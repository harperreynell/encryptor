package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
)

func encrypt(filename string, key []byte) error {
	plainText, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}
	nonce := make([]byte, gcm.NonceSize())
	
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return err
	}
	
	cipherText := gcm.Seal(nonce, nonce, plainText, nil)
	
	err = ioutil.WriteFile(filename, cipherText, 0700)
	if err != nil {
		return err
	}
	
	return nil
}
