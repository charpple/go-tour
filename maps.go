package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {


    wordMap := make(map[string]int)

    // split sentences into words
    words := strings.Fields(s)

  
    for _, word := range words {
        wordMap[word]++
    }

    return wordMap
}

func main() {
    wc.Test(WordCount)
}
