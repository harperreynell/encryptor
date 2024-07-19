package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"log"
)

func encrypt(filename string, key []byte) {
	plainText, err := ioutil.ReadFile(filename)
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
	
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Print("Nonce err: %v", err.Error())
	}
	
	cipherText := gcm.Seal(nonce, nonce, plainText, nil)
	
	err = ioutil.WriteFile(filename, cipherText, 0700)
	if err!= nil {
		log.Print("Write file err: %v", err.Error())
	}
}
