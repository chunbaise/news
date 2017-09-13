package config

import (
	"io/ioutil"
	"log"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type MySQL struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DbIndex  string `yaml:"dbindex"`
}
type Conf struct {
	MySQL
	Redis
}

var C = new(Conf)

var confiAbsPath = "config/config.yaml"

func init() {
	strProjectPath, _ := filepath.Abs("..")
	yamlFile, err := ioutil.ReadFile(filepath.Join(strProjectPath, "/", confiAbsPath))

	if err != nil {
		log.Printf("YamlFile Parsing Failed: #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, C)
	if err != nil {
		log.Fatalf("Yaml Unmarshal Failed: %v", err)
	}
}
