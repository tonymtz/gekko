package server

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
	"fmt"
)

type Config struct {
	AppName  string `yaml:"appname"`
	Port     int    `yaml:"httpport"`
	Database string `yaml:"database"`
}

var Config *Config

func init() {
	Config = getConfig()
}

func getConfig() *Config {

	env := os.Getenv("GEKKO_ENV")

	var (
		dat []byte
		err error
	)

	if env == "" {
		dat, err = ioutil.ReadFile("config/env.conf.sample")
	} else {
		filePath := fmt.Sprintf("config/%v.conf", env)

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			dat, err = ioutil.ReadFile("config/env.conf.sample")
		} else {
			dat, err = ioutil.ReadFile(filePath)
		}
	}

	Check(err)

	config := &Config{}

	err = yaml.Unmarshal([]byte(dat), config)

	fmt.Println(config)

	Check(err)

	return config
}
