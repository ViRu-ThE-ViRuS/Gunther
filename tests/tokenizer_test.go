package lexer_tests

import (
	"errors"
	"testing"

	"github.com/ViRu-ThE-ViRuS/Gunther/src/lexer"
	"github.com/stretchr/testify/assert"
)

func TestTokenizerBasic(t *testing.T) {
	var tests = []struct {
		input          string
		expectedOutput int // testing output length
	}{
		{"", 0},
		{"1", 1},
		{"( + 1 2 3 )", 6},
	}

	for _, test := range tests {
		output, err := lexer.Tokenize(test.input)
		assert.Equal(t, test.expectedOutput, len(output))
		assert.NoError(t, err)
	}
}

func TestTokenizerCompound(t *testing.T) {
	var tests = []struct {
		input          string
		expectedOutput int // testing input length
	}{
		{"( + 1 2 ( + 1 2 ) )", 10},
		{"( + ( - 1 2 ) 1 )", 9},
	}

	for _, test := range tests {
		output, err := lexer.Tokenize(test.input)
		assert.Equal(t, test.expectedOutput, len(output))
		assert.NoError(t, err)
	}
}

func TestTokenizerErrors(t *testing.T) {
	var tests = []struct {
		input          string
		expectedOutput int //testing output length
		err            error
	}{
		{"(+ 1 2 3 (+ 2 3))", 0, errors.New("invalid input")},
		{"(+()", 0, errors.New("invalid input")},
		{"))+", 0, errors.New("invalid input")},
		{"(1)", 0, errors.New("invalid input")},
	}

	for _, test := range tests {
		output, err := lexer.Tokenize(test.input)
		assert.Error(t, test.err, err)
		assert.Nil(t, output)
		assert.Equal(t, test.expectedOutput, len(output))
	}
}
