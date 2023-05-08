package main

import (
	"net/http"

	redis "github.com/go-redis/redis/v8"
	"github.com/teo/vega/controller"
	"github.com/teo/vega/database"
)

func init() {

	//redis connection
	database.Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost" + ":" + "6379", // host:port of the redis server
		Password: "",                         // no password set
		DB:       1,
	})

}

func main() {

	http.HandleFunc("/scores", controller.HttpHandler)

	http.ListenAndServe(":8080", nil)

}
