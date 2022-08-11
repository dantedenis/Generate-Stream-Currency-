package model

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/rand"
	"sync"
	"time"
)

var (
	chance float64
	start  int64
	end    int64
)

type Rater struct {
	sync.Mutex
	limiter chan time.Time
	Rate    map[time.Time]float64
}

func newRate() *Rater {
	return &Rater{
		limiter: make(chan time.Time, 100),
		Rate:    map[time.Time]float64{},
	}
}

func (r *Rater) Add(k time.Time, v float64) error {
	r.Lock()
	defer r.Unlock()
	if len(r.Rate) > 99 {
		delete(r.Rate, <-r.limiter)
	}
	if _, ok := r.Rate[k]; !ok {
		r.Rate[k] = v
	} else {
		return fmt.Errorf("value is exist")
	}
	r.limiter <- k
	return nil
}

// /// ---- logic worker for pair-value
type pair struct {
	t time.Time
	f float64
}

func (r *Rater) Worker(ctx context.Context, k string) {
	value := make(chan pair)
	go func() {
		generator(value)
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println(k, "finish")
			return
		default:
			t := <-value
			err := r.Add(t.t, t.f)
			if err != nil {
				log.Println(err, k)
			}
		}
	}
}

func generator(c chan pair) {
	for {
		switch {
		case rand.Float64() < chance:
			time.Sleep(50 * time.Millisecond)
		default:
			c <- pair{
				t: time.Unix(rand.Int63n(end-start)+start, 0),
				f: math.Floor(rand.Float64()*10000) / 10000,
			}
		}
	}
}
