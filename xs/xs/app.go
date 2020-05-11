package xs

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"plugin"
	"time"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/xs/def"
)

type App struct {
	env        interface{}
	config     interface{}
	mtime      time.Time
	configFile string
	pluginFile string
	p          *plugin.Plugin
	s          *Service
}

func In(p *plugin.Plugin, config interface{}, s def.IService) error {

	v, err := p.Lookup("In")

	if err != nil {
		return err
	}

	fn, ok := v.(def.In)

	if !ok {
		return errors.New("In fail")
	}

	return fn(config, s)
}

func NewApp(configFile string, env interface{}) (*App, error) {
	st, err := os.Stat(configFile)
	if err != nil {
		return nil, err
	}
	config, err := GetConfigWithFileEnv(configFile, env)
	if err != nil {
		return nil, err
	}

	dir := filepath.Dir(configFile)

	dynamic.Set(config, "dir", dir)

	path, err := filepath.Abs(filepath.Join(dir, dynamic.StringValue(dynamic.Get(config, "dylib"), "")))
	if err != nil {
		return nil, err
	}
	p, err := GetPlugin(path)
	if err != nil {
		return nil, err
	}
	s := NewService()
	err = In(p, config, s)
	if err != nil {
		return nil, err
	}
	return &App{config: config, p: p, mtime: st.ModTime(), configFile: configFile, pluginFile: path, env: env, s: s}, nil
}

func (A *App) GetConfigFile() string {
	return A.configFile
}

func (A *App) GetPluginFile() string {
	return A.pluginFile
}

func (A *App) Valid() error {

	st, err := os.Stat(A.configFile)

	if err != nil {
		return err
	}

	if A.mtime != st.ModTime() {

		log.Println("[UPDATE]", A.configFile)

		config, err := GetConfigWithFileEnv(A.configFile, A.env)

		if err != nil {
			log.Println("[ERROR] [UPDATE]", A.configFile, err)
			return err
		}

		dir := filepath.Dir(A.configFile)

		dynamic.Set(config, "dir", dir)

		path, err := filepath.Abs(filepath.Join(dir, dynamic.StringValue(dynamic.Get(config, "dylib"), "")))

		if err != nil {
			log.Println("[ERROR] [UPDATE]", A.configFile, err)
			return err
		}

		p, err := GetPlugin(path)

		if err != nil {
			log.Println("[ERROR] [UPDATE]", A.configFile, err)
			return err
		}

		s := NewService()
		err = In(p, config, s)

		if err != nil {
			log.Println("[ERROR] [UPDATE]", A.configFile, err)
			return err
		}

		v := A.s
		A.s = s
		A.config = config
		A.pluginFile = path
		A.mtime = st.ModTime()
		A.p = p

		v.Recycle()

		log.Println("[UPDATE]", A.configFile, "DONE")
	}

	return nil
}

func (A *App) Recycle() {
	if A.s != nil {
		A.s.Recycle()
		A.s = nil
	}
}

func (A *App) Handle(resp http.ResponseWriter, req *http.Request) bool {
	if A.s == nil {
		return false
	}
	return A.s.Handle(resp, req)
}
