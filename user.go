package main

import (
	"github.com/goji/param"
	"github.com/zenazn/goji/web"
	"log"
	"net/http"
)

type User struct {
	Username string `param:"username", json:"username"`
	Name     string `param:"name", json:"name"`
}

func createUser(c web.C, w http.ResponseWriter, r *http.Request) {

	var user User

	r.ParseForm()
	err := param.Parse(r.Form, &user)
	logIf(err)

	log.Printf("Creating user: %s - name: %s\n", user.Username, user.Name)

}

func createUserProto(c web.C, w http.ResponseWriter, r *http.Request) {

	var user User

	r.ParseForm()
	err := param.Parse(r.Form, &user)
	logIf(err)

	log.Printf("Creating Proto user: %s - name: %s\n", user.Username, user.Name)

}
