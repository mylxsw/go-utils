package str_test

import (
	"testing"

	"github.com/mylxsw/go-utils/assert"
	"github.com/mylxsw/go-utils/str"
)

func TestStringUnique(t *testing.T) {
	arr := []string{
		"aaa",
		"bbb",
		"ccc",
		"aaa",
		"ddd",
		"ccc",
	}

	assert.EqualValues(t, 4, len(str.Distinct(arr)))
}

func TestStringsContainPrefix(t *testing.T) {
	s1 := "Hello, world"
	assert.True(t, str.HasPrefixes(s1, []string{"xxxx", "yyyy", "Hell"}))
	assert.False(t, str.HasPrefixes(s1, []string{"xxxx", "yyyy", "oops"}))
}

func TestStringDiff(t *testing.T) {
	itemsA := []string{"aaa", "bbb", "ccc", "ddd"}
	itemsB := []string{"ccc", "bbb", "eee"}

	res := str.Diff(itemsA, itemsB)
	assert.Equal(t, 2, len(res))
	assert.True(t, str.In("aaa", res))
	assert.True(t, str.In("ddd", res))
}

func TestUnion(t *testing.T) {
	itemsA := []string{"aaa", "bbb", "ccc", "ddd"}
	itemsB := []string{"ccc", "bbb", "eee"}

	res := str.Union(itemsA, itemsB)
	assert.Equal(t, 5, len(res))
}
