package main

import (
	"fmt"
	"strings"
	"syscall/js"
)

func exportRemoveDiacritics() {

	js.Global().Set("frenchD", map[string]interface{}{"removeDiacritics": RemoveDiacritics})
}

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
}