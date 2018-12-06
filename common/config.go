package common

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"flag"
	"strings"
	"time"
	"fmt"
	"os/exec"
	"bytes"
	"text/template"
	"strconv"
)

var (
	Config         TestConfig
	configFile     string
	serverOvr      string
	exportOvr      string
	traceOvr       string
	bkgCmd         *exec.Cmd
	funcMap        template.FuncMap
)

func init() {
	flag.StringVar(&configFile, "config", "", "Config File")
	flag.StringVar(&serverOvr, "server", "", "NFS Server")
	flag.StringVar(&exportOvr, "share", "", "NFS Share")
	flag.StringVar(&traceOvr, "trace", "", "Print Packets")
	funcMap = template.FuncMap{"timestamp":timestamp, "homedir":homedir}
}

type TestConfig struct {
	Server        string    `json:"server"`
	Port          int       `json:"port"`
	Export        string    `json:"export"`
	RebootCmd     string    `json:"reboot-cmd"`
	BkgCmd        string    `json:"background-cmd"`
	Trace         bool      `json:"trace"`
	SuiteName     string
}

func ParseOptions() (config TestConfig) {
	if configFile != "" {
		jsonFile, err := os.Open(configFile)
		if err != nil {
			panic(err.Error())
		}
		defer jsonFile.Close()
		bytes, err := ioutil.ReadAll(jsonFile)
		if nil != err {
			panic(err.Error())
		}
		json.Unmarshal(bytes, &config)
	}
	if "" != serverOvr {
		config.Server = serverOvr
	}
	if "" != exportOvr {
		config.Export = exportOvr
	}
	if "" != traceOvr {
		if t,e:=strconv.ParseBool("true"); e==nil {
			config.Trace = t
		}
	}
	if "" == config.Server {
		panic("No servers in config")
	}
	config.SuiteName = "default"
	return config
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

func (c *TestConfig) RebootServer() {
	cmdLine := c.RebootCmd
	if "" == cmdLine {
		return
	}
	cmdSlice := strings.Split(cmdLine, " ")
	cmdLine = cmdSlice[0]
	cmdSlice = cmdSlice[1:len(cmdSlice)]
	cmd := exec.Command(cmdLine, cmdSlice...)
	err := cmd.Run()
	if nil != err {
		fmt.Printf("Error running %s : %v", cmdLine, err)
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

//func Df() {
//	cmd := exec.Command("df", "-h")
//	out, _ := cmd.CombinedOutput()
//	fmt.Printf("\n>>>>>>>>>>>>>>\n%s\n>>>>>>>>>>>>>>\n", out)
//}
