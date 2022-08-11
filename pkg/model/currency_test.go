package model

import (
	"log"
	"testing"
	"time"
)

/*
func Test_Rate(t *testing.T) {
	c := newRate()

	for i := 0; i < 150; i++ {
		assert.Nil(t, c.Add(timeEvent(i), 123123.13))
	}
	assert.NotNil(t, c.Add(100, 6674.2))
}

func Test_Rate2(t *testing.T) {
	c := newRate()

	assert.Nil(t, c.Add(100, 123123.13))
	assert.NotNil(t, c.Add(100, 6674.2))
}

*/

func Test(t *testing.T) {
	timer := time.Now()
	log.Println(timer.Unix())
}
