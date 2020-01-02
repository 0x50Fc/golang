package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/mq"
	_ "github.com/hailongz/golang/mq/ali"
	_ "github.com/hailongz/golang/mq/nsq"
	less "github.com/hailongz/golang/serverless/app"
	_ "github.com/hailongz/golang/serverless/base"
	_ "github.com/hailongz/golang/serverless/db"
	_ "github.com/hailongz/golang/serverless/geoip"
	_ "github.com/hailongz/golang/serverless/influx"
	_ "github.com/hailongz/golang/serverless/mq"
	_ "github.com/hailongz/golang/serverless/redis"
	_ "github.com/hailongz/golang/serverless/view"
	"github.com/hailongz/golang/stat"
	"github.com/hailongz/golang/svc"
	"github.com/hailongz/golang/text"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	app, err := less.NewAppWithEnv()

	if err != nil {
		log.Panicln(err)
	}

	dynamic.Each(dynamic.Get(app.GetConfig(), "fonts"), func(_ interface{}, p interface{}) bool {
		text.Load(dynamic.StringValue(p, ""))
		return true
	})

	log.Println("[Font]", strings.Join(text.GetFamilys(), ","))

	block := false

	{
		init := dynamic.Get(app.GetConfig(), "init")

		if init != nil {

			input := less.Input{
				Path: "__init__",
				Data: init,
			}

			input.Trace = fmt.Sprintf("%d", app.NewID())

			log.Printf("[%s] [INIT]\n", input.Trace)

			_, err := app.Exec("/main.js", &input)

			if err != nil {
				log.Printf("[%s] [__init__] [ERROR] %s\n", input.Trace, err)
			} else {
				log.Printf("[%s] [__init__] [OK]\n", input.Trace)
			}

		}
	}

	{
		cfg := dynamic.GetWithKeys(app.GetConfig(), []string{"mq", "consumer"})

		if cfg != nil {

			stype := dynamic.StringValue(dynamic.Get(cfg, "type"), "")

			if stype != "" {

				q, err := mq.OpenConsumer(stype, cfg)

				if err != nil {
					log.Panicln(err)
				}

				defer q.Close()

				log.Println("[MQ]", dynamic.Get(cfg, "addr"))

				err = q.Open(func(name string, data interface{}) error {

					input := less.Input{
						Path: name,
						Data: data,
					}

					input.Trace = fmt.Sprintf("%d", app.NewID())

					log.Printf("[%s] [MQ] %s %s\n", input.Trace, name, data)

					_, err := app.Exec("/main.js", &input)

					if err != nil {
						log.Printf("[%s] [MQ] [ERROR] %s %s\n", input.Trace, name, err)
						return err
					}

					log.Printf("[%s] [MQ] [OK] %s\n", input.Trace, name)

					return nil

				}, int(dynamic.IntValue(dynamic.Get(cfg, "concurrency"), 1)))

				if err != nil {
					log.Panicln(err)
				}

				block = true

			}
		}
	}

	{
		collector := dynamic.Get(app.GetConfig(), "collector")

		if collector != nil {

			dynamic.Each(collector, func(key interface{}, value interface{}) bool {

				input := less.Input{
					Path: dynamic.StringValue(dynamic.Get(value, "name"), ""),
					Data: dynamic.Get(value, "data"),
				}

				tv := dynamic.IntValue(dynamic.Get(value, "interval"), 6000)

				go func(input *less.Input, tv time.Duration) {

					t := time.NewTicker(tv)

					for {

						input.Trace = fmt.Sprintf("%d", app.NewID())

						log.Printf("[%s] [COLLECTOR] %s %s\n", input.Trace, input.Path, input.Data)

						_, err := app.Exec("/main.js", input)

						if err != nil {
							log.Printf("[%s] [COLLECTOR] [ERROR] %s %s\n", input.Trace, input.Path, err)
						} else {
							log.Printf("[%s] [COLLECTOR] [OK] %s\n", input.Trace, input.Path)
						}

						<-t.C
					}

				}(&input, time.Duration(tv)*time.Millisecond)

				return true
			})

			block = true
		}
	}

	if dynamic.Get(app.GetConfig(), "httpd") != nil {

		s := svc.New()

		address := dynamic.StringValue(dynamic.GetWithKeys(app.GetConfig(), []string{"httpd", "addr"}), ":80")

		http.HandleFunc("/__stat", stat.HandleFunc())

		http.HandleFunc("/", less.HandleFunc(app, s))

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

	} else if block {
		select {}
	} else {

		input := less.Input{
			Path: "main.js",
			Data: os.Args,
		}

		_, err := app.Exec("/main.js", &input)

		if err != nil {
			log.Println(err)
		}

	}

}
