package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
)

func encrypt(filename string, key []byte) {
	plainText, _ := ioutil.ReadFile(filename)

	block, _ := aes.NewCipher(key)

	gcm, _ := cipher.NewGCM(block)
	
	nonce := make([]byte, gcm.NonceSize())
	
	_, _ = io.ReadFull(rand.Reader, nonce)

	cipherText := gcm.Seal(nonce, nonce, plainText, nil)
	
	_ = ioutil.WriteFile(filename, cipherText, 0700)
}
