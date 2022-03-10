package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func execCipherMethod(method string, encryptKey string, decryptKey string, plaintextName string, shouldFail bool, t *testing.T) {

	methodFlag := fmt.Sprintf("-%s", method)
	plaintextFile := fmt.Sprintf("%s.txt", plaintextName)
	cipherFile := fmt.Sprintf("%s_ciphertext_%s.txt", plaintextName, method)
	decryptedFile := fmt.Sprintf("%s_decryptedtext_%s.txt", plaintextName, method)

	var cmd = struct {
		g string
		r string
		c string
	}{"go", "run", "classic-cypher.go"}

	var args = []struct {
		a string
		m string
		s string
		t string
	}{
		{"-encrypt", methodFlag, plaintextFile, cipherFile},
		{"-decrypt", methodFlag, cipherFile, decryptedFile},
	}

	encryptCmd := exec.Command(cmd.g, cmd.r, cmd.c, args[0].a, args[0].m, args[0].s, args[0].t)
	encryptCmd.Stdin = strings.NewReader(encryptKey)
	_, err := encryptCmd.Output()
	if err != nil {
		t.Fatal("Couldnt complete encryption command.")
	}

	decryptCmd := exec.Command(cmd.g, cmd.r, cmd.c, args[1].a, args[1].m, args[1].s, args[1].t)
	decryptCmd.Stdin = strings.NewReader(decryptKey)
	_, err = decryptCmd.Output()
	if err != nil {
		t.Fatal("Couldnt complete decryption command.")
	}

	var diff = struct {
		d string
		a string
		b string
	}{"diff", args[0].s, args[1].t}

	diffCmd := exec.Command(diff.d, diff.a, diff.b)
	_, err = diffCmd.Output()
	if err != nil {
		if !shouldFail {
			t.Fatal("Couldnt complete diff command.")
		}
	}
}

func TestCaesarMethod(t *testing.T) {
	execCipherMethod("caesar", "3", "3", "texts/plaintext", false, t)
}

func TestCaesarNegMethod(t *testing.T) {
	execCipherMethod("caesar", "-22", "-22", "texts/plaintext_neg", false, t)
}

func TestCaesarCapsMethod(t *testing.T) {
	execCipherMethod("caesar", "3", "3", "texts/plaintext_caps", false, t)
}

func TestCaesarCombinedMethod(t *testing.T) {
	execCipherMethod("caesar", "3", "3", "texts/plaintext_combined", false, t)
}

func TestCaesarLongMethod(t *testing.T) {
	execCipherMethod("caesar", "3", "3", "texts/plaintext_long", false, t)
}

func TestWrongKeyCaesarMethod(t *testing.T) {
	execCipherMethod("caesar", "3", "15", "texts/plaintext_wrong", true, t)
}

func TestVigenereMethod(t *testing.T) {
	execCipherMethod("vigenere", "point", "point", "texts/plaintext", false, t)
}

func TestVigenereCapsMethod(t *testing.T) {
	execCipherMethod("vigenere", "point", "point", "texts/plaintext_caps", false, t)
}

func TestVigenereCombinedMethod(t *testing.T) {
	execCipherMethod("vigenere", "point", "point", "texts/plaintext_combined", false, t)
}

func TestVigenereLongMethod(t *testing.T) {
	execCipherMethod("vigenere", "point", "point", "texts/plaintext_long", false, t)
}

func TestWrongKeyVigenereMethod(t *testing.T) {
	execCipherMethod("vigenere", "point", "tniop", "texts/plaintext_wrong", true, t)
}
