package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"
)

type configuration struct {
	ServerInfo serverInfo
	RedisInfo  redisInfo
}

type serverInfo struct {
	Host string
}

type redisInfo struct {
	Host        string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var Configuration = configuration{}

func init() {
	filePath := path.Join(os.Getenv("GOPATH"), "src/go-chat/config/config.json")
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("open file error: %v\n", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	// Configuration = configuration{}
	err = decoder.Decode(&Configuration)
	fmt.Printf("Configuration: %v\n", Configuration.RedisInfo)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}