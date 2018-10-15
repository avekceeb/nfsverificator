package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type TestConfig struct {
	ServerHost string `json:"server-host"`
	ServerPort int `json:"server-port"`
	ClientHost string `json:"client-host"`
	ClientId string `json:"client-id"`
}

func ReadConfig(configPath string) (config TestConfig) {
	jsonFile, err := os.Open(configPath)
	if err != nil {
		panic(err.Error())
	}
	defer jsonFile.Close()
	bytes, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(bytes, &config)
	return config
}