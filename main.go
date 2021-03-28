package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	key, err := getKey()

	if err != nil {
		log.Fatal(err)
	}

	textoAEnviar := []byte("holaaa :v")

	textoEncriptadoString, err := encriptar(key, textoAEnviar)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Encriptando")

	fmt.Printf("%s -> %s\n", textoAEnviar, textoEncriptadoString)

	textoDesencriptado, err := desencriptar(key, textoEncriptadoString)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("---------------------------------")

	fmt.Println("Desencriptando")

	fmt.Printf("%s -> %s\n", textoEncriptadoString, textoDesencriptado)

}

func encriptar(key, textoAEnviar []byte) (textoEncriptadoString string, err error) {

	block, err := aes.NewCipher(key)

	if err != nil {
		return
	}

	aesGCM, err := cipher.NewGCM(block)

	if err != nil {
		return
	}

	nonce := make([]byte, aesGCM.NonceSize())

	textoCifrado := aesGCM.Seal(nonce, nonce, textoAEnviar, nil)

	textoEncriptadoString = fmt.Sprintf("%x", textoCifrado)
	return
}

func desencriptar(key []byte, textoEncriptado string) (textoDesencriptado string, err error) {

	enc, err := hex.DecodeString(textoEncriptado)

	if err != nil {
		return
	}

	block, err := aes.NewCipher(key)

	if err != nil {
		return
	}

	aesGCM, err := cipher.NewGCM(block)

	if err != nil {
		return
	}

	nonceSize := aesGCM.NonceSize()

	nonce, cipherText := enc[:nonceSize], enc[nonceSize:]

	textoDesencriptadoBytes, err := aesGCM.Open(nil, nonce, cipherText, nil)

	if err != nil {
		return
	}

	textoDesencriptado = fmt.Sprintf("%s", textoDesencriptadoBytes)

	return
}

func getKey() (key []byte, err error) {

	f, err := os.Open("key.txt")

	if err != nil {
		return
	}

	defer f.Close()

	dataKey, err := io.ReadAll(f)

	if err != nil {
		return
	}

	key, err = hex.DecodeString(string(dataKey))

	if err != nil {
		return
	}

	return
}
