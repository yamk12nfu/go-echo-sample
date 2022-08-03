package main

import "go-echo-sample/app/infrastructure"

func main() {
	router := infrastructure.NewRouter()
	router.Start()
}
