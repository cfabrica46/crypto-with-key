package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	fkey, err := os.Open("key.txt")

	if err != nil {
		return
	}

	defer fkey.Close()

	key, err := io.ReadAll(fkey)

	if err != nil {
		return
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

	f1, err := os.Create("encriptado.jpg")

	if err != nil {
		log.Fatal(err)
	}

	defer f1.Close()

	_, err = f1.Write(dataEncriptada)

	if err != nil {
		log.Fatal(err)
	}

	dataDesencriptada, err := desencriptar(key, dataEncriptada)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("---------------------------------")

	fmt.Println("Desencriptando")

	f2, err := os.Create("desencriptado.jpg")

	if err != nil {
		log.Fatal(err)
	}

	defer f2.Close()

	_, err = f2.Write(dataDesencriptada)

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
