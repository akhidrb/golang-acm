package main

import "strings"

func toGoatLatin(sentence string) string {
	vowels := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}
	suffix := "ma"
	words := strings.Split(sentence, " ")
	for i, word := range words {
		first := word[0]
		suffix += "a"
		if _, ok := vowels[rune(first)]; ok {
			words[i] = word + suffix
		} else {
			words[i] = word[1:] + string(first) + suffix
		}
	}
	return strings.Join(words, " ")
}
