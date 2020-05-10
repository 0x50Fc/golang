package main

import (
	"log"
	"net/http"
	"time"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/xs/xs"
)

func main() {

	config, err := xs.GetConfig()

	if err != nil {
		log.Panicln(err)
	}

	s, err := xs.NewSrv(config)

	if err != nil {
		log.Panicln(err)
	}

	srv := &http.Server{
		Addr:        dynamic.StringValue(dynamic.Get(config, "addr"), ":80"),
		IdleTimeout: time.Duration(dynamic.IntValue(dynamic.Get(config, "idle"), 6)) * time.Second,
		Handler:     s,
	}

	go func() {

		t := time.NewTicker(6 * time.Second)

		for {
			_, ok := <-t.C
			if !ok {
				break
			}
			s.Valid()
		}

	}()

	log.Println("[HTTPD]", srv.Addr)

	err = srv.ListenAndServe()

	if err != nil {
		log.Panicln(err)
	}
}
