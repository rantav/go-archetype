package inputs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBeNiceAndTryToConvert(t *testing.T) {
	assert := assert.New(t)

	p := newYesNoPrompter(InputSpec{})

	assert.Equal("true", p.beNiceAndTryToConvert("yes"))
	assert.Equal("true", p.beNiceAndTryToConvert("true"))

	assert.Equal("false", p.beNiceAndTryToConvert("no"))
	assert.Equal("false", p.beNiceAndTryToConvert("false"))
}
