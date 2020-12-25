package config

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
	"os"
)

type ServerConfig struct {
	ClientID string
	ClientSecret string
	RedirectURL string
	AuthURL string
	TokenURL string
	Listen string
}

func ReadConfig(file string) ServerConfig {
	var config ServerConfig
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	return config
}
