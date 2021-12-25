package configs

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	DB     DBConfig     `yaml:"db"`
	Server ServerConfig `yaml:"server"`
	Logger LoggerConfig `yaml:"logger"`
}

type DBConfig struct {
	DatabaseURL string `yaml:"database_url"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Port        string `yaml:"port"`
}

type ServerConfig struct {
	BindAddr string `yaml:"bind_addr"`
}

type LoggerConfig struct {
	Level string `yaml:"level"`
}

func InitConfig(configFilePath string) (*Config, error) {
	c := &Config{}
	yamlFile, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		logrus.Fatalf("yamlFile.Get err   #%v ", err)
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		logrus.Fatalf("Unmarshal: %v", err)
		return nil, err
	}

	return c, nil
}
