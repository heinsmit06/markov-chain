package main

import (
	"markov-chain/features"
)

func main() {
	inputText, wordCount, prefix := features.InputHandler()
	features.Base(inputText, wordCount, prefix)
}
