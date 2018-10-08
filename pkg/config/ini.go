package config

import (
	"fmt"
	"os"

	"github.com/go-ini/ini"
	"github.com/hungranger/Go_Resume_ChatBot/pkg/helper"
)

type IniConfig struct {
	Port         string `ini:"port"`
	Facebook_api string `ini:"facebook_api"`
	Image        string `ini:"image"`
}

func (c *IniConfig) LoadConfig() IConfig {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		cfg, err = ini.Load(helper.GetRootPath() + "/config.ini")
		if err != nil {
			fmt.Printf("Fail to read file: %v", err)
			os.Exit(1)
		}
	}

	err = cfg.MapTo(c)
	helper.CheckError(err)

	return c
}

func (c IniConfig) GetPort() string {
	return c.Port
}

func (c IniConfig) GetFbApi() string {
	return c.Facebook_api
}

func (c IniConfig) GetImage() string {
	return c.Image
}
