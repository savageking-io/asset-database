package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Database *DatabaseConfig `yaml:"database"`
	Rest     *RestConfig     `yaml:"rest"`
}

type DatabaseConfig struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Prefix   string `yaml:"prefix"`
}

type RestConfig struct {
	Port int `yaml:"port"`
}

func (c *Config) Init(filename string) error {
	log.Infof("Reading configuration from %s", filename)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Config not found: %s", err.Error())
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return fmt.Errorf("Couldn't parse yaml: %s", err.Error())
	}

	return nil
}
