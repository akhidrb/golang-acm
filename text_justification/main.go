package main

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	resList := make([]string, 0)
	wordsList := make([]string, 0)
	sum := 0
	for _, word := range words {
		if len(word)+sum+len(wordsList) > maxWidth {
			for i := 0; i < maxWidth-sum; i++ {
				wordsList[i%max(1, len(wordsList)-1)] += " "
			}
			resList = append(resList, strings.Join(wordsList, ""))
			wordsList = make([]string, 0)
			sum = 0
		}
		wordsList = append(wordsList, word)
		sum += len(word)
	}

	if len(wordsList) > 0 {
		lastLine := strings.Join(wordsList, " ")
		for len(lastLine) < maxWidth {
			lastLine += " "
		}
		resList = append(resList, lastLine)
	}

	return resList
}
