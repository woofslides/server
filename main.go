package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/getTip", getTip)
	http.HandleFunc("/getView", getView)
	http.HandleFunc("/navNext", navNext)
	http.HandleFunc("/navPrev", navPrev)
	http.HandleFunc("/getFiles", getFiles)
	http.HandleFunc("/load", load)

	http.HandleFunc("/connectToController", connectToController)
	http.HandleFunc("/connectToViewer", connectToViewer)

	fmt.Println(http.ListenAndServe(":8080", nil))
}

func getTip(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.RemoteAddr)
}

func getView(w http.ResponseWriter, r *http.Request) {
}

func navNext(w http.ResponseWriter, r *http.Request) {}

func navPrev(w http.ResponseWriter, r *http.Request) {}

func getFiles(w http.ResponseWriter, r *http.Request) {}

func load(w http.ResponseWriter, r *http.Request) {}

func connectToController(w http.ResponseWriter, r *http.Request) {}

func connectToViewer(w http.ResponseWriter, r *http.Request) {}
