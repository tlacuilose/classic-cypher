package caesar

import (
	"fmt"
	"testing"
)

func TestEncryptRune(t *testing.T) {
	var tests = []struct {
		a, b rune
		off  int32
	}{
		{'a', 'b', 1},
		{'A', 'B', 1},
		{'c', 'h', 5},
		{'C', 'H', 5},
		{'z', 'b', 2},
		{'Z', 'B', 2},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("Encrypt %s to %s", string(tt.a), string(tt.b))
		t.Run(testname, func(t *testing.T) {
			r := encryptRune(tt.a, tt.off)
			if r != tt.b {
				t.Errorf("got %s, wanted %s", string(tt.a), string(r))
			}
		})

	}
}

func TestDecryptRune(t *testing.T) {
	var tests = []struct {
		a, b rune
		off  int32
	}{
		{'b', 'a', 1},
		{'B', 'A', 1},
		{'h', 'c', 5},
		{'H', 'C', 5},
		{'b', 'z', 2},
		{'B', 'Z', 2},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("Decrypt %s to %s", string(tt.a), string(tt.b))
		t.Run(testname, func(t *testing.T) {
			r := decryptRune(tt.a, tt.off)
			if r != tt.b {
				t.Errorf("got %s, wanted %s", string(tt.a), string(r))
			}
		})

	}
}
