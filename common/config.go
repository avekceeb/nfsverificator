package common

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"flag"
	"path/filepath"
)

var (
    Config TestConfig
    ConfigFile string

)

func init() {
    flag.StringVar(&ConfigFile, "config",
        filepath.Join(os.Getenv("GOPATH"),
        "src/github.com/avekceeb/nfsverificator/config.json"), "Config File")
}

type Server struct {
	Host          string    `json:"host"`
	Port          int       `json:"port"`
	ExportsRW     []string  `json:"exports-rw"`
	ExportsRO     []string  `json:"exports-ro"`
	ExportsBlock  []string  `json:"exports-block-layout"`
	RebootCmd     string    `json:"reboot-cmd"`
	// these would be obtained by requests to server itself
	//LeaseTime  int ?? (in client)
}

type TestConfig struct {
	DefaultServer string            `json:"default-server"`
	Servers       map[string]Server `json:"servers"`
}

func ReadConfig(configPath string) (config TestConfig) {
	jsonFile, err := os.Open(configPath)
	if err != nil {
		panic(err.Error())
	}
	defer jsonFile.Close()
	bytes, err := ioutil.ReadAll(jsonFile)
	if nil != err {
		panic(err.Error())
	}
	json.Unmarshal(bytes, &config)
	if len(config.Servers) < 1 {
		panic("No servers in config")
	}
	if ! config.SetDefaultServer(config.DefaultServer) {
		panic("Wrong default server specified")
	}
	// TODO: not mandatory ?
	if len(config.Servers[config.DefaultServer].ExportsRW) < 1 {
		panic("No rw exports specified")
	}
	return config
}

func (c *TestConfig) SetDefaultServer(srv string) (bool) {
	_, exists := c.Servers[srv]
	if exists {
		c.DefaultServer = srv
	}
	return exists
}

func (c *TestConfig) GetHost() string {
	return c.Servers[c.DefaultServer].Host
}

func (c* TestConfig) GetPort() int {
	return c.Servers[c.DefaultServer].Port
}

func (c *TestConfig) GetRWExport() string {
	return c.Servers[c.DefaultServer].ExportsRW[0]
}

func (c *TestConfig) GetBlockExport() string {
	if 0 != len(c.Servers[c.DefaultServer].ExportsBlock) {
		return c.Servers[c.DefaultServer].ExportsBlock[0]
	} else {
		return ""
	}
}
