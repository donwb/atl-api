package main

import (
	"flag"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

func main() {
	flag.Set("bind", ":3000")

	goji.Get("/", RootRoute)

	// User routes
	goji.Get("/v1/getUser/:user", findUser)
	goji.Post("/v1/createUser", createUser)

	// URL routes
	goji.Get("/v1/getURLs/:user", getURLs)
	goji.Post("/v1/createShortURL", createShortURL)

	// Light it up!
	goji.Serve()
}

func RootRoute(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Move along.... nothing to see here..."))
}
