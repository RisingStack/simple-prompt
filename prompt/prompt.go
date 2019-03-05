// Package simple-prompt exposes provides a simple prompt for user input in CLI applications.
// while there are other fantastic packages that provide such functionality eg. as https://github.com/c-bata/go-prompt
// or https://github.com/AlecAivazis/survey most of them provide functionality one might not need while hacking together simple CLI appplications
// simple prompt aims to be quick to understand and use, or to just provide an example you can copy and paste from the parts necessary for your use-case.
package prompt

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// AskOptions ...
type AskOptions struct {
	InvalidAnswerMessage string
	Answers              []rune
	Reader               io.RuneReader
	FailHandler          func(string, *AskOptions) (rune, error)
}

func setDefaults(opts *AskOptions) {
	if opts.InvalidAnswerMessage == "" {
		opts.InvalidAnswerMessage = "Invalid answer, please try again"
	}

	if opts.Answers == nil {
		opts.Answers = []rune{'y', 'n'}
	}

	if opts.Reader == nil {
		opts.Reader = bufio.NewReader(os.Stdin)
	}

	if opts.FailHandler == nil {
		opts.FailHandler = Ask
	}
}

func Ask(question string, opts *AskOptions) (rune, error) {
	setDefaults(opts)

	fmt.Println(question)

	char, _, err := opts.Reader.ReadRune()

	if !isRuneContained(char, opts.Answers) {
		fmt.Println(opts.InvalidAnswerMessage)
		return opts.FailHandler(question, opts)
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
