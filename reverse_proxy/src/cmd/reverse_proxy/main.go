package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	rp        *httputil.ReverseProxy
	from      string = "now-15m"
	to        string = "now"
	country   string
	msisdn    string = "*"
	role      string
	config    tomlConfig
	token     string
	signature string
)

func setupProxy(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse arguments

	//if token cookie exists
	token_cookie, err := r.Cookie("token")
	if err == nil {
		log.Println("Cookie found:", token_cookie.Value)
		token = token_cookie.Value
		process(w, r)

	} else if r.Form.Get("token") != "" {
		log.Println("Token Query param found")
		process(w, r)
	}

	w.Write([]byte("Access Forbidden"))
	return
}

func process(w http.ResponseWriter, r *http.Request) {
	setParams(w, r)
	decrypt_token()

	log.Println("isAuth", isAuthorized())
	if isAuthorized() {
		redirect(w, r)
	}
}

func setParams(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	log.Println("Params => token:", r.Form.Get("token"), "country:", r.Form.Get("country"), "from:", r.Form.Get("from"), "to:", r.Form.Get("to"), "msisdn:", r.Form.Get("msisdn"))

	if r.Form.Get("token") != "" {
		token = r.Form.Get("token")
		setTokenAsCookie(w, r)
	}

	if r.Form.Get("country") != "" {
		country = r.Form.Get("country")
	}

	if r.Form.Get("from") != "" {
		from = r.Form.Get("from")
	}

	if r.Form.Get("to") != "" {
		to = r.Form.Get("to")
	}

	if r.Form.Get("msisdn") != "" {
		msisdn = r.Form.Get("msisdn")
	}

	if r.Form.Get("id_partner") != "" {
		role = r.Form.Get("id_partner")
	}
}

func isAuthorized() bool {
	//log.Println("signature", signature)
	//log.Println("config signature", config.Signature)
	log.Println("SIGNATUre a comaprer:" + config.Signature + "&" + country + "&" + from + "&" + to + "&" + msisdn + "&" + role)
	return (signature == config.Signature+"&"+country+"&"+from+"&"+to+"&"+msisdn+"&"+role)
}

func index(w http.ResponseWriter, r *http.Request) {
	//log.Println(r.URL)

	if r.URL.Path == "/" {
		//log.Println("setupproxy")
		setupProxy(w, r)
	} else {
		toencode := ""
		//add Authorization in Header

		for _, currentRole := range config.Roles {
			if currentRole.Id == role {
				toencode = currentRole.Username + ":" + currentRole.Password
				break
			}
		}

		auth_key := base64.StdEncoding.EncodeToString([]byte(toencode))
		r.Header.Set("Authorization", "Basic "+auth_key)
		//log.Println("auth_key", auth_key)

		//log.Println("reverseproxy")
		rp.ServeHTTP(w, r)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	var redirectPage = `<!DOCTYPE html>
	<html lang="en">
	<head>
	    <meta charset="UTF-8">
	    <title>Redirection</title>
	    <SCRIPT language="JavaScript">
		    document.location.href="http://localhost:9090/app/kibana#/dashboard/dashboard_canalv2?_g=(refreshInterval:(display:Off,pause:!f,value:0),time:(from:'` + from + `',mode:absolute,to:'` + to + `T21:59:59.999Z'))&_a=(query:(query_string:(analyze_wildcard:!t,query:'msisdn:` + msisdn + `')))"
	    </SCRIPT>
	</head>
	</html>`
	log.Println("redirect")
	w.Header().Add("Content-Type", "text/html")
	w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Add("Pragma", "no-cache")
	w.Header().Add("Expires", "0")
	w.Header().Add("Www-Authenticate", "Basic")
	w.Header().Add("Authorization", "Basic c3VwZXJhZG1pbjpzdXBlcmFkbWlu")
	w.Write([]byte(redirectPage))
}

func setupLog() {
	f, err := os.OpenFile("/var/log/oma-canalv2/reverseproxy.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Print("error opening file: %v", err)
	}

	log.SetOutput(f)
}

func main() {
	setupLog()

	readConfig()

	//set up reverse proxy
	rp = httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   config.UrlKibana,
		User:   url.UserPassword("admin_kibana", "admin_kibana"),
	})

	log.Println("Reverse proxy is listening on ", config.Addr)
	http.ListenAndServe(config.Addr, http.HandlerFunc(index))
}

///////////////////////////////////////////////
//                TOKEN
//////////////////////////////////////////////
func decrypt_token() {
	//replace spaces with '+'
	token = strings.Replace(token, " ", "+", -1)

	//get decrypted token
	signature = string(cbcDecrypt(token))
	//replace null characters
	signature = strings.Replace(signature, "\x00", "", -1)
	log.Println("Decrypted token :", signature)
}

func cbcDecrypt(text1 string) []byte {
	key := []byte(config.DecryptKey)
	ciphertext, _ := base64.StdEncoding.DecodeString(string(text1))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)
	ciphertext = PKCS5UnPadding(ciphertext)
	return ciphertext
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	//unpadding := int(src[length-1])
	//log.Println("8: len:", src, "  unpadding:", unpadding)
	return src[:(length)]
}

func setTokenAsCookie(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("Set Token as Cookie")
	cookie := &http.Cookie{Name: "token", Value: r.Form.Get("token"), Expires: time.Now().Add(30 * 24 * time.Hour), HttpOnly: true}
	http.SetCookie(w, cookie)
	r.AddCookie(cookie)
}

///////////////////////////////////////////////
//                CONFIG
//////////////////////////////////////////////
type tomlConfig struct {
	UrlKibana  string
	Addr       string
	Country    string
	DecryptKey string
	Signature  string
	Roles      map[string]roles
}

type roles struct {
	Id       string
	Username string
	Password string
}

func readConfig() {
	if _, err := toml.DecodeFile("src/cmd/reverse_proxy/conf.ini", &config); err != nil {
		log.Println(err)
		return
	}

	log.Println("===============Config file===============")
	log.Println("         country ==>", config.Country)
	log.Println("       urlKibana ==>", config.UrlKibana)
	log.Println("            addr ==>", config.Addr)
	log.Println("      decryptKey ==>", config.DecryptKey)
	log.Println("       signature ==>", config.Signature)
	log.Println("  Admin username ==>", config.Roles["Admin"].Username)
	log.Println("        Admin pw ==>", config.Roles["Admin"].Password)
	log.Println("Partner username ==>", config.Roles["Partner"].Username)
	log.Println("      Partner pw ==>", config.Roles["Partner"].Password)
	log.Println("=========================================\n")

	for _, v := range config.Roles {
		log.Println(v)
	}

}
