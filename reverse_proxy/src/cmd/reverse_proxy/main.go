package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

var (
	rp *httputil.ReverseProxy
	//from      string = "now-15m"
	//to        string = "now"
	//country   string
	//msisdn    string = "*"
	profile string
	config  tomlConfig
	//token     string
	//signature string
)

func main() {
	readConfig()
	setupLog()

	//set up reverse proxy
	rp = httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   config.UrlKibana,
	})

	log.Println("Reverse proxy is listening on ", config.Addr)
	http.ListenAndServe(config.Addr, http.HandlerFunc(index))
}

func index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("Path loaded:" + r.URL.Path)
	if r.URL.Path == "/" {
		log.Println("Path 3:" + r.URL.Path)
		setupProxy(w, r)
	} else if r.URL.Path == "/test" {
		log.Println("Path 2:" + r.URL.Path)
		redirect(w, r)
	} else {
		log.Println("Path 1:" + r.URL.Path)
		r.Header.Set("Authorization", "Basic "+profile)
		rp.ServeHTTP(w, r)
	}
}

func setupProxy(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse arguments

	if r.Form.Get("token") != "" {
		//log.Println("Params:", r.Form)
		tokenDecrypted := decryptToken(r)
		if isAuthorized(r, tokenDecrypted) {
			redirect(w, r)
		}
	}

	w.Write([]byte("Access Forbidden"))
	return
}

func isAuthorized(r *http.Request, tokenDecrypted string) bool {
	//log.Println("SIGNATURE:" + tokenDecrypted)
	isAuth := (tokenDecrypted == config.Signature+"&"+r.Form.Get("country")+"&"+r.Form.Get("from")+"&"+r.Form.Get("to")+"&"+r.Form.Get("profile"))
	log.Println("User isAuth:", isAuth)
	return isAuth
}

func redirect(w http.ResponseWriter, r *http.Request) {
	log.Println("Path:", r.URL.Path, "   Build Kibana URL with params")
	//var redirectPage = `<!DOCTYPE html>
	//<html lang="en">
	//<head>
	//    <meta charset="UTF-8">
	//    <title>Redirection</title>
	//    <SCRIPT language="JavaScript">
	//    	document.location.href="http://localhost:9090/app/kibana#/dashboard/dashboard_canalv2?_g=(refreshInterval:(display:Off,pause:!f,value:0),time:(from:'` + r.Form.Get("from") + `',mode:absolute,to:'` + r.Form.Get("to") + `T21:59:59.999Z'))&_a=(query:(query_string:(analyze_wildcard:!t,query:'msisdn:` + r.Form.Get("msisdn") + `')))"
	//    </SCRIPT>
	//</head>
	//</html>`

	//add Authorization in Header
	for _, currentRole := range config.Roles {
		if currentRole.Profile == r.Form.Get("profile") {
			log.Println("User Profile:", currentRole.Profile, "User Role:", currentRole.Username)
			toencode := currentRole.Username + ":" + currentRole.Password
			profile = base64.StdEncoding.EncodeToString([]byte(toencode))
			//set profile as cookie
			//cookie := &http.Cookie{Name: "profile", Value: profile, Expires: time.Now().Add(30 * 24 * time.Hour), HttpOnly: true}
			//http.SetCookie(w, cookie)
			//r.AddCookie(cookie)
			break
		}
	}

	log.Println("Access to kibana")
	//
	if r.URL.Path == "/" {
		http.PostForm("http://localhost:9090/test", url.Values{})

	} else {
		http.Redirect(w, r, "http://localhost:9090/app/kibana#/dashboard/dashboard_canalv2?_g=(refreshInterval:(display:Off,pause:!f,value:0),time:(from:'"+r.Form.Get("from")+"',mode:absolute,to:'"+r.Form.Get("to")+"T21:59:59.999Z'))&_a=(query:(query_string:(analyze_wildcard:!t,query:'msisdn:"+r.Form.Get("msisdn")+"')))", 302)
	}

	//http.PostForm("http://localhost:9090/app/kibana#/dashboard/dashboard_canalv2?_g=(refreshInterval:(display:Off,pause:!f,value:0),time:(from:'"+r.Form.Get("from")+"',mode:absolute,to:'"+r.Form.Get("to")+"T21:59:59.999Z'))&_a=(query:(query_string:(analyze_wildcard:!t,query:'msisdn:"+r.Form.Get("msisdn")+"')))", url.Values{})

	//http.Redirect(w, r, "http://localhost:9090/app/kibana#/dashboard/dashboard_canalv2?_g=(refreshInterval:(display:Off,pause:!f,value:0),time:(from:'"+r.Form.Get("from")+"',mode:absolute,to:'"+r.Form.Get("to")+"T21:59:59.999Z'))&_a=(query:(query_string:(analyze_wildcard:!t,query:'msisdn:"+r.Form.Get("msisdn")+"')))", 302)
	//w.Write([]byte(redirectPage))
}

///////////////////////////////////////////////
//                TOKEN
//////////////////////////////////////////////
func decryptToken(r *http.Request) string {
	//replace spaces with '+'
	newToken := strings.Replace(r.Form.Get("token"), " ", "+", -1)
	r.Form.Set("token", newToken)

	//get decrypted token
	decryptedToken := string(cbcDecrypt(newToken))

	//replace null characters
	decryptedToken = strings.Replace(decryptedToken, "\x00", "", -1)
	log.Println("Decrypted token :", decryptedToken)

	return decryptedToken
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

///////////////////////////////////////////////
//                CONFIG
//////////////////////////////////////////////
type tomlConfig struct {
	UrlKibana   string
	Addr        string
	Country     string
	PathLogFile string
	DecryptKey  string
	Signature   string
	Roles       map[string]roles
}

type roles struct {
	Profile  string
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
	log.Println("  path log file  ==>", config.PathLogFile)
	log.Println("      decryptKey ==>", config.DecryptKey)
	log.Println("       signature ==>", config.Signature)
	log.Println("  Admin username ==>", config.Roles["Manager"].Username)
	log.Println("        Admin pw ==>", config.Roles["Manager"].Password)
	log.Println("Partner username ==>", config.Roles["Partner"].Username)
	log.Println("      Partner pw ==>", config.Roles["Partner"].Password)
	log.Println("=========================================\n")

}

func setupLog() {
	f, err := os.OpenFile(config.PathLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Print("error opening file: %v", err)
	}

	log.SetOutput(f)
}
