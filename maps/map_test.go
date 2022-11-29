package maps

import (
	"fmt"
	"testing"

	"github.com/mylxsw/go-utils/assert"
)

func TestKeys(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	keys := Keys(m)
	if len(keys) != 3 {
		t.Errorf("len(keys) = %d, want 3", len(keys))
	}
}

func TestValues(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	values := Values(m)
	if len(values) != 3 {
		t.Errorf("len(values) = %d, want 3", len(values))
	}
}

func TestMap(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	mapped := Map(m, func(item int, key string) string {
		return fmt.Sprintf("item-%s-%d", key, item)
	})

	if len(mapped) != 3 {
		t.Errorf("len(mapped) = %d, want 3", len(mapped))
	}

	assert.Equal(t, "item-one-1", mapped["one"])
	assert.Equal(t, "item-two-2", mapped["two"])
	assert.Equal(t, "item-three-3", mapped["three"])
}

func TestFilter(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	filtered := Filter(m, func(item int, key string) bool {
		return item > 1 && key != "three"
	})

	if len(filtered) != 1 {
		t.Errorf("len(filtered) = %d, want 1", len(filtered))
	}

	assert.Equal(t, 2, filtered["two"])
}

func TestEach(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	var total int
	Each(m, func(item int, key string) {
		total += item
		assert.True(t, key == "one" || key == "two" || key == "three")
	})

	assert.Equal(t, 6, total)
}

func TestOrderedEach(t *testing.T) {
	m := map[string]int{
		"1-one":   1,
		"2-two":   2,
		"3-three": 3,
	}

	var lastValue, index int
	keys := []string{"1-one", "2-two", "3-three"}
	OrderedEach(m, func(item int, key string) {
		assert.Equal(t, lastValue+1, item)
		assert.Equal(t, keys[index], key)
		lastValue = item
		index++
	})
}

func TestOrderedEachBy(t *testing.T) {
	m := map[string]int{
		"1-one":   1,
		"2-two":   2,
		"3-three": 3,
	}

	var orderBy = func(k1, k2 string) bool {
		return k1 > k2
	}

	keys := []string{"3-three", "2-two", "1-one"}
	var index int
	OrderedEachBy(m, func(item int, key string) {
		assert.Equal(t, keys[index], key)
		index++
	}, orderBy)
}

func TestOrderedEachByValue(t *testing.T) {
	m := map[string]int{
		"1-one":   1,
		"2-two":   2,
		"3-three": 3,
	}

	var orderBy = func(v1, v2 int) bool {
		return v1 > v2
	}

	keys := []string{"3-three", "2-two", "1-one"}
	var index int
	OrderedEachByValue(m, func(item int, key string) {
		assert.Equal(t, keys[index], key)
		index++
	}, orderBy)
}

func TestReduce(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	assert.Equal(t, 106, Reduce(m, func(carry int, item int) int { return carry + item }, 100))
}

func TestReduceWithKey(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	assert.Equal(t, 106, ReduceWithKey(m, func(carry int, item int, key string) int { return carry + item }, 100))
}

func TestOrderedReduceWithKey(t *testing.T) {
	m := map[string]int{
		"1-one":   1,
		"2-two":   2,
		"3-three": 3,
	}

	assert.Equal(t, "init,1-one,2-two,3-three", OrderedReduceWithKey(m, func(carry string, item int, key string) string { return carry + "," + key }, "init"))
}
