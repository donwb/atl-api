package models

import (
	"errors"
	_ "fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"strconv"
	"time"
)

const ROOT_URL = "http://s.co/"

func AddURL(username string, url string) (string, error) {
	// check user exists
	ex := userExists(username)
	if !ex {
		return "", errors.New("The user doesn't exist")
	}

	// create the score
	score := time.Now().Unix()

	// shorten url
	shortURL := ROOT_URL + strconv.FormatInt(score, 10)

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

func Resolve(shortURL string) string {
	urlToResolve := ROOT_URL + shortURL

	conn, _ := Connect()

	conn.Do("SELECT", DB_URLRESOLVE)

	// get the user and url
	values, _ := redis.Values(conn.Do("HGETALL", urlToResolve))

	user := string(values[1].([]byte))
	fullURL := string(values[3].([]byte))

	// increment the counter for this url
	conn.Do("HINCRBY", urlToResolve, "clicks", 1)

	// increment the overall counter for this user
	conn.Do("SELECT", DB_USER)
	conn.Do("HINCRBY", user, "clicks", 1)

	return fullURL
}

func GetUrls(username string) []map[string]string {
	conn, _ := Connect()

	conn.Do("SELECT", DB_URLHISTORY)

	values, _ := redis.Strings(conn.Do("ZREVRANGE", username, 0, -1))

	urlsSlice := make([]map[string]string, len(values))

	for i, val := range values {
		log.Println("value", val)
		m := map[string]string{"url": val}
		urlsSlice[i] = m
	}

	return urlsSlice
}
