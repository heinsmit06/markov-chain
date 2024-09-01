package main

import (
	"markov-chain/features"
)

func main() {
	inputText := features.InputHandler()
	features.Base(inputText)
}
