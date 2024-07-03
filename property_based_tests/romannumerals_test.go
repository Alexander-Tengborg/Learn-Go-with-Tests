package romannumerals

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases []struct {
	Arabic uint16
	Roman  string
} = []struct {
	Arabic uint16
	Roman  string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 798, Roman: "DCCXCVIII"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
}

func TestConvertingToRoman(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got, err := ConvertToRoman(test.Arabic)

			if err != nil {
				t.Fatalf("could not convert %q to %q, number was too large", got, test.Roman)
			}

			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got, err := ConvertToArabic(test.Roman)

			if err != nil {
				t.Fatalf("could not convert %d to %d, number was too large", got, test.Arabic)
			}

			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		roman, err := ConvertToRoman(arabic)
		if err != nil {
			return true
		}

		fromRoman, err := ConvertToArabic(roman)
		if err != nil {
			return true
		}

		t.Log("testing", arabic)
		return arabic == fromRoman
	}

	qConfig := &quick.Config{
		MaxCount: 1000,
	}

	if err := quick.Check(assertion, qConfig); err != nil {
		t.Error("failed checks", err)
	}
}

