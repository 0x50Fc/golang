package xs

import (
	"errors"
	"log"
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
	p          *plugin.Plugin
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
	path, err := filepath.Abs(filepath.Join(filepath.Dir(configFile), dynamic.StringValue(dynamic.Get(config, "dylib"), "")))
	if err != nil {
		return nil, err
	}
	log.Println("[PLUGIN]", path)
	p, err := GetPlugin(path)
	if err != nil {
		return nil, err
	}
	log.Println("[PLUGIN]", path, "[OK]")
	return &App{config: config, p: p, mtime: st.ModTime(), configFile: configFile, env: env}, nil
}

func (A *App) GetConfigFile() string {
	return A.configFile
}

func (A *App) Valid() (bool, error) {

	st, err := os.Stat(A.configFile)

	if err != nil {
		return false, err
	}

	if A.mtime != st.ModTime() {
		config, err := GetConfigWithFileEnv(A.configFile, A.env)
		if err != nil {
			return false, err
		}
		p, err := GetPlugin(dynamic.StringValue(dynamic.Get(config, "dylib"), ""))
		if err != nil {
			return false, err
		}
		A.config = config
		A.p = p
		A.mtime = st.ModTime()
		return false, nil
	}

	return true, nil
}

func (A *App) In(s ...def.IService) error {

	v, err := A.p.Lookup("In")

	if err != nil {

		return err
	}

	fn, ok := v.(def.In)

	if !ok {
		return errors.New("In fail: " + A.configFile)
	}

	return fn(A.config, s...)

}
