package models

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

var rPool *redis.Pool
var ConnectString = "127.0.0.1:6379" // localhost by default for tests

const DB_USER = 0
const DB_URLHISTORY = 1
const DB_URLRESOLVE = 2

/*
	Provides a pooling mechinism on top of a REDIS connection
		server: address of the redis server
		password: pass "" if no password is set
*/
func NewPool(server *string, password string) *redis.Pool {
	log.Println("REDIS Server:", *server)

	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", *server)
			if err != nil {
				fmt.Println("ERROR:", err)
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

/*
	Call Connect any time you need a connection to redis.  Uses a singleton
	pattern to initialze the pool if not already set.
*/
func Connect() (redis.Conn, error) {

	if rPool == nil {
		log.Println("REDIS connection nil, initializing pool")
		rPool = NewPool(&ConnectString, "")
	}

	conn := rPool.Get()

	return conn, nil
}
