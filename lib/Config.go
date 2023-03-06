package lib

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// 配置总结构体
type Config struct {
	MacAddr                    string `yaml:"MacAddr"`
	ProcessStatusCheckInterval int64  `yaml:"ProcessStatusCheckInterval"`
	MainProcessPingInterval    int64  `yaml:"MainProcessPingInterval"`

	Mqtt struct {
		Host         string `yaml:"Host"`
		ClientId     string `yaml:"ClientId"`
		Username     string `yaml:"Username"`
		Password     string `yaml:"Password"`
		CleanSession bool   `yaml:"CleanSession"`
	} `yaml:"Mqtt"`

	Apps []App `yaml:"Apps"`
}

// 程序结构体
type App struct {
	Name         string `yaml:"Name"`         // 应用名称
	ProcessName  string `yaml:"ProcessName"`  // 进程名称
	ShortcutName string `yaml:"ShortcutName"` // 快捷方式名称
	UniqueId     string `yaml:"UniqueId"`     // 不允许重复的唯一识别
}

func LoadConfig() (*Config, error) {

	var config Config

	appConfig, err := ioutil.ReadFile("./config/config.yaml")

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(appConfig, &config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}
