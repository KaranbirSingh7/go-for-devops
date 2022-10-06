package main

import (
	"io/fs"
	"log"
	"path/filepath"
)

var (
	RootDir string //string for walking main path
)

func main() {
	RootDir = "/tmp/"
	log.Println("Starting to walk filesystem inside:", RootDir)

	filepath.WalkDir(RootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Println("ERROR:", err.Error())
			return err
		}
		if d.IsDir() {
			log.Printf("%q is a directory", d.Name())
		}
		return nil
	})

}
