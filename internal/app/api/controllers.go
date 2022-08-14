package api

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
)

func (s *Server) health(ctx *fasthttp.RequestCtx) {
	log.Println("handle health")
	ctx.SetStatusCode(http.StatusOK)
}

func (s *Server) getValues(ctx *fasthttp.RequestCtx) {
	target := string(ctx.FormValue("target"))
	if target == "" {
		ctx.SetStatusCode(http.StatusOK)
		return
	}

	getCurrency := s.service.GetCurrency()
	currency := getCurrency.GetVal(target)
	currency.Lock()
	defer currency.Unlock()

	b, err := json.Marshal(map[string]interface{}{
		target: currency,
	})
	if err != nil {
		log.Println(err)
		ctx.SetStatusCode(http.StatusInternalServerError)
		return
	}

	ctx.Response.Header.Set("Content-Type", "application/json; charset=utf-8")
	ctx.SetStatusCode(http.StatusOK)

	_, err = ctx.Write(b)
	if err != nil {
		log.Println(err)
		return
	}
}
