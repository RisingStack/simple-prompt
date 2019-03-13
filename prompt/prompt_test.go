package prompt_test

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/RisingStack/simple-prompt/prompt"
)

func TestAskInput(t *testing.T) {
	input := "c"

	reader := bufio.NewReader(strings.NewReader(input))

	question := "Please tell me your name"
	options := &prompt.AskOptions{
		Answers: []rune{'c', 'a'},
		Reader:  reader,
	}

	expected := 'c'

	actual, err := prompt.Ask(question, options)

	if err != nil {
		panic(err)
	}

	if expected != actual {
		t.Error("Expected", string(expected), "got", string(actual), "instead")
	}
}

func TestAskWrongInput(t *testing.T) {
	input := "u"

	reader := bufio.NewReader(strings.NewReader(input))

	failHandlerCalled := false
	question := "Please tell me your name"
	options := &prompt.AskOptions{
		Answers: []rune{'c', 'a'},
		Reader:  reader,
		FailHandlerFunc: func(question string, opts *prompt.AskOptions) (rune, error) {
			failHandlerCalled = true
			return 0, nil
		},
	}

	prompt.Ask(question, options)

	if !failHandlerCalled {
		t.Error("FailHandlerFunc has not been called")
	}
}

func ExampleAsk() {
	input := "a"

	reader := bufio.NewReader(strings.NewReader(input))

	question := "Will you marry me? [(r)efuse, (a)ccept]"
	options := &prompt.AskOptions{
		Answers: []rune{'r', 'a'},
		Reader:  reader,
	}

	response, err := prompt.Ask(question, options)
	if err != nil {
		panic(err)
	}

	switch response {
	case 'a':
		fmt.Println("You made me the happiest dog on Earth!")
	case 'o':
		fmt.Println("Woof!")
	}
	// Output:
	// Will you marry me? [(r)efuse, (a)ccept]
	// You made me the happiest dog on Earth!
}

func ExampleAsk_wrongInput() {
	input := "u"

	reader := bufio.NewReader(strings.NewReader(input))

	failHandler := func(question string, o *prompt.AskOptions) (rune, error) {
		return 0, errors.New("You only needed to press either \"c\" or \"a\", yet you chose another character. I am disappointed")
	}

	question := "Will you marry me? [(r)efuse, (a)ccept]"
	options := &prompt.AskOptions{
		InvalidAnswerMessage: "Accepted responses are \"c\" and \"a\"",
		Answers:              []rune{'c', 'a'},
		Reader:               reader,
		FailHandlerFunc:      failHandler,
	}

	response, err := prompt.Ask(question, options)

	if err != nil {
		fmt.Println(err)
	}

	switch response {
	case 'a':
		fmt.Println("You made me the happiest dog on Earth!")
	case 'o':
		fmt.Println("Woof!")
	}
	// Output:
	// Will you marry me? [(r)efuse, (a)ccept]
	// Accepted responses are "c" and "a"
	// You only needed to press either "c" or "a", yet you chose another character. I am disappointed
}
