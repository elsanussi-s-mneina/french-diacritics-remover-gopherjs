package main

import (
	"fmt"
	"strings"
	"syscall/js"
)

var replacements = [][2]string{
	{"é", "e"},
	{"è", "e"},
	{"ê", "e"},
	{"ë", "e"},
	{"É", "E"},
	{"È", "E"},
	{"Ê", "E"},
	{"Ë", "E"},
	{"á", "a"},
	{"à", "a"},
	{"â", "a"},
	{"ä", "a"},
	{"Á", "A"},
	{"À", "A"},
	{"Â", "A"},
	{"Ä", "A"},
	{"í", "i"},
	{"ì", "i"},
	{"î", "i"},
	{"ï", "i"},
	{"Í", "I"},
	{"Ì", "I"},
	{"Î", "I"},
	{"Ï", "I"},
	{"ó", "o"},
	{"ò", "o"},
	{"ô", "o"},
	{"ö", "o"},
	{"Ó", "O"},
	{"Ò", "O"},
	{"Ô", "O"},
	{"Ö", "O"},
	{"ú", "u"},
	{"ù", "u"},
	{"û", "u"},
	{"ü", "u"},
	{"Ú", "U"},
	{"Ù", "U"},
	{"Û", "U"},
	{"Ü", "U"},
	{"ç", "c"},
	{"Ç", "C"},
}

// exportRemoveDiacritics allows the RemoveDiacritics function to be called
// by Javascript code.
// This function must be called once in the Go source code (in the main function), in order for the
// RemoveDiacritics function to be accessible from Javascript.
//
// In the Javascript code, call the RemoveDiacritics function thus:
// frenchD.removeDiacritics("Où")
func exportRemoveDiacritics() {
	js.Global().Set("frenchD", map[string]interface{}{"removeDiacritics": RemoveDiacritics})
}

// RemoveDiacritics takes text in the French language
// and returns similar text but without the diacritics
// (the diacritics are circumflex, acute, grave accents, and diareses; and cedilla)
// For example: "éÉàç" becomes "eEac"
func RemoveDiacritics(frenchText string) string {
	var result string

	result = frenchText
	for _, replacement := range replacements {
		result = strings.ReplaceAll(result, replacement[0], replacement[1])
	}
	return result
}

func main() {
	fmt.Println("Removing French Diacritics program")
	exportRemoveDiacritics()
	testRemoveDiacritics()
}

func testRemoveDiacritics() {
	testRemoveDiacriticsCase("áÁéÉÍíóÓÚú", "aAeEIioOUu")
	testRemoveDiacriticsCase("ça où", "ca ou")

}

func testRemoveDiacriticsCase(input, expected string) {
	input = "áÁéÉÍíóÓÚú"
	expected = "aAeEIioOUu"
	var actual string
	actual = RemoveDiacritics(input)
	if actual != expected {
		fmt.Printf("Test failed, expected: %s, actual %s",
			expected, actual)
		panic("Something is not working")
	}
}

