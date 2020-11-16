package sorting_test

import (
	"fmt"
	"testing"

	"github.com/mylxsw/go-utils/assert"
	"github.com/mylxsw/go-utils/sorting"
)

func TestInt64s(t *testing.T) {
	data := []int64{4121, 555, 11, 21314, 56, 0, -100}
	sorting.Int64s(data)
	assert.Equal(t, "[-100 0 11 56 555 4121 21314]", fmt.Sprintf("%v", data))
}
