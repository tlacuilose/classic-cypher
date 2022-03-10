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

// Returns next offset and next pointer
func takeNextOffset(word string, pointer int) (int32, int, error) {
	// Save the pointer to prevent infinite loop
	firstPointer := pointer

	// Get and offset
	var offset int32 = -1
	var b byte
	for offset == -1 {
		// Take the current letter pointed in the word
		b = word[pointer]

		if b >= 65 && b <= 90 {
			// Return offset for uppercase
			offset = int32(1 + b - 'A')
		} else if b >= 97 && b <= 122 {
			// Return offset for lowercase
			offset = int32(1 + b - 'a')
		}
		// Dont change offset if character in word is not a letter

		// Update pointer to point into next letter
		pointer += 1
		if pointer >= len(word) {
			// Return to first letter at the end
			pointer = 0
		}

		// Break loop if the pointer looped
		if pointer == firstPointer {
			return 0, 0, errors.New("Could not find a valid offset in word (key), only use azAZ.")
		}
	}

	// Return the next offset and the currect pointer
	return offset, pointer, nil
}

func runVigenereMethod(plainTextFile *os.File, cipherFile *os.File, runeAction func(rune, int32) rune) error {
	// Ask for offset
	word, err := askWord()
	if err != nil {
		return err
	}

	// Offset for that letter
	var offset int32
	/// Pointer for the next offset.
	wordPointer := 0

	// Scan character by character
	scanner := bufio.NewScanner(plainTextFile)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {

		// Get the next offset and update the word pointer.
		offset, wordPointer, err = takeNextOffset(word, wordPointer)
		if err != nil {
			return err
		}

		// Get encrypted character
		r := runeAction(rune(scanner.Bytes()[0]), offset)

		// Write encrypted character to file
		_, err = cipherFile.WriteString(string(r))
		if err != nil {
			return err
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
