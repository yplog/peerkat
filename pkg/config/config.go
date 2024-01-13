package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

func path() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(homeDir, ".config", "peerkat")
	configFile := filepath.Join(configDir, "config.yaml")

	return configFile, nil
}

func Exists() bool {
	configFilePath, err := path()
	if err != nil {
		return false
	}

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func Generate() bool {
	config := Config{
		Server: ServerConfig{
			Port: 8080,
			Host: "localhost",
		},
	}

	yamlConfig, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatal(err)
		return false
	}

	configFilePath, err := path()
	if err != nil {
		log.Fatal(err)
		return false
	}

	err = os.MkdirAll(filepath.Dir(configFilePath), 0755)
	if err != nil {
		log.Fatal(err)
		return false
	}

	err = os.WriteFile(configFilePath, yamlConfig, 0644)
	if err != nil {
		log.Fatal(err)
		return false
	}

	log.Printf("Config file generated at %s", configFilePath)

	return true
}

func Read() (*Config, error) {
	configFilePath, err := path()
	if err != nil {
		log.Println("Config file path error")
		return nil, err
	}

	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Println("File read error")
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Println("Unmarshal error")
		return nil, err
	}

	log.Printf("Config file read from %s", configFilePath)
	log.Printf("Config: %+v", config)

	return &config, nil
}
