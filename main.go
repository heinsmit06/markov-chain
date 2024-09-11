package main

import (
	"markov-chain/features"
)

func main() {
	inputText, wordCount, prefix, prefixLength := features.InputHandler()
	features.Base(inputText, wordCount, prefix, prefixLength)
}
