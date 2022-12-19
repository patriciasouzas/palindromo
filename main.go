package palindromo

func main() {
	isPalindrome("")
}

func isPalindrome(text string) bool {
	if len(text) == 0 {
		return false
	}
	return invertText(text) == text
}

func isPalindrome2(text string) bool {
	if len(text) == 0 {
		return false
	}

	for i, j := 0, len(text)-1; i < j; i, j = i+1, j-1 {
		if text[i] != text[j] {
			return false
		}
	}

	return true
}

func isPalindrome3(text string) bool {
	if len(text) == 0 {
		return false
	}
	return invertText2(text) == text
}

func invertText(text string) string {
	var invertedText string

	for i := len(text) - 1; i >= 0; i-- {
		invertedText += string(text[i])
	}

	return invertedText
}

func invertText2(text string) string {
	var invertedText []byte = make([]byte, len(text))

	for i := len(text) - 1; i >= 0; i-- {
		invertedText = append(invertedText, text[i])
	}

	return string(invertedText)
}
