package main

import (
	"fmt"
	"strings"
	"syscall/js"
)

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
	result = strings.ReplaceAll(result, "é", "e")
	result = strings.ReplaceAll(result, "è", "e")
	result = strings.ReplaceAll(result, "ê", "e")
	result = strings.ReplaceAll(result, "ë", "e")
	result = strings.ReplaceAll(result, "É", "E")
	result = strings.ReplaceAll(result, "È", "E")
	result = strings.ReplaceAll(result, "Ê", "E")
	result = strings.ReplaceAll(result, "Ë", "E")

	result = strings.ReplaceAll(result, "á", "a")
	result = strings.ReplaceAll(result, "à", "a")
	result = strings.ReplaceAll(result, "â", "a")
	result = strings.ReplaceAll(result, "ä", "a")
	result = strings.ReplaceAll(result, "Á", "A")
	result = strings.ReplaceAll(result, "À", "A")
	result = strings.ReplaceAll(result, "Â", "A")
	result = strings.ReplaceAll(result, "Ä", "A")

	result = strings.ReplaceAll(result, "í", "i")
	result = strings.ReplaceAll(result, "ì", "i")
	result = strings.ReplaceAll(result, "î", "i")
	result = strings.ReplaceAll(result, "ï", "i")
	result = strings.ReplaceAll(result, "Í", "I")
	result = strings.ReplaceAll(result, "Ì", "I")
	result = strings.ReplaceAll(result, "Î", "I")
	result = strings.ReplaceAll(result, "Ï", "I")

	result = strings.ReplaceAll(result, "ó", "o")
	result = strings.ReplaceAll(result, "ò", "o")
	result = strings.ReplaceAll(result, "ô", "o")
	result = strings.ReplaceAll(result, "ö", "o")
	result = strings.ReplaceAll(result, "Ó", "O")
	result = strings.ReplaceAll(result, "Ò", "O")
	result = strings.ReplaceAll(result, "Ô", "O")
	result = strings.ReplaceAll(result, "Ö", "O")

	result = strings.ReplaceAll(result, "ú", "u")
	result = strings.ReplaceAll(result, "ù", "u")
	result = strings.ReplaceAll(result, "û", "u")
	result = strings.ReplaceAll(result, "ü", "u")
	result = strings.ReplaceAll(result, "Ú", "U")
	result = strings.ReplaceAll(result, "Ù", "U")
	result = strings.ReplaceAll(result, "Û", "U")
	result = strings.ReplaceAll(result, "Ü", "U")

	result = strings.ReplaceAll(result, "ç", "c")
	result = strings.ReplaceAll(result, "Ç", "C")

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
