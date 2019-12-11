package config

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func GetConfig(section string, value string) string {
	cfg, err := ini.Load("/etc/keylime.conf")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	cfgvalue := cfg.Section(section).Key(value).String()
	return cfgvalue
}
