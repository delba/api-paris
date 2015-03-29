package models

import "github.com/garyburd/redigo/redis"

var c redis.Conn

func init() {
	var err error

	c, err = redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
}
