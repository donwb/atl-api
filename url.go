package main

import (
	"encoding/json"
	models "github.com/donwb/atl-api/models"
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

	urls := models.GetUrls("donwb")

	res, _ := json.Marshal(urls)

	w.Write(res)
}

func createShortURL(c web.C, w http.ResponseWriter, r *http.Request) {
	var url URL
	r.ParseForm()
	err := param.Parse(r.Form, &url)
	logIf(err)

	log.Printf("Create a short url for: %s  User: %s\n", url.Url, url.Username)

	shortURL, err := models.AddURL(url.Username, url.Url)

	w.Write([]byte(shortURL))

}

func resolveURL(c web.C, w http.ResponseWriter, r *http.Request) {
	var shortURL = c.URLParams["shortURL"]

	fullURL := models.Resolve(shortURL)

	m := map[string]string{"fullURL": fullURL}
	res, _ := json.Marshal(m)

	w.Write(res)

}
