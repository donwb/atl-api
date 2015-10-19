package main

import (
	"flag"
	"github.com/goji/param"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"log"
	"net/http"
)

func main() {
	flag.Set("bind", ":3000")

	goji.Get("/", RootRoute)
	goji.Get("/v1/getURLs/:user", getURLs)
	goji.Post("/v1/createUser", createUser)
	goji.Post("/v1/createShortURL", createShortURL)

	goji.Serve()
}

func RootRoute(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Move along.... nothing to see here..."))
}

func getURLs(c web.C, w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("results of getURLS"))
}

func createUser(c web.C, w http.ResponseWriter, r *http.Request) {

	var user User

	r.ParseForm()
	err := param.Parse(r.Form, &user)
	if err != nil {
		log.Println("oh shit", err)
	}

	log.Printf("Creating user: %s - name: %s\n", user.Username, user.Name)

}

func createShortURL(c web.C, w http.ResponseWriter, r *http.Request) {
	var url URL
	r.ParseForm()
	err := param.Parse(r.Form, &url)
	if err != nil {
		log.Println("oh shit")
	}

	log.Printf("Create a short url for: %s  User: %s\n", url.Url, url.Username)
}

type User struct {
	Username string `param:"username"`
	Name     string `param:"name"`
}

type URL struct {
	Username string `param:"username"`
	Url      string `param:"url"`
}
