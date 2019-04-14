package main

import (
	"encoding/json"
	"net/http"

	"github.com/henry0475/diffLove/apis/dlVisitedAPI/visited"
	"github.com/micro/go-micro/errors"
)

func listOfPoint(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "GET" {
		if len(r.Header.Get("Authorization")) < 40 {
			http.Error(
				w,
				errors.BadRequest("com.liwenbin.dev.dl.api.dlVisited", "Wrong parameters.").Error(),
				403,
			)
			return
		}

		listInfo, err := visited.GetPoints(r.Header.Get("Authorization"), r.Form.Get("vid"))
		if err != nil {
			http.Error(
				w,
				errors.BadRequest("com.liwenbin.dev.dl.api.dlVisited", err.Error()).Error(),
				502,
			)
			return
		}

		// marshal response
		b, _ := json.Marshal(map[string]interface{}{
			"Status": 200,
			"Data":   listInfo,
		})
		// write response
		w.Write(b)
	} else {
		http.Error(
			w,
			errors.BadRequest("com.liwenbin.dev.dl.api.dlVisited", "The requesting method is wrong").Error(),
			502,
		)
		return
	}
}

func infoOfDescPoint(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "GET" {
		if len(r.Header.Get("Authorization")) < 40 {
			http.Error(
				w,
				errors.BadRequest("com.liwenbin.dev.dl.api.dlVisited", "Wrong password or username").Error(),
				403,
			)
			return
		}

		info, err := visited.GetDescOfPoint(r.Header.Get("Authorization"), r.Form.Get("id"))
		if err != nil {
			http.Error(
				w,
				errors.BadRequest("com.liwenbin.dev.dl.api.dlVisited", err.Error()).Error(),
				502,
			)
			return
		}

		// marshal response
		b, _ := json.Marshal(map[string]interface{}{
			"Status": 200,
			"Data":   info,
		})
		// write response
		w.Write(b)
	} else {
		http.Error(
			w,
			errors.BadRequest("com.liwenbin.dev.dl.api.dlVisited", "The requesting method is wrong").Error(),
			502,
		)
		return
	}
}
