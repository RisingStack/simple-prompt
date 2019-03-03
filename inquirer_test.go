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

	expected := 'c'
	actual, err := inquirer.Ask("Please tell me your name", []rune{'c', 'a'}, reader)

	if err != nil {
		panic(err)
	}

	if expected != actual {
		t.Error("Expected", string(expected), "got", string(actual), "instead")
	}

}

func ExampleAsk_output() {
	input := "c"

	reader := bufio.NewReader(strings.NewReader(input))

	inquirer.Ask("Please tell me your name", []rune{'c', 'a'}, reader)
	// Output: Please tell me your name
}

func ExampleAskWrongInput() {
	input := "u"

	reader := bufio.NewReader(strings.NewReader(input))

	inquirer.Ask("Please tell me your name", []rune{'c', 'a'}, reader)

	// Output:
	// Please tell me your name
	// Invalid answer, please try again
	// Please tell me your name
}
