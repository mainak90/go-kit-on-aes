package helpers

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

// Implement EncryptService interface
type EncryptServiceInstance struct{}

func(EncryptServiceInstance) Encrypt(_ context.Context, key string, text string) (string, error) {

	//Since the key is in string, we need to convert decode it to bytes
	key_s, _ := hex.DecodeString(key)
	plaintext := []byte(text)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key_s)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return err.Error(), err
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func(EncryptServiceInstance) Decrypt(_ context.Context, message string, key string) (string, error) {

	key_s, _ := hex.DecodeString(key)
	enc, _ := hex.DecodeString(message)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key_s)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err.Error(), err
	}

	return string(plaintext), nil
}

func (EncryptServiceInstance) GenKey(_ context.Context, bit int) (string, error) {
	bytes := make([]byte, bit)
	if _, err := rand.Read(bytes); err != nil {
		return err.Error(), err
	}
	return hex.EncodeToString(bytes), nil
}