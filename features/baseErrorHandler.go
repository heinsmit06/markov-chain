package features

import (
	"fmt"
	"os"
)

func checkWordCount(wordCount int) {
	if wordCount < 0 || wordCount > 10000 {
		fmt.Fprintln(os.Stderr, "Number of words must be in between [0, 10000]")
		os.Exit(0)
	}
}
