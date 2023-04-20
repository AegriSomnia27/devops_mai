package server

import (
	utils "go_rest_api/pkg"
	"log"
	"net/http"
	"os"
	"time"
)

func helloWorld(w http.ResponseWriter, req *http.Request) {
	log.Println("helloWorld, from: ", req.URL.User, "time: ", time.Now())

	resp := utils.SystemResponse{
		Message: "Hello, world!",
		Time:    time.Now(),
	}
	_, err := w.Write([]byte(resp.String()))
	if err != nil {
		log.Println("something bad happened, send help")
	}

	return
}

func destroyAllHumans(w http.ResponseWriter, req *http.Request) {
	log.Println("why are we still here?", "time: ", time.Now())

	os.Exit(1)
}
