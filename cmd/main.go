package main

import (
	"generate_stream_currency/internal/app/api"
	"generate_stream_currency/pkg/cache"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("error loading .env file:", err)
	}

	rand.Seed(time.Now().Unix())

}

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	service := cache.NewCache()
	err := service.Run()
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		serv := api.NewServer(service)
		err := serv.Run()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	<-c
	log.Println("stop service")
}
