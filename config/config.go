package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Port        string `yaml:"port"`
	Env         string `yaml:"env"`
	Crt         string `yaml:"crt"`
	Key         string `yaml:"key"`
	StaticFiles string `yaml:"static_files"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type WechatConfig struct {
	AppID     string `yaml:"app_id"`
	AppSecret string `yaml:"app_secret"`
}

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Wechat   WechatConfig   `yaml:"wechat"`
}

var AppConfig Config

// LoadConfig 从 config.yaml 文件加载配置
func LoadConfig() {
	file, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}

	fmt.Println("Config file loaded successfully.")
}
