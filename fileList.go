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

			if info.IsDir() && path == "/dev" {
				return filepath.SkipDir
			}
			if info.IsDir() && path == "/proc" {
				return filepath.SkipDir
			}
			if info.IsDir() && path == "/run" {
				return filepath.SkipDir
			}
			if info.IsDir() && path == "/sys" {
				return filepath.SkipDir
			}

			if info != nil {
				List = append(List, path)
			}
			return nil
		})

	return List
}
