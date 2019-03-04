package inquirer

import (
	"bufio"
	"io"
	"reflect"
	"strings"
	"testing"
)

type bareOptions struct {
	question             string
	invalidAnswerMessage string
	answers              []rune
	reader               io.RuneReader
}

// Cannot test equality on functions, so we need to get rid of it
func stripOptions(o *Options) *bareOptions {
	return &bareOptions{o.Question, o.InvalidAnswerMessage, o.Answers, o.Reader}
}

func TestSetDefaults(t *testing.T) {
	question := "Question?"
	invalidAnswerMessage := "Why u do this?"
	answers := []rune{'c', 'a'}
	reader := bufio.NewReader(strings.NewReader("t"))

	failHandlerRan := false

	failHandler := func(o *Options) (rune, error) {
		failHandlerRan = true
		return 0, nil
	}

	expected := &bareOptions{
		question:             question,
		invalidAnswerMessage: invalidAnswerMessage,
		answers:              answers,
		reader:               reader,
	}

	actual := &Options{
		Question:             question,
		InvalidAnswerMessage: invalidAnswerMessage,
		Answers:              answers,
		Reader:               reader,
		FailHandler:          failHandler,
	}

	err := setDefaults(actual)

	if err != nil {
		t.Error("setDefaults should not throw error when Question is set")
	}

	actual.FailHandler(actual)
	if !failHandlerRan {
		t.Error("Could not call FailHandler on actual Options. FailHandler should not be modified once set")
	}

	strippedActual := stripOptions(actual)
	if !reflect.DeepEqual(expected, strippedActual) {
		t.Errorf("Expected %v got %v instead", expected, strippedActual)
	}
}

func TestDefaultsEmpty(t *testing.T) {
	err := setDefaults(&Options{})
	if err == nil {
		t.Error("setDefaults should throw an error if Question is not set")
	}
}

func TestDefaultsWithQuestion(t *testing.T) {
	question := "Question?"

	actual := &Options{Question: question}

	err := setDefaults(actual)
	if err != nil {
		t.Error("setDefaults should not throw error when Question is set")
	}

	if actual.Question != question {
		t.Errorf("setDefaults should not change Question, expected %s go %s instead", question, actual.Question)
	}

	if actual.InvalidAnswerMessage == "" {
		t.Error("setDefaults should set InvalidAnswerMessage when not set. Got empty string")
	}

	if actual.Answers == nil {
		t.Error("setDefaults should set Answer when not set. Got empty string.")
	}

	if actual.Reader == nil {
		t.Error("setDefaultsq should set Reader when not set. Got empty Reader")
	}
}
