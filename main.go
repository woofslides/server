package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/getip", getip)
	http.HandleFunc("/getview", getview)
	http.HandleFunc("/navnext", navnext)
	http.HandleFunc("/navprev", navprev)
	http.HandleFunc("/getfiles", getfiles)
	http.HandleFunc("/load", load)
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func getip(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.RemoteAddr)
}

func getview(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr= %q\n", r.RemoteAddr)
}

func navnext(w http.ResponseWriter, r *http.Request) {}

func navprev(w http.ResponseWriter, r *http.Request) {}

func getfiles(w http.ResponseWriter, r *http.Request) {}

func load(w http.ResponseWriter, r *http.Request) {}
