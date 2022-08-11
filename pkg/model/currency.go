package model

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Currency struct {
	Value map[string]*Rater
}

func (c *Currency) RunGenerate() (err error) {
	c.Value = map[string]*Rater{}
	pairStr := os.Getenv("CUR_PAIR")
	pairSlice := strings.Split(pairStr, ",")

	chance, err = strconv.ParseFloat(os.Getenv("CHANCE"), 64)
	if err != nil {
		return err
	}

	startTime, err := time.Parse("2006-01-02", os.Getenv("START_TIME"))
	if err != nil {
		return err
	}
	endTime, err := time.Parse("2006-01-02", os.Getenv("END_TIME"))
	if err != nil {
		return err
	}
	if startTime.After(endTime) {
		return fmt.Errorf("error start and end time")
	}

	start = startTime.Unix()
	end = endTime.Unix()

	for _, p := range pairSlice {
		c.Value[p] = newRate()
	}

	for key, val := range c.Value {
		go func(k string, v *Rater) {
			v.Worker(context.Background(), k)
		}(key, val)
	}
	return nil
}

func (c *Currency) GetVal(key string) *Rater {
	return c.Value[key]
}
