package main

import (
	"flag"
	"fmt"
	"nated_crypto/EncDec"
	"os"
)

func main() {
	// Define command-line flags
	setKeyFlag := flag.String("setkey", "", "Set the encryption and decryption key (16 characters long)")
	encryptFlag := flag.String("encrypt", "", "Encrypt the provided string")
	decryptFlag := flag.String("decrypt", "", "Decrypt the provided string")

	// Parse the command-line flags
	flag.Parse()

	// Validate the 'setkey' flag - it must not be empty and should be 16 characters long
	if *setKeyFlag == "" {
		fmt.Println("Please provide a key with --setkey flag.")
		os.Exit(1) // Exit the program if the condition is not met
	}
	if len(*setKeyFlag) != 16 {
		fmt.Println("Key must be exactly 16 characters long.")
		os.Exit(1) // Exit the program if the condition is not met
	}

	// Check if at least one of the encrypt or decrypt flags is set
	if *decryptFlag == "" && *encryptFlag == "" {
		fmt.Println("Please provide an action with --encrypt or --decrypt flag.")
		os.Exit(1) // Exit the program if neither flag is set
	}

	// Create an instance of NFC struct and set its Key field
	nfc := EncDec.NFC{
		Key: []byte(*setKeyFlag), // Convert the key string to a byte slice and assign it to the NFC struct
	}

	// Encrypt the input string if the 'encrypt' flag is set
	if *encryptFlag != "" {
		nfc.EncryptUserInput(*encryptFlag)
		return // Exit the program after encryption
	}

	// Decrypt the input string if the 'decrypt' flag is set
	if *decryptFlag != "" {
		nfc.DecryptUserInput(*decryptFlag)
		return // Exit the program after decryption
	}

	// If neither encrypt nor decrypt actions were taken, remind the user to provide a flag
	fmt.Println("Please provide an action with --encrypt or --decrypt flag.")
}
