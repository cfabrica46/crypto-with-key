package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	key, err := ioutil.ReadFile("key.txt")

	if err != nil {
		log.Fatal(err)
	}

	dataOrigen, err := ioutil.ReadFile("image.jpg")

	if err != nil {
		log.Fatal(err)
	}

	dataEncriptada, err := encriptar(key, dataOrigen)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Encriptando")

	err = ioutil.WriteFile("encriptado.jpg", dataEncriptada, 0644)

	if err != nil {
		log.Fatal(err)
	}

	dataDesencriptada, err := desencriptar(key, dataEncriptada)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("---------------------------------")

	fmt.Println("Desencriptando")

	err = ioutil.WriteFile("desencriptado.jpg", dataDesencriptada, 0644)

	if err != nil {
		log.Fatal(err)
	}
}

func encriptar(key, dataOrigen []byte) (dataEncriptada []byte, err error) {

	block, err := aes.NewCipher(key)

	if err != nil {
		return
	}

	aesGCM, err := cipher.NewGCM(block)

	if err != nil {
		return
	}

	nonce := make([]byte, aesGCM.NonceSize())

	dataEncriptada = aesGCM.Seal(nonce, nonce, dataOrigen, nil)

	return
}

func desencriptar(key, dataEncriptada []byte) (dataDesencriptada []byte, err error) {

	block, err := aes.NewCipher(key)

	if err != nil {
		return
	}

	aesGCM, err := cipher.NewGCM(block)

	if err != nil {
		return
	}

	nonceSize := aesGCM.NonceSize()

	nonce, cipherText := dataEncriptada[:nonceSize], dataEncriptada[nonceSize:]

	dataDesencriptada, err = aesGCM.Open(nil, nonce, cipherText, nil)

	if err != nil {
		return
	}

	return
}
