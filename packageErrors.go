package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
)

type Config struct {
	// Configuration fields go here (redacted)
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "can't open configuration file")
	}

	defer file.Close()

	cfg := &Config{}
	// Parse file here (redacted)
	return cfg, nil
}

func setUpLogging() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(out)
}

func main() {
	setUpLogging()
	cfg, err := readConfig("/path/to/config.toml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		log.Printf("error: %+v", err)
		os.Exit(1)
	}

	// Normal operation (redacted)
	fmt.Println(cfg)
}
