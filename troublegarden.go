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
	expires Time
}

var sessions = make([]session, 0, 10)

func main() {
	html := flag.String("html", "/var/www/html/troublegarden", "Specify the root path to HTML templates.")
	ip := flag.String("ip", "", "Specify the IP address on which to listen. By default, listen on all interfaces.")
	port := flag.String("port", "9000", "Specify the TCP port on which to listen for incoming connections.")
	sqlite := flag.String("sqlite", "/var/lib/troublegarden.db", "Specify the path to an SQLite3 database where we save messages.")
	tlscert := flag.String("tlscert", "/etc/ssl/certs/ssl-cert-snakeoil.pem", "Specify the path the the TLS certificate file.")
	tlskey := flag.String("tlskey", "/etc/ssl/private/ssl-cert-snakeoil.key", "Specify the path to the TLS key file.")
	verbose := flag.Bool("v", false, "Enable verbose logging.")
	flag.Parse()

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)

	var l net.Listener
	if *tlskey != "" && *tlscert != "" {
		l, err = http.ListenAndServeTLS(*ip+":"+*port, *tlscert, *tlskey, nil)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		l, err = net.ListenAndServe(*ip+":"+*port, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
	if *verbose {
		log.Println("Trouble Garden listening on", *ip+":"+*port)
	}
	defer l.Close()
}
