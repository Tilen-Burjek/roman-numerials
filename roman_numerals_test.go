package romannumerals
import (
	"errors"
	"testing"
)
func TestRomanNumerals(t *testing.T) {
	tc := append(romanNumeralTests, []romanNumeralTest{
		{0, "", true},
		{-1, "", true},
		{3001, "", true},
	}...)

	for _, test := range tc {
		actual, err := ToRomanNumeral(test.arabic)
		if err == nil && test.hasError {
			t.Errorf("ToRomanNumeral(%d) should return an error.", test.arabic)
			continue
		}
		if err != nil && !test.hasError {
			var _ error = err
			t.Errorf("ToRomanNumeral(%d) should not return an error.", test.arabic)
			continue
		}
		if !test.hasError && actual != test.roman {
			t.Errorf("ToRomanNumeral(%d): %s, expected %s", test.arabic, actual, test.roman)
		}
	}
}

func BenchmarkRomanNumerals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range romanNumeralTests {
			ToRomanNumeral(test.arabic)
		}
	}
}

type A2R struct {
	arab int
	roman string
}

var lookup = []A2R {
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},

}

func ToRomanNumeral(inputNum int) (string, error) {
	if inputNum < 1 || inputNum > 3000 {
		return "", errors.New("number is out of bounds")
	}
	var output string

	for _, n := range lookup {
		for inputNum >= n.arab {
			output += n.roman
			inputNum -= n.arab
		}
	}

	return output, nil
}