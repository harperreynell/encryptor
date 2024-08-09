package main

import (
	"os"
	"path/filepath"
)

func fileList(root string) []string {
	List := []string{}
	_ = filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}

			// Skip the /dev folder
			if info.IsDir() && path == "/dev" {
				return filepath.SkipDir
			}

			if info != nil {
				List = append(List, path)
			}
			return nil
		})

	return List
}
