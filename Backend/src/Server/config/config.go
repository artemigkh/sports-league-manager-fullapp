package config

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config interface {
	GetDbConnString() string
	GetPortString() string
	GetIconsDir() string
	GetMarkdownDir() string
	GetKeys() ([]byte, []byte)

	GetLeagueOfLegendsApiKey() string
}

type Configuration struct {
	DbUser        string `json:"dbUser"`
	DbPass        string `json:"dbPass"`
	DbName        string `json:"dbName"`
	Port          string `json:"port"`
	IconsDir      string `json:"iconsDir"`
	MarkdownDir   string `json:"markdownDir"`
	AuthKey       string `json:"authKey"`
	EncryptionKey string `json:"encryptionKey"`

	LeagueOfLegendsApiKey string `json:"leagueOfLegendsApiKey"`
}

func (c *Configuration) GetDbConnString() string {
	return fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", c.DbUser, c.DbPass, c.DbName)
}

func (c *Configuration) GetPortString() string {
	return fmt.Sprintf("0.0.0.0:%v", c.Port)
}

func (c *Configuration) GetIconsDir() string {
	return c.IconsDir
}

func (c *Configuration) GetMarkdownDir() string {
	return c.MarkdownDir
}

func (c *Configuration) GetKeys() ([]byte, []byte) {
	authKey, _ := hex.DecodeString(c.AuthKey)
	encryptionKey, _ := hex.DecodeString(c.EncryptionKey)
	return authKey, encryptionKey
}

func (c *Configuration) GetLeagueOfLegendsApiKey() string {
	return c.LeagueOfLegendsApiKey
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
