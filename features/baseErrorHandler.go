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

func checkPrefixPresence(combinations map[string][]string, prefix string) {
	if len(prefix) == 2 {
		_, prs := combinations[prefix]
		if !prs {
			fmt.Fprintln(os.Stderr, "Error: there is no such prefix in the text")
			os.Exit(1)
		}
	} else {
	}
}
