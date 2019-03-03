package inquirer

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

// Options ...
type Options struct {
	Question             string
	InvalidAnswerMessage string
	Answers              []rune
	Reader               io.RuneReader
	FailHandler          func(*Options) (rune, error)
}

func setDefaults(opts *Options) error {
	if opts.Question == "" {
		return errors.New("The Question Option is mandatory")
	}

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
	return nil
}

func Ask(opts *Options) (rune, error) {
	err := setDefaults(opts)
	if err != nil {
		return 0, err
	}

	fmt.Println(opts.Question)

	char, _, err := opts.Reader.ReadRune()

	if err != nil {
		return 0, err
	}

	if !isRuneContained(char, opts.Answers) {
		fmt.Println(opts.InvalidAnswerMessage)
		return opts.FailHandler(opts)
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
