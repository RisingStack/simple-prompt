package inquirer_test

import (
	"bufio"
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

func ExampleAsk_output() {
	input := "c"

	reader := bufio.NewReader(strings.NewReader(input))

	options := &inquirer.Options{
		Question: "Please tell me your name",
		Answers:  []rune{'c', 'a'},
		Reader:   reader,
	}

	inquirer.Ask(options)
	// Output: Please tell me your name
}

func ExampleAskWrongInput() {
	input := "u"

	reader := bufio.NewReader(strings.NewReader(input))

	options := &inquirer.Options{
		Question: "Please tell me your name",
		Answers:  []rune{'c', 'a'},
		Reader:   reader,
	}

	inquirer.Ask(options)

	// Output:
	// Please tell me your name
	// Invalid answer, please try again
	// Please tell me your name
}
