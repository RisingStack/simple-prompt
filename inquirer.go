package inquirer

import (
	"fmt"
	"io"
)

func Confirm(question string, reader io.RuneReader) (rune, error) {
	return Ask(question, []rune{'y', 'n'}, reader)
}

// Pass os.Stdin to read from stdin
func Ask(question string, answers []rune, reader io.RuneReader) (rune, error) {

	fmt.Println(question)
	char, _, err := reader.ReadRune()

	if err != nil {
		return 0, err
	}

	if !isRuneContained(char, answers) {
		fmt.Println("Invalid answer, please try again")
		return Ask(question, answers, reader)
	}

	return char, err
}

func isRuneContained(r rune, runeSlice []rune) bool {
	for _, item := range runeSlice {
		if item == r {
			return true
		}
	}
	return false
}
