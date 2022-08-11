package api

import (
	"generate_stream_currency/internal/app/contract"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"strings"
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

func (s *Server) Run() error {
	port := strings.Join([]string{":", os.Getenv("SERV_PORT")}, "")
	log.Println("Service will be started on", port)

	return fasthttp.ListenAndServe(port, s.fastRouter.Handler)
}

func (s *Server) configureRouter() {
	s.fastRouter.GET("/values", s.GetValues)
	s.fastRouter.GET("/health", s.health)
}
