package models

import (
	"errors"
	_ "fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"strconv"
	"time"
)

func AddURL(username string, url string) (string, error) {
	// check user exists
	ex := userExists(username)
	if !ex {
		return "", errors.New("The user doesn't exist")
	}

	// create the score
	score := time.Now().Unix()

	// shorten url
	shortURL := "http://this.com/" + strconv.FormatInt(score, 10)

	// add to zlist by username
	conn, _ := Connect()
	conn.Do("SELECT", DB_URLHISTORY)

	conn.Do("ZADD", username, score, shortURL)

	// add to url resolver db
	conn.Do("SELECT", DB_URLRESOLVE)
	conn.Do("HMSET", shortURL,
		"username", username,
		"url", url)

	return shortURL, nil
}

func userExists(username string) bool {
	conn, err := Connect()
	if err != nil {
		log.Println("Connection error:", err)
	}

	conn.Do("SELECT", DB_USER)

	exists, _ := redis.Bool(conn.Do("EXISTS", username))

	return exists
}
