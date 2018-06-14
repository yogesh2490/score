package main

import (
	"log"
	"net/http"
	"score/common"
	sr "score/router"
)

func main() {
	//db connection
	healthController := common.GetHealthController()

	router := sr.NewRouter(healthController)
	log.Fatal(http.ListenAndServe(":8081", router))
}