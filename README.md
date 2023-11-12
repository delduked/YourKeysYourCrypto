# Cryptographic CLI Tool

## Overview

This command-line tool, written in Go, is designed for encrypting and decrypting strings using AES encryption. It's particularly useful for securing sensitive data like Bitcoin private keys. The tool uses a user-specified key for encryption and decryption processes.

## Prerequisites

- Go (Golang) installed on your system.
- Access to a terminal or command-line interface.

## Installation

Clone the repository containing the Go code:

```bash
git clone https://github.com/delduked/YourKeysYourCrypto.git
cd YourKeysYourCrypto
```

## Usage

The tool supports three main operations:

1. **Set Encryption/Decryption Key**: Use the `--setkey` flag to specify a 16-character long key for encryption and decryption.

2. **Encrypt a String**: Use the `--encrypt` flag to encrypt a string.

3. **Decrypt a String**: Use the `--decrypt` flag to decrypt a string.

### Examples

- **Setting a Key**: 

  ```bash
  go run main.go --setkey "your-16-char-key"
  ```

- **Encrypting a String**: 

  ```bash
  go run main.go --setkey "your-16-char-key" --encrypt "string to encrypt"
  ```

- **Decrypting a String**: 

  ```bash
  go run main.go --setkey "your-16-char-key" --decrypt "encrypted string"
  ```

Note: Replace `your-16-char-key` with your actual encryption key, and replace `string to encrypt` or `encrypted string` with the actual string you want to encrypt or decrypt.