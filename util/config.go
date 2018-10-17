package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type TestConfig struct {
	ServerHost string    `json:"server-host"`
	ServerPort int       `json:"server-port"`
	Exports    []string  `json:"exports"`
	ExportsRO  []string  `json:"exports-ro"`
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