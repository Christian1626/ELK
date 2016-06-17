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
	"github.com/gorilla/sessions"
)

var (
	rp     *httputil.ReverseProxy
	config tomlConfig
	store  = sessions.NewCookieStore([]byte("authentication_profile_key"))
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
		setupProxy(w, r)
	} else {
		session, err := store.Get(r, "session-profile")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//first page called by kibana
		if r.URL.Path == "/app/kibana" {
			if profile := r.Form.Get("profile"); strings.Contains(decryptToken(r.Form.Get("token")), profile) {
				// Set some session values.
				session.Values["auth"] = getBasicAuth(profile)
				session.Save(r, w)
			}
		}

		//set Authorization header
		r.Header.Set("Authorization", "Basic "+session.Values["auth"].(string))
		rp.ServeHTTP(w, r)
	}
}

func setupProxy(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse arguments

	if token := r.Form.Get("token"); token != "" {
		//log.Println("Params:", r.Form)
		tokenDecrypted := decryptToken(token)
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
	//<meta charset="UTF-8">
	//<title>Title</title>
	//</head>
	//<body>
	//	<form method="GET" action="http://localhost:9090/app/kibana#/dashboard/dashboard_canalv2?_g=(refreshInterval:(display:Off,pause:!f,value:0),time:(from:'` + r.Form.Get("from") + `',mode:absolute,to:'` + r.Form.Get("to") + `T21:59:59.999Z'))&_a=(query:(query_string:(analyze_wildcard:!t,query:'msisdn:` + r.Form.Get("msisdn") + `')))" id="form">
	//		<input type="text" name="profile" value="` + r.Form.Get("profile") + `"><br><br>
	//		<input type="submit" value="Submit">
	//	</form>
	//	<script language="JavaScript">
	//		document.getElementById("form").submit();
	//	</script>
	//</body>
	//</html>`

	log.Println("Access to kibana")
	http.Redirect(w, r, "http://localhost:9090/app/kibana?token="+r.Form.Get("token")+"&profile="+r.Form.Get("profile")+"#/dashboard/dashboard_canalv2?_g=(refreshInterval:(display:Off,pause:!f,value:0),time:(from:'"+r.Form.Get("from")+"',mode:absolute,to:'"+r.Form.Get("to")+"T21:59:59.999Z'))&_a=(query:(query_string:(analyze_wildcard:!t,query:'msisdn:"+r.Form.Get("msisdn")+"')))", 302)
	//w.Write([]byte(redirectPage))
}

func getBasicAuth(profile string) string {
	//add Authorization in Header
	for _, currentRole := range config.Roles {
		if currentRole.Profile == profile {
			auth := currentRole.Username + ":" + currentRole.Password
			return base64.StdEncoding.EncodeToString([]byte(auth))
		}
	}
	return ""
}

///////////////////////////////////////////////
//                TOKEN
//////////////////////////////////////////////
func decryptToken(token string) string {
	//replace spaces with '+'
	newToken := strings.Replace(token, " ", "+", -1)

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
