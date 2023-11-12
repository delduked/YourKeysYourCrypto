package EncDec

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"os"
)

// NFC struct holds the encryption key
type NFC struct {
	Key []byte
}

// EncryptUserInput takes a string input and encrypts it
func (tag *NFC) EncryptUserInput(input string) {
	// Convert the input string to bytes and encrypt it
	encryptedData, err := tag.encryptECB([]byte(input))
	if err != nil {
		fmt.Println("Encryption failed:", err)
		os.Exit(1) // Exit if encryption fails
	}

	// Encode the encrypted data into base64 for easy text representation
	encodedData := base64.StdEncoding.EncodeToString(encryptedData)

	// Print the base64 encoded encrypted data
	fmt.Println("Encrypted (Base64):", encodedData)
}

// encryptECB handles the low-level details of the encryption using AES in ECB mode
func (tag *NFC) encryptECB(data []byte) ([]byte, error) {
	// Create a new AES cipher with the key
	block, err := aes.NewCipher(tag.Key)
	if err != nil {
		return nil, err // Return error if cipher creation fails
	}

	// Apply PKCS7 padding to the data
	blockSize := block.BlockSize()
	data = tag.pad(data, blockSize)

	// Encrypt the data block by block
	encrypted := make([]byte, len(data))
	for i := 0; i < len(data); i += blockSize {
		block.Encrypt(encrypted[i:i+blockSize], data[i:i+blockSize])
	}

	return encrypted, nil
}

// DecryptUserInput takes a base64 encoded encrypted string and decrypts it
func (tag *NFC) DecryptUserInput(input string) {
	// Decode the base64 input to get the encrypted data
	encryptedData, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		fmt.Println("Base64 decoding failed:", err)
		os.Exit(1) // Exit if base64 decoding fails
	}

	// Decrypt the data
	decryptedData, err := tag.decryptECB(encryptedData)
	if err != nil {
		fmt.Println("Decryption failed:", err)
		os.Exit(1) // Exit if decryption fails
	}

	// Print the decrypted text
	fmt.Println("Decrypted text:", string(decryptedData))
}

// decryptECB handles the low-level details of the decryption using AES in ECB mode
func (tag *NFC) decryptECB(data []byte) ([]byte, error) {
	// Create a new AES cipher with the key
	block, err := aes.NewCipher(tag.Key)
	if err != nil {
		return nil, err // Return error if cipher creation fails
	}

	// Decrypt the data block by block
	decrypted := make([]byte, len(data))
	blockSize := block.BlockSize()
	for i := 0; i < len(data); i += blockSize {
		block.Decrypt(decrypted[i:i+blockSize], data[i:i+blockSize])
	}

	// Remove the PKCS7 padding
	return tag.unpad(decrypted), nil
}

// pad applies PKCS7 padding to the data
func (tag *NFC) pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// unpad removes PKCS7 padding from the data
func (tag *NFC) unpad(data []byte) []byte {
	length := len(data)
	if length == 0 {
		return data // Return as is if data is empty
	}

	// Calculate and verify the padding
	padding := int(data[length-1])
	if padding > length || padding == 0 {
		// Handle error: invalid padding
		return nil // Return nil or handle the error appropriately
	}

	// Check the padding bytes
	for i := 0; i < padding; i++ {
		if int(data[length-1-i]) != padding {
			// Handle error: invalid padding
			return nil // Return nil or handle the error appropriately
		}
	}

	return data[:length-padding] // Return the data without padding
}
