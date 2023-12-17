package sfv

import (
	"errors"
	"fmt"
	"github.com/skirkyn/dcw/cmd/controller/generators/gerrorrs"
	"testing"
)

var expectedDecimal = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

var expectedIncorrectLengthError error = gerrorrs.NewIncorrectResultLength()
var expectedIncorrectVocabularyLengthError error = gerrorrs.NewIncorrectVocabularyLength()
var expectedIncorrectFormatter = gerrorrs.NewIncorrectFormatter()
var expectedCustomNotSupported = gerrorrs.NewCustomNotSupported()

func TestForCustomSuccess(t *testing.T) {

	res, err := ForCustom(8, vocabularyCharacters[Decimals], Simple)

	if err != nil {
		t.Error(err.Error())
	}
	state := res.state
	config := state.Config

	for i := 0; i < len(state.CurrentPositions); i++ {
		if state.CurrentPositions[i] != 0 {
			fmt.Println(state.CurrentPositions)
			t.Errorf("expected current position at %d to be 0, got %d", i, state.CurrentPositions[i])
		}
	}

	if config.ResultLength != 8 {
		t.Errorf("expected result length 8, got %d", config.ResultLength)
	}

	for i := 0; i < max(len(config.Vocabulary), len(expectedDecimal)); i++ {
		if (config.Vocabulary)[i] != expectedDecimal[i] {
			t.Errorf("expected rune at the position %d to be %c, got %c", i, expectedDecimal[i], config.Vocabulary[i])
		}
	}
	if config.Formatter != Simple {
		t.Errorf("expected simple formatted got %d", config.Formatter)
	}

}

func TestForCustomError(t *testing.T) {

	_, err := ForCustom(0, vocabularyCharacters[Decimals], Simple)

	if err == nil {
		t.Error("expected error")
	}

	if !errors.Is(err, expectedIncorrectLengthError) {
		t.Errorf("expected error %s, got %s", expectedIncorrectLengthError.Error(), err.Error())
	}

	_, err = ForCustom(1, []rune{}, Simple)
	if err == nil {
		t.Error("expected error")
	}

	if !errors.Is(err, expectedIncorrectVocabularyLengthError) {
		t.Errorf("expected error %s, got %s", expectedIncorrectVocabularyLengthError.Error(), err.Error())
	}
	_, err = ForCustom(2, vocabularyCharacters[Decimals], 3)
	if err == nil {
		t.Error("expected error")
	}

	if !errors.Is(err, expectedIncorrectFormatter) {
		t.Errorf("expected error %s, got %s", expectedIncorrectFormatter.Error(), err.Error())
	}

}

func TestForStandardSuccess(t *testing.T) {

	res, err := ForStandard(Decimals, 8, Simple)

	if err != nil {
		t.Error(err.Error())
	}
	state := res.state
	config := state.Config

	for i := 0; i < len(state.CurrentPositions); i++ {
		if state.CurrentPositions[i] != 0 {
			fmt.Println(state.CurrentPositions)
			t.Errorf("expected current position at %d to be 0, got %d", i, state.CurrentPositions[i])
		}
	}
	if config.ResultLength != 8 {
		t.Errorf("expected result length 8, got %d", config.ResultLength)
	}

	for i := 0; i < max(len(config.Vocabulary), len(expectedDecimal)); i++ {
		if (config.Vocabulary)[i] != expectedDecimal[i] {
			t.Errorf("expected rune at the position %d to be %c, got %c", i, expectedDecimal[i], (config.Vocabulary)[i])
		}
	}
	if config.Formatter != Simple {
		t.Errorf("expected simple formatted got %d", config.Formatter)
	}

}

func TestForStandardError(t *testing.T) {

	_, err := ForStandard(Custom, 8, Simple)

	if err == nil {
		t.Error("expected error")
	}

	if !errors.Is(err, expectedCustomNotSupported) {
		t.Errorf("expected error %s, got %s", expectedCustomNotSupported.Error(), err.Error())
	}

}

func TestRecalculatePositions(t *testing.T) {
	res, _ := ForStandard(Decimals, 8, Simple)
	p, err := res.recalculatePositions(5)
	validatePositions(t, err, p, []int{0, 0, 0, 0, 0, 0, 0, 5})

	res, _ = ForStandard(Decimals, 8, Simple)
	p, err = res.recalculatePositions(16)
	validatePositions(t, err, p, []int{0, 0, 0, 0, 0, 0, 1, 6})

	res, _ = ForStandard(Decimals, 8, Simple)
	p, err = res.recalculatePositions(16)
	validatePositions(t, err, p, []int{0, 0, 0, 0, 0, 0, 0, 5})
}

func validatePositions(t *testing.T, err error, actual []int, expected []int) {
	if err != nil {
		t.Error(err.Error())
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			fmt.Println(expected)
			fmt.Println(actual)
			t.Errorf("different positions at %d, actual %d, expected: %d", i, actual[i], expected[i])
		}
	}
}
