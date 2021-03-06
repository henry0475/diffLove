package main

import (
	"log"

	"github.com/micro/go-web"
)

func main() {
	service := web.NewService(
		web.Name("com.test.api.auth"),
		web.Version("0.0.1"),
	)

	service.HandleFunc("/auth/v1/login", regularLogin)
	service.HandleFunc("/auth/v1/egister", regularRegister)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
