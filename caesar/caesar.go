package caesar

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
func askOffset() (int32, error) {
	log.Println("Give me the key, (and offset between -26 and 26):")
	var off int32
	_, err := fmt.Scanf("%d", &off)
	if err != nil {
		return 0, errors.New("Key for caesar should be an integer.")
	}
	if off < -26 || off > 26 {
		return 0, errors.New("Key for caesar should be between -26 and 26.")
	}

	if off < 0 {
		return capMod + off, nil
	}
	return off, nil
}

// Caesar encryption
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

// Caesar decryption
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

func runCaesarMethod(plainTextFile *os.File, cipherFile *os.File, runeAction func(rune, int32) rune) error {
	// Ask for offset
	offset, err := askOffset()
	if err != nil {
		return err
	}

	// Scan character by character
	scanner := bufio.NewScanner(plainTextFile)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {

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
	// Start encrypting with caesar method.
	log.Printf("Encrypting file %s with caesar method.\n", plainTextFile.Name())

	err := runCaesarMethod(plainTextFile, cipherFile, encryptRune)
	if err != nil {
		return err
	}

	log.Printf("Encrypted into %s with caesar method.\n", cipherFile.Name())

	return nil
}

// Decrypt file a to file b, Decrypt(a, b)
func Decrypt(cipherFile *os.File, plainTextFile *os.File) error {
	// Start decrypting with caesar method.
	log.Printf("Decrypting file %s with caesar method.\n", cipherFile.Name())

	err := runCaesarMethod(cipherFile, plainTextFile, decryptRune)
	if err != nil {
		return err
	}

	log.Printf("Decrypted into %s with caesar method.\n", plainTextFile.Name())

	return nil
}
