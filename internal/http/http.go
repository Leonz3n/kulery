package http

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"

	"github.com/Leonz3n/k8s-job-massage/internal/base"
)

type Server struct {
	httpsvr *http.Server

	logger base.Logger

	// Port http server listen port
	broker base.Broker
}

// NewServer new an HTTP server.
func NewServer(logger base.Logger, broker base.Broker) *Server {
	return &Server{
		httpsvr: nil,
		logger:  logger,
		broker:  broker,
	}
}

// Start listen and start.
func (s *Server) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		router := httprouter.New()
		router.GET("/hello", s.Hello)
		s.httpsvr = &http.Server{Addr: ":8080", Handler: router}

		if err := s.httpsvr.ListenAndServe(); err != nil {
			s.logger.Fatal("While serving HTTP: %s", err.Error())
		}
	}()
}

func (s *Server) Stop() {

}

// Hello world!
func (s *Server) Hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := fmt.Fprintf(w, "hello, k8s-job-massage!\n")
	if err != nil {
		return
	}
}

// CreateJob create a job.
func (s *Server) CreateJob(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := fmt.Fprintf(w, "hello, k8s-job-massage!\n")
	if err != nil {
		return
	}
}
