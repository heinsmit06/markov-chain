package features

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func checkWordCount(wordCount int) {
	if wordCount < 0 || wordCount > 10000 {
		fmt.Fprintln(os.Stderr, "Error: number of words must be in between [0, 10000]")
		os.Exit(1)
	}
}

func checkPrefixPresence(combinations, oneWordKeys map[string][]string, prefix string, prefixLength int) string {
	if prefixLength == 1 {
		_, prs := oneWordKeys[prefix] // checks if there is such a word in the input text
		if !prs {
			fmt.Fprintln(os.Stderr, "Error: there is no such prefix in the text")
			os.Exit(1)
		}

		// if prefix length is 1, it refers to another map based on the input text
		// and creates new random prefix which is already present in the main Map
		slcLen := len(oneWordKeys[prefix])
		if slcLen == 0 {
			fmt.Println("")
			os.Exit(1)
		}
		idxRnd := rand.Intn(slcLen)
		startingPrefix := prefix + " " + oneWordKeys[prefix][idxRnd]
		return startingPrefix
	} else {
		slc := strings.Fields(prefix)
		startingPrefix := slc[prefixLength-2] + " " + slc[prefixLength-1]
		checkTwoWordPresence(combinations, startingPrefix)
		return startingPrefix
	}
}

func checkPrefixLength(length int) {
	if length < 0 || length > 5 {
		fmt.Fprintln(os.Stderr, "Error: the starting prefix length must be in between [0, 5]")
		os.Exit(1)
	}
}

// check whether the prefix is present in the main Map
func checkTwoWordPresence(combinations map[string][]string, prefix string) {
	if len(strings.Fields(prefix)) == 2 {
		_, prs := combinations[prefix]
		if !prs {
			fmt.Fprintln(os.Stderr, "Error: there is no such suffix in the text")
			os.Exit(1)
		}
	} else {
		fmt.Fprintln(os.Stderr, "Error: incorrect number of words in a prefix, minimum of 2 required")
		os.Exit(1)
	}
}

func checkFullPrefixPresence(inputText []string, prefix string) bool {
	prefixWords := strings.Fields(prefix)
	prefixLen := len(prefixWords)

	// Loop through inputText and check for a match with the prefix
	for i := 0; i <= len(inputText)-prefixLen; i++ {
		// Check if the slice from inputText matches the prefix words
		if matchPrefix(inputText[i:i+prefixLen], prefixWords) {
			return true
		}
	}

	// If no match is found
	return false
}

// Helper function to compare two string slices element by element
func matchPrefix(inputSlice, prefixSlice []string) bool {
	for i := range prefixSlice {
		if inputSlice[i] != prefixSlice[i] {
			return false
		}
	}
	return true
}
