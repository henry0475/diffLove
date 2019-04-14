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

	service.HandleFunc("/visited/v1/points", listOfPoint)
	service.HandleFunc("/visited/v1/point/desc", infoOfDescPoint)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
