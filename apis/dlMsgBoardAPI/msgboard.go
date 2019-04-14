package main

import (
	"encoding/json"
	"net/http"

	"github.com/henry0475/diffLove/apis/dlMsgBoardAPI/msgboard"
	"github.com/micro/go-micro/errors"
)

func msgPublish(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "POST" {
		if len(r.Header.Get("Authorization")) < 40 {
			http.Error(
				w,
				errors.BadRequest("com.liwenbin.dev.dl.api.dlMsgBoard", "Wrong parameters.").Error(),
				403,
			)
			return
		}

		err := msgboard.PublishMsg(r.Header.Get("Authorization"), r.Form.Get("bid"), r.Form.Get("content"), r.Form.Get("isPublic"))
		if err != nil {
			http.Error(
				w,
				errors.BadRequest("com.liwenbin.dev.dl.api.dlMsgBoard", err.Error()).Error(),
				502,
			)
			return
		}

		// marshal response
		b, _ := json.Marshal(map[string]interface{}{
			"Status": 200,
		})
		// write response
		w.Write(b)
	} else {
		http.Error(
			w,
			errors.BadRequest("com.liwenbin.dev.dl.api.dlMsgBoard", "The requesting method is wrong").Error(),
			502,
		)
		return
	}
}

func msgList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "GET" {
		if len(r.Header.Get("Authorization")) < 40 {
			http.Error(
				w,
				errors.BadRequest("com.liwenbin.dev.dl.api.dlMsgBoard", "Wrong password or username").Error(),
				403,
			)
			return
		}

		msgList, err := msgboard.GetListOfMsg(r.Header.Get("Authorization"), r.Form.Get("startLine"))
		if err != nil {
			http.Error(
				w,
				errors.BadRequest("com.liwenbin.dev.dl.api.dlMsgBoard", err.Error()).Error(),
				502,
			)
			return
		}

		// marshal response
		b, _ := json.Marshal(map[string]interface{}{
			"Status": 200,
			"Data":   msgList,
		})
		// write response
		w.Write(b)
	} else {
		http.Error(
			w,
			errors.BadRequest("com.liwenbin.dev.dl.api.dlMsgBoard", "The requesting method is wrong").Error(),
			502,
		)
		return
	}
}
