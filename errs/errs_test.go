package errs

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Must_should_not_panic(t *testing.T) {
	Must(1, nil)
}

func Test_Must_should_return_the_given_value(t *testing.T) {
	assert.Equal(t, 1, Must(1, nil))
	assert.Equal(t, "hello", Must("hello", nil))
}

func Test_Must_should_panic(t *testing.T) {
	assert.PanicsWithError(t, "an error", func() { Must(1, errors.New("an error")) })
}
