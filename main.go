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

	key, err := getKey()

	if err != nil {
		log.Fatal(err)
	}

	dataOrigen, err := getData()

	if err != nil {
		log.Fatal(err)
	}

	dataEncriptada, err := encriptar(key, dataOrigen)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Encriptando")

	err = crearArchivo("encriptado.jpg", dataEncriptada)

	if err != nil {
		log.Fatal(err)
	}

	dataDesencriptada, err := desencriptar(key, dataEncriptada)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("---------------------------------")

	fmt.Println("Desencriptando")

	err = crearArchivo("desencriptado.jpg", dataDesencriptada)

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

func getKey() (key []byte, err error) {

	f, err := os.Open("key.txt")

	if err != nil {
		return
	}

	defer f.Close()

	key, err = io.ReadAll(f)

	if err != nil {
		return
	}

	return
}

func getData() (data []byte, err error) {

	data, err = ioutil.ReadFile("image.jpg")

	if err != nil {
		return
	}

	return
}

func crearArchivo(name string, data []byte) (err error) {

	f, err := os.Create(name)

	if err != nil {
		return
	}

	defer f.Close()

	_, err = f.Write(data)

	if err != nil {
		return
	}

	return
}
