package obfuscation

import (
	"fmt"
	"math/rand"
	"strings"
	"Ebyte-VBS-Obf/utils"
)

// RandCapitalization randomly capitalizes letters in a string.
func RandCapitalization(characters string) string {
	var capiString strings.Builder
	for _, char := range characters {
		if rand.Intn(2) == 0 {
			capiString.WriteRune(utils.ToUpper(char))
		} else {
			capiString.WriteRune(utils.ToLower(char))
		}
	}
	return capiString.String()
}

// Obfu obfuscates a string by encoding characters.
func Obfu(body string) string {
	var encBody strings.Builder
	for i, char := range body {
		if i == 0 {
			encBody.WriteString(Expr(int(char)))
		} else {
			encBody.WriteString("*" + Expr(int(char)))
		}
	}
	return encBody.String()
}

func Expr(char int) string {
	rangeVal := rand.Intn(9901) + 100
	exp := rand.Intn(3)

	switch exp {
	case 0:
		return fmt.Sprintf("%d-%d", rangeVal+char, rangeVal)
	case 1:
		return fmt.Sprintf("%d+%d", char-rangeVal, rangeVal)
	case 2:
		return fmt.Sprintf("%d/%d", char*rangeVal, rangeVal)
	}
	return ""
}
