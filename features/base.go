package features

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

func Base(input []string, wordCount int, prefix string, prefixLength int) {
	// fmt.Fprintln(os.Stderr, "DEBUG: Entering BASE section")

	if len(input) < 3 {
		fmt.Fprintln(os.Stdout, "Error: not enough words to generate text")
		os.Exit(1)
	} else if len(strings.Fields(prefix)) > len(input) {
		fmt.Fprintln(os.Stdout, "Error: suffix not found")
		os.Exit(1)
	}

	combinations := make(map[string][]string)
	for i := 1; i < len(input); i++ {
		if i < len(input)-1 {
			combinations[input[i-1]+" "+input[i]] = append(combinations[input[i-1]+" "+input[i]], input[i+1])
		}
	}

	oneWordKeys := make(map[string][]string) // stores two-word pairs needed for the case when prefix length is 1
	for i := 0; i < len(input)-1; i++ {
		oneWordKeys[input[i]] = append(oneWordKeys[input[i]], input[i+1])
	}

	checkWordCount(wordCount) // ensures that the words displayed are within the limits
	if prefix != "" && prefixLength == -654321 {
		checkTwoWordPresence(combinations, prefix)
		printText(combinations, input, wordCount, prefix)
	} else if prefixLength != -654321 && prefix != "" {
		checkPrefixLength(prefixLength) // checks if the starting prefix length is within the limits
		if len(strings.Fields(prefix)) != prefixLength {
			fmt.Fprintln(os.Stderr, "Error: prefix length and inputted lengths do not match")
			os.Exit(1)
		}
		startingPrefix := checkPrefixPresence(combinations, oneWordKeys, prefix, prefixLength)
		prefixPrs := checkFullPrefixPresence(input, prefix)

		if !prefixPrs {
			fmt.Fprintln(os.Stderr, "Error: the prefix is not fully present in the text")
			os.Exit(1)
		}

		prefixSlc := strings.Fields(prefix)
		for i := 0; i < prefixLength-2; i++ {
			fmt.Print(prefixSlc[i] + " ")
		}

		printText(combinations, input, wordCount, startingPrefix)
	} else if prefix == "" {
		if prefixLength == -654321 {
			// fmt.Fprintln(os.Stderr, "DEBUG: Entering PREFIX == '' and PREFIXLENGTH == -654321 section")
			printText(combinations, input, wordCount, prefix)
		} else if prefixLength != -654321 {
			if prefixLength != 0 {
				fmt.Fprintln(os.Stderr, "Error: inputted length must be 0 when prefix is empty string")
				os.Exit(1)
			} else {
				printText(combinations, input, wordCount, prefix)
			}
		}
	}
}

func printText(combinations map[string][]string, inputText []string, wordCount int, startingPrefix string) {
	// fmt.Fprintln(os.Stderr, "DEBUG: Entering PRINTTEXT section")
	prefixLength := 2
	var firstPartPrefix string
	var secondPartPrefix string

	// when the starting prefix is not set or empty - it takes in
	// as default values the first 2 words in the text
	if startingPrefix == "" {
		firstPartPrefix = inputText[0]
		secondPartPrefix = inputText[1]
	} else {
		firstPartPrefix = strings.Fields(startingPrefix)[0]
		secondPartPrefix = strings.Fields(startingPrefix)[1]
	}

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
		// fmt.Fprintln(os.Stderr, "DEBUG: Entering PRINTTEXT -> LOOP section")
		slcLen := len(combinations[prefix])
		if slcLen == 0 {
			fmt.Fprintln(os.Stdout, "\nSuffix for the last prefix is empty")
			os.Exit(0)
		}

		idxRnd := rand.Intn(slcLen)
		suffix := combinations[prefix][idxRnd]
		fmt.Print(suffix + " ")

		firstPartPrefix = secondPartPrefix
		secondPartPrefix = suffix
		prefix = firstPartPrefix + " " + secondPartPrefix
	}
	fmt.Println()
}

func InputHandler() ([]string, int, string, int) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Fprintln(os.Stderr, "Error: no input into Stdin")
		os.Exit(1)
	}

	wordCountPtr := flag.Int("w", 100, "an int flag that allows to set the maximum number of words to display")
	prefixPtr := flag.String("p", "", "a string flag that allows to set the starting prefix from the given text")
	prefixLengthPtr := flag.Int("l", -654321, "an int flag that allows to set the length of the starting prefix")
	helpPtr := flag.Bool("help", false, "a bool flag that prints usage information")
	flag.Parse()

	if *helpPtr {
		fmt.Fprintln(os.Stdout, "\nMarkov Chain text generator.\n")
		fmt.Fprintln(os.Stdout, "Usage:")
		fmt.Fprintln(os.Stdout, "	markovchain [-w <N>] [-p <S>] [-l <N>]")
		fmt.Fprintln(os.Stdout, "	markovchain --help\n")
		fmt.Fprintln(os.Stdout, "Options:")
		fmt.Fprintln(os.Stdout, "	--help  Show this screen")
		fmt.Fprintln(os.Stdout, "	-w N  	Number of maximum words")
		fmt.Fprintln(os.Stdout, "	-p S  	Starting prefix")
		fmt.Fprintln(os.Stdout, "	-l N  	Prefix length\n")
		os.Exit(0)
	}

	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
		os.Exit(1)
	}

	input := string(bytes)
	inputText := strings.Fields(input)

	return inputText, *wordCountPtr, *prefixPtr, *prefixLengthPtr
}
