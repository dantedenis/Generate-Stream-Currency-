package main

import (
	"context"
	"generate_stream_currency/internal/app/api"
	"generate_stream_currency/pkg/cache"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	service := cache.NewCache()
	err := service.Run()
	if err != nil {
		log.Fatalln(err)
	}

	serv := api.NewServer(service)
	err = serv.Run(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
}
