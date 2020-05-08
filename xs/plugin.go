package main

import (
	"errors"
	"plugin"

	"github.com/hailongz/golang/xs/def"
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

func In(p *plugin.Plugin, config interface{}, s ...def.IService) error {
	v, err := p.Lookup("In")
	if err != nil {
		return err
	}
	fn, ok := v.(def.In)
	if !ok {
		return errors.New("In fail")
	}
	fn(config, s...)
	return nil
}
