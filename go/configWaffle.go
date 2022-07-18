package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	EnvWafflePath = "WAFFLE_ADDR"
)

type ConfigWaffle struct {
	Path string
}

func DefaultConfig() *ConfigWaffle {
	config := &ConfigWaffle{
		Path: ".",
	}
	return config
}

func configWaffle() *ConfigWaffle {
	config := DefaultConfig()
	if err := config.ReadEnvironmentVariables(); err != nil {
		fmt.Println("INFO Could not read all environment variables. ", err)
	}
	fmt.Println("Environment variables: ", EnvWafflePath, config.Path)
	return config
}

func (config *ConfigWaffle) ReadEnvironmentVariables() error {
	envWafflePath := os.Getenv(EnvWafflePath)
	if envWafflePath != "" {
		config.Path = envWafflePath
	} else {
		return errors.New(EnvWafflePath + " is empty - defaults to current directory.")
	}
	return nil
}