package controllers

import (
	"encoding/json"
	"fmt"
	models "github.com/donwb/atl-api/models"
	"github.com/goji/param"
	"github.com/golang/protobuf/proto"
	"github.com/zenazn/goji/web"
	"net/http"
)

type User struct {
	Username string `param:"username" json:"username"`
	Name     string `param:"name" json:"name"`
	Clicks   int    `param:"clicks" json:"clicks"`
}

func CreateUser(c web.C, w http.ResponseWriter, r *http.Request) {

	var user User

	r.ParseForm()
	err := param.Parse(r.Form, &user)
	logIf(err)

	exists := models.AddUser(user.Username, user.Name)

	m := map[string]bool{"exists": exists}
	res, _ := json.Marshal(m)

	w.Write(res)

}

func FindUser(c web.C, w http.ResponseWriter, r *http.Request) {

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

func FindUserProto(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-protobuf")

	username := c.URLParams["user"]

	userModel := models.FindByUsername(username)

	u := UserProto{
		Username: username,
		Name:     userModel.Name,
		Clicks:   int64(userModel.Clicks),
	}

	fmt.Println(u)

	protoRes, _ := proto.Marshal(&u)

	w.Write(protoRes)
}
