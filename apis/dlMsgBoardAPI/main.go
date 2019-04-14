package main

import (
	"log"

	"github.com/micro/go-web"
)

func main() {
	service := web.NewService(
		web.Name("com.liwenbin.dev.dl.api.dlMsgBoard"),
		web.Version("0.0.1"),
	)

	service.HandleFunc("/msgboard/v1/publish", msgPublish)
	service.HandleFunc("/msgboard/v1/list", msgList)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
