package main

import (
	"strings"
	"unicode"
)

func properTitle(input string) string {
	words := strings.Split(input, " ")
	smallwords := " a an on the to "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") && word != string(word[0]) {
			words[index] = word
		} else {
			r := []rune(word)
			words[index] = string(append([]rune{unicode.ToUpper(r[0])}, r[1:]...))
		}
	}
	return strings.Join(words, " ")
}
