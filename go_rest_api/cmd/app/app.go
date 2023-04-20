package main

import (
	"context"
	"fmt"
	server "go_rest_api/internal"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	if err := godotenv.Load("./config/.env"); err != nil {
		panic(err)
	}

	appCfg := &server.AppConfig{}
	if err := envconfig.Process("", appCfg); err != nil {
		panic(err)
	}

	log.Println("starting server...")

	server, err := server.StartHTTPServer(appCfg.Host, appCfg.Port)
	if err != nil {
		panic(err)
	}

	log.Println("server has started")

	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)

	<-term

	log.Println("shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println(err)
	}
}
