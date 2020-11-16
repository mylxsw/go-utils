package assert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func True(t *testing.T, res bool, msgAndArgs ...interface{}) bool {
	return assert.True(t, res, msgAndArgs...)
}

func False(t *testing.T, res bool, msgAndArgs ...interface{}) bool {
	return assert.False(t, res, msgAndArgs...)
}

func Equal(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	return assert.Equal(t, expected, actual, msgAndArgs...)
}

func EqualValues(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	return assert.EqualValues(t, expected, actual, msgAndArgs...)
}

func NoError(t *testing.T, err error, msgAndArgs ...interface{}) bool {
	return assert.NoError(t, err, msgAndArgs...)
}
