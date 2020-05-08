package xs

import (
	"plugin"
)

var pluginMap = map[string]*plugin.Plugin{}

func GetPlugin(path string) (*plugin.Plugin, error) {
	v := pluginMap[path]
	if v == nil {
		var err error
		v, err = plugin.Open(path)
		if err != nil {
			return nil, err
		}
		pluginMap[path] = v
	}
	return v, nil
}
