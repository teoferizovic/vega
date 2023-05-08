package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/teo/vega/service"
)

func HttpHandler(w http.ResponseWriter, req *http.Request) {

	params := map[string]interface{}{}

	resp := map[string]interface{}{}

	var err error

	if req.Method == "GET" {

		for k, v := range req.URL.Query() {

			params[k] = v[0]

		}

		resp, err = service.GetScores(params)

	} else if req.Method == "POST" {

		err = json.NewDecoder(req.Body).Decode(&params)

		resp, err = service.AddScore(params)

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
