package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config interface {
	GetDbConnString() string
	GetPortString() string
}

type Configuration struct {
	DbUser string `json:"dbUser"`
	DbPass string `json:"dbPass"`
	DbName string `json:"dbName"`
	Port   string `json:"port"`
}

func (c *Configuration) GetDbConnString() string {
	return fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", c.DbUser, c.DbPass, c.DbName)
}

func (c *Configuration) GetPortString() string {
	return fmt.Sprintf("0.0.0.0:%v", c.Port)
}

func GetConfig(location string) Config {
	file, err := os.Open(location)
	if err != nil {
		log.Fatal("error opening config: ", err)
		return nil
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var config Configuration
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("error decoding config: ", err)
		return nil
	}
	return &config
}