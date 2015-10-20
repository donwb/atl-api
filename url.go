package main

import (
	"github.com/goji/param"
	"github.com/zenazn/goji/web"
	"log"
	"net/http"
)

type URL struct {
	Username string `param:"username"`
	Url      string `param:"url"`
}

func getURLs(c web.C, w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("results of getURLS"))
}

func getURLsProto(c web.C, w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("results of getURLS"))
}

func createShortURL(c web.C, w http.ResponseWriter, r *http.Request) {
	var url URL
	r.ParseForm()
	err := param.Parse(r.Form, &url)
	logIf(err)

	log.Printf("Create a short url for: %s  User: %s\n", url.Url, url.Username)
}

func createShortURLProto(c web.C, w http.ResponseWriter, r *http.Request) {
	var url URL
	r.ParseForm()
	err := param.Parse(r.Form, &url)
	logIf(err)

	log.Printf("Create a short url for: %s  User: %s\n", url.Url, url.Username)
}
