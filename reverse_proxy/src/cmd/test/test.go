package main

import (
	"fmt"
	"log"
	"net/http"

	"net/url"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {
	log.Println("Listening on :9090")
	http.ListenAndServe(":9090", http.HandlerFunc(index))

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Path:", r.URL.Path)
	if r.URL.Path == "/" {
		session, err := store.Get(r, "session-name")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set some session values.
		session.Values["foo"] = "bar"
		session.Values[42] = 43
		// Save it before we write to the response/return from the handler.
		session.Save(r, w)

		r.ParseForm()
		r.PostForm.Set("TEST", "TEST")
		r.Form.Set("TEST", "TEST")
		//r.FormValue("TEST") = "TOTO"
		//r.PostFormValue("TEST")
		//r.FormValue("TEST")

		//http.Redirect(w, r, "http://localhost:9090/test", 301)
		http.PostForm("http://localhost:9090/test", url.Values{"toto": {"toto"}})
	} else if r.URL.Path == "/test" {
		r.ParseMultipartForm(64)
		r.ParseForm()

		session, err := store.Get(r, "session-name")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println(fmt.Sprint(r.PostForm.Get("TEST")))
		log.Println(fmt.Sprint(r.Form.Get("TEST")))
		log.Println(fmt.Sprint(r.Form))
		log.Println(fmt.Sprint(r.PostFormValue("TEST")))
		//log.Println(r.MultipartForm.Value("TEST"))
		log.Println(fmt.Sprint(session.Values["foo"]))
		w.Write([]byte("Hello:"))
	}
}
