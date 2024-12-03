package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func Read(configPath string) (c Config, err error) {
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		err = fmt.Errorf("%s: %w", "Error reading config file", err)
		return
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		err = fmt.Errorf("%s: %w", "Error parsing config file", err)
	}
	return
}

type Config struct {
	Server struct {
		Port         string `yaml:"port"`
		ReadTimeout  int    `yaml:"readTimeout"`
		WriteTimeout int    `yaml:"writeTimeout"`
	} `yaml:"server"`
	PSQL struct {
		DSN string `yaml:"dsn"`
	} `yaml:"psql"`
	Bucket struct {
		IPLimit       int `yaml:"ipLimit"`
		LoginLimit    int `yaml:"loginLimit"`
		PasswordLimit int `yaml:"passwordLimit"`
		BucketTTL     int `yaml:"bucketTtl"`
	} `yaml:"bucket"`
}
