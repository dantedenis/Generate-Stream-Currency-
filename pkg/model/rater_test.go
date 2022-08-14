package model

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalln("error loading .env file:", err)
	}

	chance, err = strconv.ParseFloat(os.Getenv("CHANCE"), 64)
	if err != nil {
		log.Fatalln(err)
	}

	startTime, err := time.Parse("2006-01-02", os.Getenv("START_TIME"))
	if err != nil {
		log.Fatalln(err)
	}
	endTime, err := time.Parse("2006-01-02", os.Getenv("END_TIME"))
	if err != nil {
		log.Fatalln(err)
	}
	if startTime.After(endTime) {
		log.Fatalln("error start and end time")
	}

	start = startTime.Unix()
	end = endTime.Unix()

	rand.Seed(time.Now().Unix())
}

func Test_Generator(t *testing.T) {
	res := make(chan pair)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	go func() {
		generator(ctx, res)
	}()

	for {
		select {
		case <-time.After(1 * time.Second):
			t.Error("goroutine is work")
			return
		default:
			temp, ok := <-res
			if !ok {
				return
			}
			if temp.f < 0 {
				t.Error("returns error value float", temp.f)
				//return
			}
			if !temp.t.Before(time.Unix(end, 0)) {
				t.Error("returns error value:", temp.t, "before", time.Unix(end, 0))
				//return
			}
			if !temp.t.After(time.Unix(start, 0)) {
				t.Error("returns error value:", temp.t, "after", time.Unix(start, 0))
			}
		}
	}
}
