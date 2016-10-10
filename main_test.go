package main

import "testing"

var testCasesNumerics = map[int]string{
	1:    "I",
	2:    "II",
	4:    "IV",
	5:    "V",
	7:    "VII",
	9:    "IX",
	10:   "X",
	12:   "XII",
	14:   "XIV",
	18:   "XVIII",
	118:  "CXVIII",
	127:  "CXXVII",
	129:  "CXXIX",
	159:  "CLIX",
	1004: "MIV",
	1009: "MIX",
	1049: "MXLIX",
	5001: "MMMMMI",
}

func TestDitit2Roman(t *testing.T) {
	for num, rom := range testCasesNumerics {
		if result := Ditit2Roman(num); result != rom {
			t.Errorf("Expect '%s', but get '%s'", rom, result)
		}
	}
}

func TestRoman2Digit(t *testing.T) {
	for num, rom := range testCasesNumerics {
		if result := Roman2Digit(rom); result != num {
			t.Errorf("Expect '%d', but get '%d'", num, result)
		}
	}
}
