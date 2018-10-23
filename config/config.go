package config

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/kataras/golog"
)

var (
	Friends       []string
	Port          int
	LogPath       string
	LogName       string
	LogFileExpire time.Duration
	LogFileSplit  time.Duration
	Name          string
)

func init() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		golog.Fatal(err)
	}
	var config struct {
		Name          string   `json:"name"`
		Port          int      `json:"port"`
		Friends       []string `json:"sendto"`
		LogPath       string   `json:"logpath"`
		LogName       string   `json:"logname"`
		LogFileExpire int      `json:"logfileexpire"`
		LogFileSplit  int      `json:"logfilesplit"`
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		golog.Error(err)
	}
	Port = config.Port
	Friends = config.Friends
	LogPath = config.LogPath
	LogName = config.LogName
	LogFileExpire = time.Duration(24*config.LogFileExpire) * time.Hour
	LogFileSplit = time.Duration(config.LogFileSplit) * time.Hour
	Name = config.Name
	if LogName == "" || LogPath == "" {
		golog.Fatal("invalid LogPath or LogName")
	}
	if Name == "" || Name == "local" {
		golog.Fatal("please set a name. tip: dont use \"local\"")
	}
}
