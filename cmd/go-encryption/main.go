package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/karanveersp/go-encryption/pkg/aes"
)

func getArgs() (string, string, bool) {
	// declare
	var key, txtFilePath string
	var isEncryptMode bool

	// flag declaration
	flag.StringVar(&key, "key", "", "Secret key to encrypt and decrypt data")
	flag.StringVar(&txtFilePath, "path", "", "Path to file containing plaintext or ciphertext")
	flag.BoolVar(&isEncryptMode, "e", false, "Encrypt mode. Decrypts if not provided.")

	flag.Parse()

	if key == "" || txtFilePath == "" {
		fmt.Println("key and file path are required.")
		flag.Usage()
		os.Exit(1)
	}

	return key, txtFilePath, isEncryptMode
}

func main() {
	fmt.Println("AES Encryption/Decryption Program v0.01")

	key, txtFilePath, encryptMode := getArgs()

	fileData, err := ioutil.ReadFile(txtFilePath)
	if err != nil {
		log.Fatalln(err)
	}

	if encryptMode {
		encryptedData, err := aes.Encrypt(key, fileData)
		if err != nil {
			log.Fatalln(err)
		}
		encFileDir, encFile := filepath.Split(txtFilePath)
		encFileName := strings.Split(encFile, ".")[0] + "_encrypted.txt"
		err = ioutil.WriteFile(path.Join(encFileDir, encFileName), encryptedData, 0777)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Encryption successful")
		return
	}

	decryptedData, err := aes.Decrypt(key, fileData)
	if err != nil {
		log.Fatalln(err)
	}
	fileDir, fileName := filepath.Split(txtFilePath)
	fileName = strings.Replace(fileName, "encrypted", "decrypted", 1)

	err = ioutil.WriteFile(path.Join(fileDir, fileName), []byte(decryptedData), 0777)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Decryption successful")
}
