package cyphercmd

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func TestParseCorrectFlagsAndArgs(t *testing.T) {
	var tests = []struct {
		action         string
		method         string
		sourceFilename string
		targetFilename string
	}{
		{"-encrypt", "-caesar", "plaintext.txt", "ciphertext.txt"},
		{"-decrypt", "-caesar", "plaintext.txt", "ciphertext.txt"},
		{"-encrypt", "-vigenere", "ciphertext.txt", "decryptedtext.txt"},
		{"-decrypt", "-vigenere", "ciphertext.txt", "decryptedtext.txt"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Parsing cmd %v", tt)
		t.Run(testname, func(t *testing.T) {
			os.Args = []string{tt.action, tt.method, tt.sourceFilename, tt.targetFilename}
			_, err := parseFlagsAndArgs()
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			if err == nil {
				t.Errorf("Expected command to not throw errors")
			}
		})
	}
}

func TestParseIncorrectFlagsAndArgs(t *testing.T) {

	var tests = []struct {
		action         string
		method         string
		sourceFilename string
		targetFilename string
	}{
		{"-encrypt", "-decrypt", "plaintext.txt", "ciphertext.txt"},
		{"-vigenere", "-caesar", "plaintext.txt", "ciphertext.txt"},
		{"-encrypt", "-vigenere", "ciphertext", "decryptedtext.txt"},
		{"-decrypt", "-vigenere", "ciphertext.txt", "decryptedtext"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Parsing cmd %v", tt)
		t.Run(testname, func(t *testing.T) {
			os.Args = []string{tt.action, tt.method, tt.sourceFilename, tt.targetFilename}
			_, err := parseFlagsAndArgs()
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			if err == nil {
				t.Errorf("Expected command to not throw errors")
			}
		})
	}
}
