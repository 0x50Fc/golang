package main

import (
	"log"
	"net/http"
	"time"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/xs/def"
)

var httpService def.IHTTPService = nil

func reload(srv *http.Server, config interface{}) {
	s := NewHttpService()
	env := dynamic.Get(config, "env")

	dynamic.Each(dynamic.Get(config, "app"), func(_ interface{}, item interface{}) bool {

		configFile := dynamic.StringValue(dynamic.Get(item, "config"), "")

		if configFile == "" {
			log.Println("[ERROR] [ConfigFile] config", item)
			return true
		}

		dylib := dynamic.StringValue(dynamic.Get(item, "dylib"), "")

		if dylib == "" {
			log.Println("[ERROR] [ConfigFile] dylib", item)
			return true
		}

		cfg, err := GetConfigWithFileEnv(configFile, env)

		if err != nil {
			log.Println("[ERROR] [ConfigFile]", item, err)
			return true
		}

		p, err := GetPlugin(dylib)

		if err != nil {
			log.Println("[ERROR] [ConfigFile] GetPlugin", item, err)
			return true
		}

		err = In(p, cfg, s)

		if err != nil {
			log.Println("[ERROR] [ConfigFile] In", item, err)
			return true
		}

		return true
	})

	httpService = s
	srv.Handler = s.Mux
}

func main() {

	config, err := GetConfig()

	if err != nil {
		log.Panicln(err)
	}

	srv := &http.Server{
		Addr:        dynamic.StringValue(dynamic.Get(config, "addr"), ":80"),
		IdleTimeout: time.Duration(dynamic.IntValue(dynamic.Get(config, "idle"), 6)) * time.Second,
	}

	reload(srv, config)

	err = srv.ListenAndServe()

	if err != nil {
		log.Panicln(err)
	}
}
