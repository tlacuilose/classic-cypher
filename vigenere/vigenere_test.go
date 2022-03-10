package vigenere

import (
	"fmt"
	"testing"
)

func TestTakeNextOffsetRune(t *testing.T) {
	var tests = []struct {
		w   string
		p   int
		off int32
		np  int
	}{
		{"point", 0, 16, 1},
		{"point", 1, 15, 2},
		{"point", 2, 9, 3},
		{"point", 3, 14, 4},
		{"point", 4, 20, 0},
		// Non letter characters should be skipped
		{"a point", 0, 1, 1},
		{"a point", 1, 16, 3},
		{"ap&oint", 1, 16, 2},
		{"ap&oint", 2, 15, 4},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("Next offset in %s at %d", tt.w, tt.p)
		t.Run(testname, func(t *testing.T) {
			offset, newPointer, err := takeNextOffset(tt.w, tt.p)
			if err != nil {
				t.Fatal("Next offset function failed")

			}
			if offset != tt.off || newPointer != tt.np {
				t.Errorf("off is %d not %d, new pointer is %d not %d", tt.off, offset, tt.np, newPointer)

			}
		})

	}
}
