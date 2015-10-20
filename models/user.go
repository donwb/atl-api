package models

import (
	_ "fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

type UserModel struct {
	Name   string `redis:"name"`
	Clicks int    `redis:"clicks"`
}

func FindByUsername(username string) UserModel {
	conn, err := Connect()
	if err != nil {
		log.Println("Connection error:", err)
	}

	conn.Do("SELECT", DB_USER)

	values, _ := redis.Values(conn.Do("HGETALL", username))

	var user UserModel
	if err := redis.ScanStruct(values, &user); err != nil {
		log.Fatal(err)
	}

	return user
}

func AddUser(username string, fullname string) bool {
	conn, err := Connect()
	if err != nil {
		log.Println("Connection error:", err)
	}

	conn.Do("SELECT", DB_USER)

	exists, _ := redis.Bool(conn.Do("EXISTS", username))

	if !exists {
		conn.Do("HMSET", username,
			"name", fullname,
			"clicks", 0)
	}

	return exists
}
