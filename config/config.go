package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	Db DB
	Ela ELA
	Btc BTC
}

type DB struct {
	DbDriverName   string
	DbDriverSource string
}

type ELA struct {
	Host       string
	ServerPort string
}

type BTC struct {
	Host      string
	Rpcuser   string
	Rpcpasswd string
}

var Conf *config

func init() {
	Conf = new(config)
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&Conf)
	if err != nil {
		fmt.Println("Error init Config :", err.Error())
		os.Exit(-1)
	}
}
