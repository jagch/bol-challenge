package server

import (
	"fmt"

	domain "jagch/boletia/freecurrency/internal"
	"log"

	"jagch/boletia/freecurrency/internal/platform/server/handler"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	//deps
	currencyUsecase domain.CurrencyUsecase
}

func New(host string, port uint, currencyUsecase domain.CurrencyUsecase) Server {
	srv := Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine:   gin.New(),

		currencyUsecase: currencyUsecase,
	}

	srv.registerRoutes()

	return srv
}

func (s *Server) Run() error {
	log.Println("server running on: ", s.httpAddr)

	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	boletia := s.engine.Group("/boletia")
	boletia.GET("/currencies/:currency", handler.CurrencyGet(s.currencyUsecase))
}
