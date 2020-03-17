package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
	"github.com/hailongz/golang/stat"
	"github.com/hailongz/golang/svc"
	_ "github.com/hailongz/pj_bitcoin/srv/ms/geetest/geetest"
)

func main() {

	s := svc.New()

	app, err := micro.NewAppWithEnv()

	if err != nil {
		log.Panicln(err)
	}

	address := dynamic.StringValue(dynamic.GetWithKeys(app.GetConfig(), []string{"httpd", "addr"}), ":80")

	http.HandleFunc("/"+app.GetName()+"/", micro.HandleFunc(app, s))

	http.HandleFunc("/", stat.HandleFunc())

	log.Println("[HTTPD]", address)

	srv := &http.Server{
		Addr:        address,
		IdleTimeout: 6 * time.Second,
	}

	{
		go func() {

			ch := make(chan os.Signal, 1)

			signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

			log.Println("[SIGNAL]", <-ch)

			signal.Stop(ch)

			s.Done()

			log.Println("[SERVER] [DONE]")

			srv.Close()

			log.Println("[SERVER] [CLOSE]")
		}()

	}

	log.Println(srv.ListenAndServe())

}