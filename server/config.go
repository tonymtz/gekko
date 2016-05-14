package server

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
	"fmt"
)

type config struct {
	AppName      string `yaml:"app_name"`
	Port         int    `yaml:"http_port"`
	Database     string `yaml:"database"`
	GoogleId     string `yaml:"google_id"`
	GoogleSecret string `yaml:"google_secret"`
}

var Config *config

func init() {
	Config = getConfig()
}

func getConfig() *config {

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

	config := &config{}

	err = yaml.Unmarshal([]byte(dat), config)

	Check(err)

	return config
}
