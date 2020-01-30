package main

import (
	"fmt"
	"strings"
)

//count frequency of words in a string of text
//return a map of words with their counts
//convert text to lowercase
//punctuation trimmed from words
//string package contains Fields, ToLower, and Trim.
//display the count for any word that occurs more that once

func countWords(text string) map[string]int {
	//format text
	//convert text to lowercase
	text = strings.ToLower(text)
	//remove punctuation
	text = strings.ReplaceAll(text, ",", " ")
	text = strings.ReplaceAll(text, ";", " ")
	text = strings.ReplaceAll(text, ".", " ")
	text = strings.ReplaceAll(text, "—", " ")
	//split text up into a slice of strings
	words := strings.Fields(text)
	//count frequency of words and display count for any word with count > 1
	wordCount := make(map[string]int, len(words))
	for _, word := range words {
		wordCount[word]++
	}
	//return map of words with their counts
	return wordCount
}

func main() {
	text := `As far as eye could reach he saw nothing but the stems 
	of the great plants about him receding in the violet shade, 
	and far overhead the multiple transparency of huge leaves 
	filtering the sunshine to the solemn splendour of twilight 
	in which he walked. Whenever he felt able he ran again; 
	the ground continued soft and springy, covered with the 
	same resilient weed which was the first thing his hands 
	had touched in Malacandra. Once or twice a small red creature 
	scuttled across his path, but otherwise there seemed 
	to be no life stirring in the wood;	nothing to fear—except the 
	fact of wandering unprovisioned and alone in a forest of unknown 
	vegetation thousands or	millions of miles beyond the reach or knowledge of man.`

	wordCount := countWords(text)
	for word, count := range wordCount {
		if count > 1 {
			fmt.Printf("'%v' occurs %d times\n", word, count)
		}
	}
}
