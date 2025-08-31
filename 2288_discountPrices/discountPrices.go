package discountPrices

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func discountPrices(sentence string, discount int) string {
	words := strings.Split(sentence, " ")
	for i, word := range words {
		if strings.HasPrefix(word, "$") && isNumeric(word[1:]) {
			price, _ := strconv.Atoi(word[1:])
			discountedPrice := float64(price) * (1 - float64(discount)/100)
			words[i] = fmt.Sprintf("$%.2f", discountedPrice)
		}
	}
	return strings.Join(words, " ")
}

func isNumeric(s string) bool {
	match, _ := regexp.MatchString("^[0-9]+$", s)
	return match
}
