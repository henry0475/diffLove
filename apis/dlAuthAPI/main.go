package main

import (
	"log"

	"github.com/micro/go-web"
)

func main() {
	service := web.NewService(
		web.Name("com.dl.api.auth"),
		web.Version("0.0.1"),
	)

	service.HandleFunc("/auth/login", regularLogin)
	service.HandleFunc("/auth/register", regularRegister)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
