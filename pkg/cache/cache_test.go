package cache

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalln("error loading .env file:", err)
	}
}

func TestNewCache(t *testing.T) {
	c := NewCache()
	if c == nil {
		t.Error("Return error")
	}

}

func TestCache_GetCurrency(t *testing.T) {
	c := NewCache()
	temp := c.GetCurrency()
	if temp.Value != nil {
		t.Error("Return error")
	}
}

func TestCache_Run(t *testing.T) {
	c := NewCache()
	assert.Nil(t, c.Run())
}
