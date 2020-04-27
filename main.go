package main

import (
	"Node-Go-AES-encryption-example/encryption"
	"fmt"
)

func main() {
	var text = "Hi, How are you?"
	var key = "01234567890123450123456789012345"

	encrypted := encryption.Encrypt(key, text)

	fmt.Println("Encrypted: ", encrypted)

	decrypted := encryption.Decrypt(key, encrypted)

	fmt.Println("Decrypted: ", string(decrypted))
}
