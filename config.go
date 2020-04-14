package main

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type conf struct {
	Ucloud     ucloudConf     `yaml:"ucloud"`
	Cloudflare cloudflareConf `yaml:"cloudflare"`
	Instances  []string       `yaml:"instances"`
}

type ucloudConf struct {
	ProjectID  string `yaml:"project_id"`
	PublicKey  string `yaml:"public_key"`
	PrivateKey string `yaml:"private_key"`
}

type cloudflareConf struct {
	Token  string `yaml:"token"`
	Zone   string `yaml:"zone"`
	Record string `yaml:"record"`
}

var config = &conf{}

func (c *conf) getConf() {
	if f, err := os.Open("config.yaml"); err != nil {
		panic(err)
	} else {
		yaml.NewDecoder(f).Decode(c)
	}
}

func (c *conf) save() {
	data, _ := yaml.Marshal(c)
	err := ioutil.WriteFile("config.yaml", data, 0644)
	if err != nil {
		panic(err)
	}
}
