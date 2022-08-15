package api

import (
	"context"
	"generate_stream_currency/internal/app/contract"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)

type Server struct {
	fastRouter *router.Router
	service    contract.ApiService
}

func NewServer(serv contract.ApiService) *Server {
	s := &Server{
		fastRouter: router.New(),
		service:    serv,
	}

	s.configureRouter()

	return s
}

func (s *Server) Run(ctx context.Context) (err error) {
	port := os.Getenv("SERV_PORT")
	log.Println("Service will be started on", port)

	go func() {
		err = fasthttp.ListenAndServe(port, s.fastRouter.Handler)
	}()
	<-ctx.Done()
	return err
}

func (s *Server) configureRouter() {
	s.fastRouter.GET("/values", s.getValues)
	s.fastRouter.GET("/health", s.health)
}
