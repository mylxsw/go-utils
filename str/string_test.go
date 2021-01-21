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

	assert.True(t, str.HasSuffixes(s1, []string{"ld"}))
	assert.False(t, str.HasSuffixes(s1, []string{"ld1"}))
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

func TestCutoff(t *testing.T) {
	assert.Equal(t, "Hell...", str.Cutoff(4, "Hello, world"))
	assert.Equal(t, "Hello, world", str.Cutoff(100, "Hello, world"))
	assert.Equal(t, "Hello, world", str.Cutoff(len("Hello, world"), "Hello, world"))
}
