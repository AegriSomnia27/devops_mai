package server

import (
	"fmt"
	"log"
	"net/http"
)

func StartHTTPServer(host string, port uint16) (*http.Server, error) {
	svc := &http.Server{
		Addr: fmt.Sprintf("%s:%d", host, port),
	}

	http.HandleFunc("/hello", helloWorld)
	http.HandleFunc("/term", destroyAllHumans)

	go func() {
		if err := svc.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	return svc, nil
}
