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
	// TODO: save and replay seed
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

func CheckFlag(flags uint32, flag int) bool {
	return 1 == (flags & uint32(flag))
}

func BytesToUint32(b []byte) uint32 {
	r := uint32(0)
	for i:=range b {
		r += uint32(b[i])
	}
	return r
}

func MakeGetAttrFlags(f ...int) uint32 {
	r := uint32(0)
	for i:=range f{
		r |= (1<<uint32(f[i])) // ??
	}
	return r
}

func MakeUint32Flags(f ...int) uint32 {
	r := uint32(0)
	for i:=range f {
		r |= uint32(f[i])
	}
	return r
}

