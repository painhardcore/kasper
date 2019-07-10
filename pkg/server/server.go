package server

import (
	"fmt"
	"github.com/painhardcore/kasper/pkg/counter"
	"github.com/painhardcore/kasper/pkg/counter/file"
	"log"
	"net/http"
	"time"
)

type Server struct {
	counter counter.Counter
}

func (s *Server) counterMW(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r)
		err := s.counter.Inc()
		if err != nil {
			log.Printf("error while increasing counter : %s", err)
		}
	}
}

func (s *Server) printCount(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, s.counter)
}

func Start(filename string, duration time.Duration) {
	fileStore, err := file.New(filename, duration)
	if err != nil {
		log.Fatalf("Can't initialize store for counting requests: %s", err)
	}
	srv := Server{fileStore}
	http.HandleFunc("/", srv.counterMW(srv.printCount))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
