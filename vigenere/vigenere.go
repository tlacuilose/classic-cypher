package vigenere

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

const (
	lowMod rune = 1 + 'z' - 'a'
	capMod rune = 1 + 'Z' - 'A'
)

// Ask for the offset (key)
func askWord() (string, error) {
	log.Println("Give me a word (the key):")
	var word string
	_, err := fmt.Scanf("%s", &word)
	if err != nil {
		return "", errors.New("Key for vigenere should be a word")
	}
	/*
		TODO: Check that its one word.
		if off < -26 || off > 26 {
			return 0, errors.New("Key for vigenere should be between -26 and 26.")
		}
	*/
	return word, nil
}

// Vigenere encryption
func encryptRune(b rune, offset int32) rune {
	switch {
	// Encrypt uppercase abc
	case b >= 65 && b <= 90:
		return rune((b-'A'+offset)%capMod) + 'A'
	// Encrypt lowercase abc
	case b >= 97 && b <= 122:
		return rune((b-'a'+offset)%lowMod) + 'a'
	// Dont encrypt other characters
	default:
		return b
	}
}

// Vigenere decryption
func decryptRune(b rune, offset int32) rune {
	switch {
	// Decrypt uppercase abc
	case b >= 65 && b <= 90:
		return rune((b-'A'+capMod-offset)%capMod) + 'A'
	// Decrypt lowercase abc
	case b >= 97 && b <= 122:
		return rune((b-'a'+lowMod-offset)%lowMod) + 'a'
	// Dont decrypt other characters
	default:
		return b
	}
}

func runVigenereMethod(plainTextFile *os.File, cipherFile *os.File, runeAction func(rune, int32) rune) error {
	// Ask for offset
	word, err := askWord()
	if err != nil {
		return err
	}

	wordPointer := 0

	// Scan character by character
	scanner := bufio.NewScanner(plainTextFile)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {

		// Get the offset from the position of the word pointer
		offset := int32(word[wordPointer])

		// Get encrypted character
		r := runeAction(rune(scanner.Bytes()[0]), offset)

		// Write encrypted character to file
		_, err = cipherFile.WriteString(string(r))
		if err != nil {
			return err
		}

		wordPointer += 1
		if wordPointer >= len(word) {
			wordPointer = 0
		}
	}

	// Write file to stable storage
	cipherFile.Sync()

	return nil
}

// Encrypt file a to file b, Encrypt(a, b)
func Encrypt(plainTextFile *os.File, cipherFile *os.File) error {
	// Start encrypting with vigenere method.
	log.Printf("Encrypting file %s with vigenere method.\n", plainTextFile.Name())

	err := runVigenereMethod(plainTextFile, cipherFile, encryptRune)
	if err != nil {
		return err
	}

	log.Printf("Encrypted into %s with vigenere method.\n", cipherFile.Name())

	return nil
}

// Decrypt file a to file b, Decrypt(a, b)
func Decrypt(cipherFile *os.File, plainTextFile *os.File) error {
	// Start decrypting with vigenere method.
	log.Printf("Decrypting file %s with vigenere method.\n", cipherFile.Name())

	err := runVigenereMethod(cipherFile, plainTextFile, decryptRune)
	if err != nil {
		return err
	}

	log.Printf("Decrypted into %s with vigenere method.\n", plainTextFile.Name())

	return nil
}
