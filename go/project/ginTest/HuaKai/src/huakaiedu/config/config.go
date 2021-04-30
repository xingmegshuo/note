/***************************
@File        : config.go
@Time        : 2021/04/20 15:16:33
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : config 配置
****************************/

package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

type Config struct {
	Sql          string
	Port         string
	PageSize     int
	SuperAccount string
	Password     string
}

var Con *Config
var lock sync.RWMutex

func init() {
	con := &Config{}
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(data, con)
	if err != nil {
		log.Panic("error:", err)
	}
	lock.Lock()
	Con = con
	lock.Unlock()
}
