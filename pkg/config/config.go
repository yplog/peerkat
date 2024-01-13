package config

import (
	"fmt"
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

func (c *Config) String() string {
	configString, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	return string(configString)
}

func Exists() bool {
	configFilePath, err := Path()
	if err != nil {
		return false
	}

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func Setup() {
	configFileExists := Exists()
	if !configFileExists {
		log.Println("Config file not found, generating...")
		generate, err := Generate()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Config file generated at", generate)
		fmt.Println("")
	}
}

func Read() (*Config, error) {
	configFilePath, err := Path()
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

	return &config, nil
}

func Generate() (string, error) {
	config := Config{
		Server: ServerConfig{
			Port: 8080,
			Host: "localhost",
		},
	}

	yamlConfig, err := yaml.Marshal(&config)
	if err != nil {
		return "", err
	}

	configFilePath, err := Path()
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(filepath.Dir(configFilePath), 0755)
	if err != nil {
		return "", err
	}

	err = os.WriteFile(configFilePath, yamlConfig, 0644)
	if err != nil {
		return "", err
	}

	return configFilePath, nil
}

func Path() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(homeDir, ".config", "peerkat")
	configFile := filepath.Join(configDir, "config.yaml")

	return configFile, nil
}
