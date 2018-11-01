package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)

type TestConfig struct {
	ServerHost string    `json:"server-host"`
	ServerPort int       `json:"server-port"`
	Exports    []string  `json:"exports"`
	ExportsRO  []string  `json:"exports-ro"`
}

func init () {
    rand.Seed(time.Now().Unix())
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

func RandInt(min int, max int) int {
    return rand.Intn(max - min) + min
}

const letters = "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"

func RandString(n int) string {
    var l int64 = int64(len(letters))
    b := make([]byte, n)
    for i := 0; i < n; i++ {
        b[i] = letters[int(rand.Int63n(l))]
    }
    return string(b)
}

