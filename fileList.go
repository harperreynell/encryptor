package main

import (
	"path/filepath"
	"os"
)

func fileList(root string) []string {
	List := []string {}
	_ = filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}

			if info != nil {
				List = append(List, path)
			}
			return nil
		})
	
	return List
}