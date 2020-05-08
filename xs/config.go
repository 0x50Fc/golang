package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/hailongz/golang/dynamic"
)

func GetConfig() (interface{}, error) {

	configFile := "./app.json"

	{
		s := os.Getenv("KK_CONFIG_FILE")
		if s != "" {
			configFile = s
			log.Println("[KK_CONFIG_FILE]", s)
		}
	}

	config, err := GetConfigWithFile(configFile)

	if err != nil {
		return nil, err
	}
	return config, nil
}

func GetFileContent(configFile string) ([]byte, error) {

	fd, err := os.Open(configFile)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	b, err := ioutil.ReadAll(fd)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func GetConfigWithFile(configFile string) (interface{}, error) {

	var config interface{} = nil

	b, err := GetFileContent(configFile)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &config)

	if err != nil {
		return nil, err
	}

	return config, nil
}

func GetConfigWithFileEnv(configFile string, env interface{}) (interface{}, error) {

	var config interface{} = nil

	b, err := GetFileContent(configFile)

	if err != nil {
		return nil, err
	}

	s := string(b)

	dynamic.Each(env, func(key interface{}, value interface{}) bool {
		s = strings.ReplaceAll(s, dynamic.StringValue(key, ""), dynamic.StringValue(value, ""))
		return true
	})

	err = json.Unmarshal([]byte(s), &config)

	if err != nil {
		return nil, err
	}

	return config, nil
}
