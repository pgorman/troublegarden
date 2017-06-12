// Trouble Garden is a web-based trouble ticket system for small IT departments.
package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"time"
)

type session struct {
	uid int
	email string
	name string
	cookie string
	expires time.Time
}

var sessions = make([]session, 0, 10)
var html *string
var templates *template.Template

func handleNotice(w http.ResponseWriter, r *http.Request) {
/*
	type pageData struct {
		Title  string
		Notice string
	}
	p := pageData{
		Title:  "Notice",
		Notice: "This is important!",
	}
	err := templates.ExecuteTemplate(w, "notice")
*/
	err := templates.ExecuteTemplate(w, "header", "Notice")
	if err != nil {
		log.Println(err)
	}
	err = templates.ExecuteTemplate(w, "notice", "This is the notice text.")
	if err != nil {
		log.Println(err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello.\n"))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("sessionid")
	for _, s := range sessions {
		if s.cookie == cookie.Value {
			templates.Execute(w, "Welcome!")
			return
		}
	}
/*
	type user struct{
		username  string
		password  string
	}
	u := user{
		"paulg",
		"53cret",
	}
	if r.Method == "POST" {
		r.ParseForm()
		if r.Form.Get("user") == u.username && r.Form.Get("password") == u.password {
			sid := sessionID()
			sessionIDs = append(sessionIDs, sid)
			http.SetCookie(w, &http.Cookie{Name: "sessionid", Value: sid, HttpOnly: true, Secure: true, Expires: time.Now().Add(7 * 24 * time.Hour)})
			t, _ := template.ParseFiles("notice.tmpl")
			t.Execute(w, "Welcome!")
			return
		} else {
			t, _ := template.ParseFiles("notice.tmpl")
			t.Execute(w, "Authentication failed. :(")
			return
		}
	}
	t, _ := template.ParseFiles("login.tmpl")
	t.Execute(w, nil)
*/
}

func main() {
	html = flag.String("html", "/var/www/html/troublegarden", "Specify the root path to HTML templates.")
	ip := flag.String("ip", "", "Specify the IP address on which to listen. By default, listen on all interfaces.")
	port := flag.String("port", "9000", "Specify the TCP port on which to listen for incoming connections.")
//	sqlite := flag.String("sqlite", "/var/lib/troublegarden.db", "Specify the path to an SQLite3 database where we save messages.")
	tlscert := flag.String("tlscert", "/etc/ssl/certs/ssl-cert-snakeoil.pem", "Specify the path the the TLS certificate file.")
	tlskey := flag.String("tlskey", "/etc/ssl/private/ssl-cert-snakeoil.key", "Specify the path to the TLS key file.")
	verbose := flag.Bool("v", false, "Enable verbose logging.")
	flag.Parse()

	templates = template.Must(template.ParseGlob(*html + "/*"))

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/notice", handleNotice)

	if *tlskey != "" && *tlscert != "" {
		err := http.ListenAndServeTLS(*ip+":"+*port, *tlscert, *tlskey, nil)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := http.ListenAndServe(*ip+":"+*port, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
	if *verbose {
		log.Println("Trouble Garden listening on", *ip+":"+*port)
	}
}
