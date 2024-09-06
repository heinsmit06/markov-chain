package features

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

func Base(input []string, wordCount int) {
	combinations := make(map[string][]string)
	for i := 1; i < len(input); i++ {
		if i < len(input)-1 {
			combinations[input[i-1]+" "+input[i]] = append(combinations[input[i-1]+" "+input[i]], input[i+1])
		}
	}

	length := len(input) - 1
	combinations[input[length-1]+" "+input[length]] = append(combinations[input[length-1]+" "+input[length]], "")

	/*for i, v := range combinations {
		fmt.Print(i + " ")
		fmt.Println(v)
	}*/

	checkWordCount(wordCount)
	printText(combinations, input, wordCount)
}

func printText(combinations map[string][]string, inputText []string, wordCount int) {
	prefixLength := 2
	firstPartPrefix := inputText[0]
	secondPartPrefix := inputText[1]

	if wordCount == 0 {
		fmt.Fprintln(os.Stdout, "")
		os.Exit(0)
	} else if wordCount == 1 {
		fmt.Fprintln(os.Stdout, firstPartPrefix)
		os.Exit(0)
	}

	fmt.Print(firstPartPrefix + " " + secondPartPrefix + " ")
	prefix := firstPartPrefix + " " + secondPartPrefix

	for i := 0; i < wordCount-prefixLength; i++ {
		slcLen := len(combinations[prefix])
		if slcLen == 0 {
			fmt.Println()
			os.Exit(0)
		}
		// fmt.Println("\nlen:", slcLen)

		idxRnd := rand.Intn(slcLen)
		// fmt.Println("id random:", idxRnd)
		suffix := combinations[prefix][idxRnd]
		fmt.Print(suffix + " ")

		firstPartPrefix = secondPartPrefix
		secondPartPrefix = suffix
		prefix = firstPartPrefix + " " + secondPartPrefix
	}
	fmt.Println()
}

func InputHandler() ([]string, int) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Fprintln(os.Stderr, "Error: no input into Stdin")
		os.Exit(1)
	}

	wordCountPtr := flag.Int("w", 100, "an int flag that allows to set the maximum number of words to display")
	flag.Parse()

	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
		os.Exit(1)
	}

	input := string(bytes)
	inputText := strings.Fields(input)

	return inputText, *wordCountPtr
}
