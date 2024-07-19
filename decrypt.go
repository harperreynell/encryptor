package main

import (
	"crypto/aes"
	"crypto/cipher"
	"io/ioutil"
	"log"
)

func decrypt(filename string, key []byte) {
	cipherText, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Print("Read file error: %v", err.Error())
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Print("Cipher error: %v", err.Error())
	}
	
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Print("Cipher GCM err: %v", err.Error())
	}
	
	nonce := cipherText[:gcm.NonceSize()]
	cipherText = cipherText[gcm.NonceSize():]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Print("Decrypt file err: %v", err.Error())
	}
	
	err = ioutil.WriteFile(filename, plainText, 0777)
	if err!= nil {
		log.Print("Write file err: %v", err.Error())
	}
}
