package assert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func True(t *testing.T, res bool, msgAndArgs ...interface{}) {
	assert.True(t, res, msgAndArgs...)
}

func False(t *testing.T, res bool, msgAndArgs ...interface{}) {
	assert.False(t, res, msgAndArgs...)
}

func Equal(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	assert.Equal(t, expected, actual, msgAndArgs...)
}

func EqualValues(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	assert.EqualValues(t, expected, actual, msgAndArgs...)
}
