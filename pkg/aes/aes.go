package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/karanveersp/go-encryption/pkg/util"
)

// Decrypt decrypts the contents of the given ciphertext
// using the given key. Returns plaintext or error.
func Decrypt(key string, ciphertext []byte) ([]byte, error) {
	keyBytes, err := util.PadZeroes32(key)
	if err != nil {
		return nil, err
	}

	c, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, err
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

// Encrypt encrypts plaintext
// using the given key. Returns an array of
// bytes representing the cipher text and an error.
func Encrypt(key string, plaintext []byte) ([]byte, error) {
	keyBytes, err := util.PadZeroes32(key)
	if err != nil {
		return nil, err
	}

	// generate a new aes cipher using key
	c, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptograhic block ciphers.
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	// creates a new byte arraythe size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// here we encrypt our text using the Seal function.
	// Seal encrypts and authenticates plaintext, authenticates
	// the additional data and appends the result to dst, returning
	// the updated slice. The nonce must be the NonceSize() bytes long
	// and unique for all time, for a given key.
	encryptedData := gcm.Seal(nonce, nonce, plaintext, nil)
	return encryptedData, nil
}
