package must_test

import (
	"errors"
	"testing"

	"github.com/mylxsw/go-utils/assert"
	"github.com/mylxsw/go-utils/must"
)

func TestMust(t *testing.T) {

	var testFunc = func(showError bool) (string, error) {
		if showError {
			return "", errors.New("oops")
		}

		return "Hello, world", nil
	}

	assert.Equal(t, "Hello, world", must.Must(testFunc(false)))
	func() {
		defer func() {
			if err := recover(); err != nil {
				if err.(error).Error() != "oops" {
					t.Errorf("test failed")
				}
			} else {
				t.Errorf("test failed")
			}
		}()

		must.Must(testFunc(true))
	}()
}
