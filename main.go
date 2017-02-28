package main

import (
	"log"
	"flag"
	"os/exec"
)

var (
	configPath = flag.String("config", "config.yaml", "config path")
)

func init() {
	flag.Parse()
}

func callback(config *Config) error {
	for _, command := range config.Commands {
		cmd := exec.Command(command)
		err := cmd.Run()

		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	config, err := NewConfig(*configPath)

	if err != nil {
		panic(err)
	}

	if config.Log != "" {
		err := InitConfig(config.Log)

		if err != nil {
			panic(err)
		}
	}

	_, err := NewNotifier(config, callback)

	if err != nil {
		panic(err)
	}
}