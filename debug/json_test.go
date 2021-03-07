package debug_test

import (
	"strings"
	"testing"

	"github.com/mylxsw/go-utils/assert"
	"github.com/mylxsw/go-utils/debug"
)

func TestMustMarshalJSON(t *testing.T) {
	assert.True(t, len(strings.Split(debug.MustMarshalJSON(map[string]interface{}{
		"name": "123",
		"age":  12,
		"subelements": map[string]interface{}{
			"third-level": "oops",
		},
	}), "\n")) > 1)
}
