package common

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"flag"
	"path/filepath"
	"strings"
	"time"
	"fmt"
	"os/exec"
	"bytes"
	"text/template"
)

var (
    Config     TestConfig
    ConfigFile string
	bkgCmd     *exec.Cmd
	funcMap    template.FuncMap
)

func init() {
    flag.StringVar(&ConfigFile, "config",
        filepath.Join(os.Getenv("GOPATH"),
        "src/github.com/avekceeb/nfsverificator/config.json"), "Config File")
	funcMap = template.FuncMap{"timestamp":timestamp, "homedir":homedir}
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
	BkgCmd        string            `json:"background-cmd"`
	SuiteName     string
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
	config.SuiteName = "default"
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

//////// external background commands ////////////////

func (c *TestConfig) RunExternalCommands() {
	if "" != c.BkgCmd {
		t, err := template.New("command").Funcs(funcMap).Parse(c.BkgCmd)
		if nil != err {
			fmt.Println("executing template:", err)
			return
		}
		buf := new(bytes.Buffer)
		err = t.Execute(buf, *c)
		if err != nil {
			fmt.Println("executing template:", err)
			return
		}
		c.BkgCmd = buf.String()
		cmdList := strings.Split(c.BkgCmd, " ")
		cmd := cmdList[0]
		cmdList = cmdList[1:len(cmdList)]
		bkgCmd = exec.Command(cmd, cmdList...)
		bkgCmd.Start()
		time.Sleep(time.Second)
	}
}

func (c *TestConfig) StopExternalCommands() {
	if nil != bkgCmd {
		if nil != bkgCmd.Process {
			fmt.Println("\nGiving 3 seconds for", c.BkgCmd, "to finish...")
			time.Sleep(time.Second*3)
			//bkgCmd.Process.Kill()
			bkgCmd.Process.Signal(os.Interrupt)
		}
	}

}

///////////////////////////////////////////////////////

func timestamp() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02d_%02d.%02d.%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}

func homedir() string {
	return os.Getenv("HOME")
}
