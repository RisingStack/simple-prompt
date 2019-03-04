package inquirer_test

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/Shadowbeetle/inquirer"
)

func TestAskInput(t *testing.T) {
	input := "c"

	reader := bufio.NewReader(strings.NewReader(input))

	options := &inquirer.Options{
		Question: "Please tell me your name",
		Answers:  []rune{'c', 'a'},
		Reader:   reader,
	}

	expected := 'c'

	actual, err := inquirer.Ask(options)

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

	options := &inquirer.Options{
		Question: "Please tell me your name",
		Answers:  []rune{'c', 'a'},
		Reader:   reader,
		FailHandler: func(opts *inquirer.Options) (rune, error) {
			failHandlerCalled = true
			return 0, nil
		},
	}

	inquirer.Ask(options)

	if !failHandlerCalled {
		t.Error("FailHandler has not been called")
	}
}

func ExampleAsk() {
	input := "a"

	reader := bufio.NewReader(strings.NewReader(input))

	options := &inquirer.Options{
		Question: "Will you marry me? [(r)efuse, (a)ccept]",
		Answers:  []rune{'r', 'a'},
		Reader:   reader,
	}

	response, err := inquirer.Ask(options)
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

	failHandler := func(o *inquirer.Options) (rune, error) {
		return 0, errors.New("You only needed to press either \"c\" or \"a\", yet you chose another character. I am disappointed")
	}

	options := &inquirer.Options{
		Question:             "Will you marry me? [(r)efuse, (a)ccept]",
		InvalidAnswerMessage: "Accepted responses are \"c\" and \"a\"",
		Answers:              []rune{'c', 'a'},
		Reader:               reader,
		FailHandler:          failHandler,
	}

	response, err := inquirer.Ask(options)

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
