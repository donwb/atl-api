package main

import (
	"flag"
	controllers "github.com/donwb/atl-api/controllers"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

func main() {
	flag.Set("bind", ":3000")

	goji.Get("/", RootRoute)

	// User routes
	goji.Get("/v1/getUser/:user", controllers.FindUser)
	goji.Get("/v2/getUser/:user", controllers.FindUserProto)
	goji.Post("/v1/createUser", controllers.CreateUser)

	// URL routes
	goji.Get("/v1/getURLs/:user", controllers.GetURLs)
	goji.Post("/v1/createShortURL", controllers.CreateShortURL)
	goji.Get("/v1/resolveURL/:shortURL", controllers.ResolveURL)

	// Light it up!
	goji.Serve()
}

func RootRoute(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Move along.... nothing to see here..."))
}
