package main

import (
	"encoding/json"
	"net/http"

	"github.com/henry0475/diffLove/apis/dlAuthAPI/auth"
	"github.com/micro/go-micro/errors"
)

func regularLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "POST" {
		if len(r.Form.Get("username")) < 6 || len(r.Form.Get("password")) < 7 {
			http.Error(
				w,
				errors.BadRequest("com.liwenbin.dev.dl.api.dlAuth", "Wrong password or username").Error(),
				403,
			)
			return
		}

		userInfo, err := auth.Login(r.Form.Get("username"), r.Form.Get("password"))
		if err != nil {
			http.Error(
				w,
				errors.BadRequest("com.liwenbin.dev.dl.api.dlAuth", err.Error()).Error(),
				502,
			)
			return
		}

		// marshal response
		b, _ := json.Marshal(map[string]interface{}{
			"Status": 200,
			"Data":   userInfo,
		})
		// write response
		w.Write(b)
	} else {
		http.Error(
			w,
			errors.BadRequest("com.liwenbin.dev.dl.api.dlAuth", "The requesting method is wrong").Error(),
			502,
		)
		return
	}
}

func regularRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "POST" {
		if len(r.Form.Get("username")) < 6 || len(r.Form.Get("password")) < 7 {
			http.Error(
				w,
				errors.BadRequest("com.liwenbin.dev.dl.api.dlAuth", "Wrong password or username").Error(),
				403,
			)
			return
		}

		userInfo, err := auth.Register(r.Form.Get("username"), r.Form.Get("password"), r.Form.Get("gender"))
		if err != nil {
			http.Error(
				w,
				errors.BadRequest("com.liwenbin.dev.dl.api.dlAuth", err.Error()).Error(),
				502,
			)
			return
		}

		// marshal response
		b, _ := json.Marshal(map[string]interface{}{
			"Status": 200,
			"Data":   userInfo,
		})
		// write response
		w.Write(b)
	} else {
		http.Error(
			w,
			errors.BadRequest("com.liwenbin.dev.dl.api.dlAuth", "The requesting method is wrong").Error(),
			502,
		)
		return
	}
}
