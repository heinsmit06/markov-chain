package main

import (
	"markov-chain/features"
)

func main() {
	inputText, wordCount := features.InputHandler()
	features.Base(inputText, wordCount)
}
