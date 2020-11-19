package main

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/mainak90/go-kit-aes/utils"
	"log"
)

func main() {
	// Create a 32 byte random key for AES
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}
	key := hex.EncodeToString(bytes)
	message := "This is the string to encrypt"
	log.Println("Original message: ", message)
	encryptedstring := utils.EncryptString(message, key)
	log.Println("Encrypted string: ", encryptedstring)
	decryptedstring := utils.DecryptString(encryptedstring, key)
	log.Println("Decryptedstring: ", decryptedstring)
}
