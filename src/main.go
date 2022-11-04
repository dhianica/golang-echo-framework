package main

import (
	"log"
	"net/http"
	"time"

	router "golang-echo-framework/src/app"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	e := router.New()

	h2s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}
	s := http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(e, h2s),
	}
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
