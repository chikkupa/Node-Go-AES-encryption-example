package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"strconv"
)

var iv = []byte("1100273601254817")

// Encrypt Encrypts a string for usage with node using AES encryption
func Encrypt(key, text string) string {
	// Convert the text to encrpyt into a byte array and then pad the array to make sure it is in blocks of 16
	ciphertext := []byte(text)
	ciphertext = pad(ciphertext)

	// Check that the bytes in the text array are devisable by the block size. This is required for the encrytion
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("Cipher text is " + strconv.Itoa(len(ciphertext)) + " long but should be " + strconv.Itoa(aes.BlockSize))
	} else {
		// Create a new cipher using the key you want to use. IMPORTANT!!!!! The key needs to be 32 characters long
		block, err := aes.NewCipher([]byte(key))

		if err != nil {
			panic("Error creating cipher")
		} else {
			// Set up the encrypter
			mode := cipher.NewCBCEncrypter(block, iv)

			// Encrypt the string
			mode.CryptBlocks(ciphertext, ciphertext)

			// We need to convert the encrypted string into something that ca be passed. We convert it into a hex string.
			return hex.EncodeToString(ciphertext)
		}
	}
}

// Decrypt Decrypts aa encrypted AES string and is compatible with node crypto AES encryption
func Decrypt(key, text string) []byte {
	// The byte array we will return with the decrypted string. This could be converted to a string if required
	var ciphertext []byte

	// Assuming that this was encypted using either the node or the other method in this file it needs to be decoded from a hex string
	ciphertext, _ = hex.DecodeString(text)

	// Create the cipher for decrypting the encrypted string.IMPORTANT!!!! this key needs to be 32 characters long
	block, _ := aes.NewCipher([]byte(key))

	// Check that the cipher array size is devisable by the block size. If not then it cannot decrypt
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("Cipher text is " + strconv.Itoa(len(ciphertext)) + " long but should be " + strconv.Itoa(aes.BlockSize))
	} else {
		mode := cipher.NewCBCDecrypter(block, []byte(iv))

		// Decrypt the text into the bye array we set up
		mode.CryptBlocks(ciphertext, ciphertext)

		// The passed in text was probably padded when encypted to make up the block size. We need to remove this padding now the decryption is done
		ciphertext = unpad(ciphertext)
	}

	// Return the decrypted text (Assuming everything went well)
	return ciphertext
}

// pad Pads a byte array to get it up to the required block size needed for encoding to AES
func pad(in []byte) []byte {
	padding := 16 - (len(in) % 16)

	if padding == 0 {
		padding = 16
	}

	for i := 0; i < padding; i++ {
		in = append(in, byte(padding))
	}

	return in
}

// unpad Remove the characters that are present after decrypting an AES string
func unpad(in []byte) []byte {
	if len(in) == 0 {
		return nil
	}

	padding := in[len(in)-1]

	if int(padding) > len(in) || padding > aes.BlockSize {
		return nil
	} else if padding == 0 {
		return nil
	}

	for i := len(in) - 1; i > len(in)-int(padding)-1; i-- {
		if in[i] != padding {
			return nil
		}
	}

	return in[:len(in)-int(padding)]
}
