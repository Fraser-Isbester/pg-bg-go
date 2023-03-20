package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func LoadConfig(filename string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

type Database struct {
	ConnectionString string `yaml:"connection_string"`
}

type LoadBalancer struct {
	Type   string             `yaml:"type"`
	Config LoadBalancerConfig `yaml:"config"`
}

type LoadBalancerConfig struct {
	Host          string `yaml:"host"`
	AdminUser     string `yaml:"admin_user"`
	AdminPassword string `yaml:"admin_password"`
	Port          int    `yaml:"port"`
}

type Config struct {
	DBs          map[string]Database `yaml:"dbs"`
	LoadBalancer LoadBalancer        `yaml:"load_balancer"`
}
