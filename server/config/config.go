package config

import (
	"io/ioutil"
	"os"
	"fmt"

	"gopkg.in/yaml.v2"

	"github.com/tonymtz/gekko/server/utils"
)

type config struct {
	AppName          string `yaml:"app_name"`
	Port             int    `yaml:"http_port"`

	User             string `yaml:"user"`
	Password         string `yaml:"password"`
	Database         string `yaml:"database"`

	GoogleId         string `yaml:"google_id"`
	GoogleSecret     string `yaml:"google_secret"`
	GoogleCallback   string `yaml:"google_callback"`

	DropboxId        string `yaml:"dropbox_id"`
	DropboxSecret    string `yaml:"dropbox_secret"`
	DropboxCallback  string `yaml:"dropbox_callback"`
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

	utils.Check(err)

	config := &config{}

	err = yaml.Unmarshal([]byte(dat), config)

	utils.Check(err)

	return config
}
