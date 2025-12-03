package property

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{11, "XI"},
	{12, "XII"},
	{13, "XIII"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{90, "XC"},
	{100, "C"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		description := fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman)
		t.Run(description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		description := fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic)
		t.Run(description, func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

// for any N, taking the result of ConvertToRoman(N) and passing it to ConvertToArabicm should return N
func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing conversion for", arabic)

		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}

func indexRomanNumeral(numeral string) int {
	return slices.IndexFunc(AllRomanNumerals, func(r RomanNumeral) bool {
		return r.Symbol == numeral
	})
}

func subtractorSortFunc(a, b string) int {
	aIdx := indexRomanNumeral(a)
	bIdx := indexRomanNumeral(b)

	order := cmp.Compare(aIdx, bIdx)

	if order < 0 && (b == "I" || b == "X" || b == "C") {
		return 1
	}

	return order
}

// only I, X, and C can be used as 'subtractors'
func TestPropertiesOfSubtractors(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}

		roman := ConvertToRoman(arabic)
		runes := strings.Split(roman, "")
		t.Log("testing subtractors for", arabic, roman)

		return slices.IsSortedFunc(runes, subtractorSortFunc)
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
