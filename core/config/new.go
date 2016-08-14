package config

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"

	"github.com/tonymtz/gekko/utils"
)

// New ...
func New(filePath string) *Config {

	var dat []byte
	var err error

	if _, err = os.Stat(filePath); !os.IsNotExist(err) {
		dat, err = ioutil.ReadFile(filePath)
	}

	utils.Check(err)

	config := &Config{}

	err = yaml.Unmarshal([]byte(dat), config)

	utils.Check(err)

	return config
}
