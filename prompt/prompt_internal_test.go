package prompt

import (
	"bufio"
	"io"
	"reflect"
	"strings"
	"testing"
)

type bareOptions struct {
	invalidAnswerMessage string
	answers              []rune
	reader               io.RuneReader
}

// Cannot test equality on functions, so we need to get rid of it
func stripOptions(o *AskOptions) *bareOptions {
	return &bareOptions{o.InvalidAnswerMessage, o.Answers, o.Reader}
}

func TestSetDefaults(t *testing.T) {
	question := "Please tell me your name"
	invalidAnswerMessage := "Why u do this?"
	answers := []rune{'c', 'a'}
	reader := bufio.NewReader(strings.NewReader("t"))

	failHandlerRan := false

	failHandler := func(qestion string, o *AskOptions) (rune, error) {
		failHandlerRan = true
		return 0, nil
	}

	expected := &bareOptions{
		invalidAnswerMessage: invalidAnswerMessage,
		answers:              answers,
		reader:               reader,
	}

	actual := &AskOptions{
		InvalidAnswerMessage: invalidAnswerMessage,
		Answers:              answers,
		Reader:               reader,
		FailHandlerFunc:      failHandler,
	}

	setDefaults(actual)

	actual.FailHandlerFunc(question, actual)
	if !failHandlerRan {
		t.Error("Could not call FailHandlerFunc on actual Options. FailHandlerFunc should not be modified once set")
	}

	strippedActual := stripOptions(actual)
	if !reflect.DeepEqual(expected, strippedActual) {
		t.Errorf("Expected %v got %v instead", expected, strippedActual)
	}
}

func TestSetDefaultsEmptyOptions(t *testing.T) {
	actual := &AskOptions{}

	setDefaults(actual)

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
