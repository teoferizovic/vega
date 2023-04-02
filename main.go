package main

import (
	"encoding/json"
	"fmt"
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

	http.HandleFunc("/scores", httpHandler)

	http.ListenAndServe(":8080", nil)

}

func httpHandler(w http.ResponseWriter, req *http.Request) {

	params := map[string]interface{}{}

	resp := map[string]interface{}{}

	var err error

	if req.Method == "GET" {

		for k, v := range req.URL.Query() {

			params[k] = v[0]

		}

		resp, err = controller.GetScores(params)

	} else if req.Method == "POST" {

		err = json.NewDecoder(req.Body).Decode(&params)

		resp, err = controller.AddScore(params)

	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")

	if err != nil {

		resp = map[string]interface{}{

			"error": err.Error(),
		}

	} else {

		if encodingErr := enc.Encode(resp); encodingErr != nil {

			fmt.Println("{ error: " + encodingErr.Error() + "}")

		}

	}

}
