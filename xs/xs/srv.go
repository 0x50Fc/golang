package xs

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/hailongz/golang/dynamic"
)

type Srv struct {
	config interface{}
	env    interface{}
	dir    string
	apps   map[string]*App
	lock   sync.RWMutex
}

func NewSrv(config interface{}) (*Srv, error) {

	apps := map[string]*App{}
	env := dynamic.Get(config, "env")
	dir := dynamic.StringValue(dynamic.Get(config, "dir"), ".")

	err := filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(p, ".json") {
			log.Println("[NEW]", p)
			app, err := NewApp(p, env)
			if err != nil {
				log.Println("[ERROR] [NEW]", p, err.Error())
				return err
			}
			log.Println("[NEW]", p, "DONE")
			apps[p] = app
			return nil
		}
		return nil
	})

	if err != nil {
		for _, app := range apps {
			app.Recycle()
		}
		return nil, err
	}

	return &Srv{config: config, apps: apps, env: env, dir: dir}, nil
}

func (S *Srv) Valid() {

	newApps := map[string]*App{}
	keys := map[string]bool{}
	rm := []string{}

	S.lock.RLock()

	filepath.Walk(S.dir, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(p, ".json") {
			app := S.apps[p]
			if app == nil {
				log.Println("[NEW]", p)
				app, err := NewApp(p, S.env)
				if err != nil {
					log.Println("[ERROR] [NEW]", p, err.Error())
					return nil
				}
				log.Println("[NEW]", p, "DONE")
				newApps[p] = app
				keys[p] = true
			} else {
				keys[p] = true
				_ = app.Valid()
			}
			return nil
		}
		return nil
	})

	for p, _ := range S.apps {
		if !keys[p] {
			rm = append(rm, p)
		}
	}

	S.lock.RUnlock()

	if len(rm) > 0 || len(newApps) > 0 {

		S.lock.Lock()

		for p, app := range newApps {
			S.apps[p] = app
		}

		for _, key := range rm {
			app := S.apps[key]
			if app != nil {
				app.Recycle()
				delete(S.apps, key)
			}
		}

		S.lock.Unlock()
	}

}

func (S *Srv) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	S.lock.RLock()
	defer S.lock.RUnlock()

	for _, app := range S.apps {
		if app.Handle(resp, req) {
			return
		}
	}

	resp.WriteHeader(http.StatusNotImplemented)
	resp.Write([]byte{})
}
