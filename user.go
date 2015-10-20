package main

import (
	"encoding/json"
	models "github.com/donwb/atl-api/models"
	"github.com/goji/param"
	"github.com/zenazn/goji/web"
	"log"
	"net/http"
)

type User struct {
	Username string `param:"username" json:"username"`
	Name     string `param:"name" json:"name"`
	Clicks   int    `param:"clicks" json:"clicks"`
}

func createUser(c web.C, w http.ResponseWriter, r *http.Request) {

	var user User

	r.ParseForm()
	err := param.Parse(r.Form, &user)
	logIf(err)

	exists := models.AddUser(user.Username, user.Name)

	m := map[string]bool{"exists": exists}
	res, _ := json.Marshal(m)

	w.Write(res)

}

func findUser(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	username := c.URLParams["user"]

	userModel := models.FindByUsername(username)

	u := User{
		Username: username,
		Name:     userModel.Name,
		Clicks:   userModel.Clicks,
	}

	res, _ := json.Marshal(u)

	w.Write(res)
}

func createUserProto(c web.C, w http.ResponseWriter, r *http.Request) {

	var user User

	r.ParseForm()
	err := param.Parse(r.Form, &user)
	logIf(err)

	log.Printf("Creating Proto user: %s - name: %s\n", user.Username, user.Name)

}
