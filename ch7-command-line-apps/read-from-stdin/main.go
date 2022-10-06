package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	var errRE = regexp.MustCompile(`(?i)error`)
	var s *bufio.Scanner

	switch len(os.Args) {
	// 1 argument == aka no argument passed
	case 1:
		log.Println("No file specified, reading from STDIN")
		s = bufio.NewScanner(os.Stdin)
	case 2:
		filepath := os.Args[len(os.Args)-1]
		log.Printf("Reading file: %s\n", filepath)
		f, err := os.Open(filepath)
		if err != nil {
			log.Fatalf(err.Error())
		}
		// close file on exit
		defer f.Close()
		s = bufio.NewScanner(f)
	default:
		log.Fatalln("Too many arguments provided.")
	}

	// read input line by line
	for s.Scan() {
		line := s.Bytes()
		if errRE.Match(line) {
			fmt.Printf("%s\n", line)
		}
	}

	if err := s.Err(); err != nil {
		log.Fatalln("Error:", err.Error())
	}
}
