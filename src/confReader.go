package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Conf struct {
	Basedir   string `json:"basedir" yaml:"basedir"`
	Commitdir string `json:"commitdir" yaml:"commitdir"`
	Baseurl   string `json:"baseurl" yaml:"baseurl"`
	Commiturl string `json:"commiturl" yaml:"commiturl"`
	Username  string `json:"username" yaml:"username"`
	Password  string `json:"password" yaml:"password"`
	Commit    string `json:"commit"   yaml:"commit"`
}

func confReader() Conf {

	path := filepath.Join("conf", "config.yaml")
	cfg := Conf{}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return Conf{}
	}
	switch filepath.Ext(path) {
	case ".yaml":
		yaml.Unmarshal(data, &cfg)
	case ".json":
		json.Unmarshal(data, &cfg)
	default:
		fmt.Println("Error in Fetching Config file")
	}
	return cfg
}
