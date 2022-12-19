package palindromo

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	type Case struct {
		input          string
		expectedResult bool
	}

	testCases := []Case{
		{
			input:          "ana",
			expectedResult: true,
		},
		{
			input:          "mamam",
			expectedResult: true,
		},
		{
			input:          "sergio",
			expectedResult: false,
		},
		{
			input:          "",
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		if isPalindrome(testCase.input) != testCase.expectedResult {
			t.Error("FAIL", testCase)
		}
	}

}

func TestIsPalindrome2(t *testing.T) {
	type Case struct {
		input          string
		expectedResult bool
	}

	testCases := []Case{
		{
			input:          "ana",
			expectedResult: true,
		},
		{
			input:          "mamam",
			expectedResult: true,
		},
		{
			input:          "sergio",
			expectedResult: false,
		},
		{
			input:          "",
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		if isPalindrome2(testCase.input) != testCase.expectedResult {
			t.Error("FAIL", testCase)
		}
	}

}

func TestInvertText(t *testing.T) {
	type Case struct {
		input          string
		expectedResult string
	}

	testCases := []Case{
		{
			input:          "ana",
			expectedResult: "ana",
		},
		{
			input:          "mamam",
			expectedResult: "mamam",
		},
		{
			input:          "sergio",
			expectedResult: "oigres",
		},
	}

	for _, testCase := range testCases {
		if invertText(testCase.input) != testCase.expectedResult {
			t.Error("FAIL", testCase)
		}
	}

}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome("MAMAM")
	}
}

func BenchmarkIsPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome2("MAMAM")
	}
}

func BenchmarkIsPalindrome3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome3("MAMAM")
	}
}
