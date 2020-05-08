package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/xs/xs"
)

var httpSrv = &xs.HttpSrv{}
var apps = map[string]*xs.App{}

func reload(srv *http.Server, config interface{}) {
	updated := false
	vs := []*xs.App{}
	keys := map[string]bool{}
	env := dynamic.Get(config, "env")
	dir := dynamic.StringValue(dynamic.Get(config, "dir"), ".")

	filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(p, ".json") {
			app := apps[p]
			if app == nil {
				app, err = xs.NewApp(p, env)
				if err != nil {
					log.Println("[ERROR]", p, err)
					return nil
				}
				apps[p] = app
				vs = append(vs, app)
				updated = true
				keys[p] = true
				return nil
			}
			b, err := app.Valid()
			if err != nil {
				log.Println("[ERROR]", p, err)
				delete(apps, p)
				updated = true
				return nil
			}
			if !b {
				vs = append(vs, app)
				updated = true
				keys[p] = true
				return nil
			}
		}
		return nil
	})

	rm := []string{}

	for key, app := range apps {
		if !keys[key] {
			b, err := app.Valid()
			if err != nil {
				log.Println("[ERROR]", key, err)
				rm = append(rm, key)
				updated = true
				continue
			}
			if !b {
				vs = append(vs, app)
				updated = true
			}
		}
	}

	for _, key := range rm {
		delete(apps, key)
	}

	if httpSrv.IsEmpty() || updated {
		s := xs.NewHttpService()
		for _, v := range vs {
			err := v.In(s)
			if err != nil {
				log.Println("[ERROR]", v.GetConfigFile(), err)
			} else {
				log.Println("[DONE]", v.GetConfigFile())
			}
		}
		httpSrv.SetService(s)
	}
}

func main() {

	config, err := xs.GetConfig()

	if err != nil {
		log.Panicln(err)
	}

	srv := &http.Server{
		Addr:        dynamic.StringValue(dynamic.Get(config, "addr"), ":80"),
		IdleTimeout: time.Duration(dynamic.IntValue(dynamic.Get(config, "idle"), 6)) * time.Second,
		Handler:     httpSrv,
	}

	go func() {

		reload(srv, config)

		t := time.NewTicker(6 * time.Second)

		for {
			_, ok := <-t.C
			if !ok {
				break
			}
			reload(srv, config)
		}

	}()

	log.Println("[HTTPD]", srv.Addr)

	err = srv.ListenAndServe()

	if err != nil {
		log.Panicln(err)
	}
}
