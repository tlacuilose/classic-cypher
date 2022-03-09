package main

import (
	"log"
	"os/exec"
	"strings"
	"testing"
)

func TestCaesarMethod(t *testing.T) {
	offset := "3"

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
		{"-encrypt", "-caesar", "texts/plaintext.txt", "texts/ciphertext.txt"},
		{"-decrypt", "-caesar", "texts/ciphertext.txt", "texts/decryptedtext.txt"},
	}

	encryptCmd := exec.Command(cmd.g, cmd.r, cmd.c, args[0].a, args[0].m, args[0].s, args[0].t)
	encryptCmd.Stdin = strings.NewReader(offset)
	_, err := encryptCmd.Output()
	if err != nil {
		t.Fatal("Couldnt complete encryption command.")
	}

	decryptCmd := exec.Command(cmd.g, cmd.r, cmd.c, args[1].a, args[1].m, args[1].s, args[1].t)
	decryptCmd.Stdin = strings.NewReader(offset)
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
	diffOut, err := diffCmd.Output()
	if err != nil {
		log.Println(err)
		t.Fatal("Couldnt complete diff command.")
	}
	if string(diffOut) != "" {
		t.Fatal("Plaintext is not equal to decryptedtext.")
	}

}

func TestWrongKeyCaesarMethod(t *testing.T) {
	offset := "3"
	offsetWrong := "5"

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
		{"-encrypt", "-caesar", "texts/plaintext.txt", "texts/ciphertext.txt"},
		{"-decrypt", "-caesar", "texts/ciphertext.txt", "texts/decryptedtext_wrong.txt"},
	}

	encryptCmd := exec.Command(cmd.g, cmd.r, cmd.c, args[0].a, args[0].m, args[0].s, args[0].t)
	encryptCmd.Stdin = strings.NewReader(offset)
	_, err := encryptCmd.Output()
	if err != nil {
		t.Fatal("Couldnt complete encryption command.")
	}

	decryptCmd := exec.Command(cmd.g, cmd.r, cmd.c, args[1].a, args[1].m, args[1].s, args[1].t)
	decryptCmd.Stdin = strings.NewReader(offsetWrong)
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
	if err == nil {
		t.Fatal("Diff should fail")
	}
}
