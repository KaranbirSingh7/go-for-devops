package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	configFileName = "config.txt"
)

func main() {
	fmt.Println("Starting")

	configFilePath := filepath.Join(".", configFileName)
	// read file
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// split key vaules
	values := strings.Split(string(data), "=")
	for _, val := range values {
		fmt.Println(val)
	}
}
