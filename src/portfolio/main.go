package main

import (
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	appengine.Main()

	http.HandleFunc("/301", redirect)
}

func redirect(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "http://www.golang.org", 301)
}
