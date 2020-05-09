package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int{

	words := strings.Fields(s)

	wordCountmap = make(map[string]int)

	for _, word := range words{
		wordCountmap[word]++
	}
	return wordCountmap
}

func main(){
	wc.Test(WordCount)
}