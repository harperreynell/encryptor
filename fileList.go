package main

import (
	"path/filepath"
	"os"
	"log"
)

func fileList() []string {
	List := []string {}
	err := filepath.Walk("/",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println(err)
				return nil
			}

			if info != nil {
				List = append(List, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	
	return List
}

