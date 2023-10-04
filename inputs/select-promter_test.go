package inputs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateAnswer(t *testing.T) {
	assert := assert.New(t)

	p := newSelectPrompter(InputSpec{Options: []string{"simple", "advanced"}})

	assert.Equal(nil, p.validateAnswer("simple"))
	assert.Equal(nil, p.validateAnswer("advanced"))
	assert.NotEqual(nil, p.validateAnswer("empty"))
}
