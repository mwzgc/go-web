package config

import (
	"fmt"

	"github.com/Unknwon/goconfig"
)

// GetConfig ...
func GetConfig() (*goconfig.ConfigFile, error) {
	configName := "conf.ini"
	cfg, err := goconfig.LoadConfigFile(configName)
	if err != nil {
		path := ""
		for i := 0; i < 5; i++ {
			path = path + "../"
			fmt.Println(path + configName)
			cfg, err = goconfig.LoadConfigFile(path + configName)
			if err == nil {
				break
			}
		}
	}

	return cfg, err
}
