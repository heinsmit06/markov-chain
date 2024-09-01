package features

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Base(input []string) {
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
}

func InputHandler() []string {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
		os.Exit(1)
	}

	input := string(bytes)
	inputText := strings.Fields(input)

	return inputText
}

/*
_, present := combinations[input[i-1]+input[i]]
			if present {
				combinations[input[i-1]+" "+input[i]] = append(combinations[input[i-1]+" "+input[i]], input[i+1])
			} else {
				combinations[input[i-1]+" "+input[i]] = append(combinations[input[i-1]+" "+input[i]], input[i+1])
			}
*/
