package array_test

import (
	"fmt"
	"testing"

	"github.com/mylxsw/go-utils/array"
	"github.com/mylxsw/go-utils/assert"
)

type Data struct {
	Key   string
	Value string
}

func TestRepeat(t *testing.T) {
	data := array.Repeat(1, 10)
	assert.Equal(t, 10, len(data))
	assert.Equal(t, 1, data[0])
	assert.Equal(t, 1, data[5])
	assert.Equal(t, 1, data[9])
}

func TestBuildMap(t *testing.T) {
	data := []Data{
		{Key: "abc", Value: "123"},
		{Key: "def", Value: "456"},
		{Key: "vja", Value: "789"},
		{Key: "lmn", Value: "101"},
	}

	res := array.BuildMap(data, func(item Data, i int) (string, string) { return item.Key + item.Value, item.Value })
	assert.Equal(t, 4, len(res))
	assert.Equal(t, "123", res["abc123"])
	assert.Equal(t, "101", res["lmn101"])
}

func TestToMap(t *testing.T) {
	data := []Data{
		{Key: "abc", Value: "123"},
		{Key: "def", Value: "456"},
		{Key: "vja", Value: "789"},
		{Key: "lmn", Value: "101"},
	}

	dataMap := array.ToMap(data, func(dat Data, i int) string { return dat.Key })
	assert.Equal(t, 4, len(dataMap))
	assert.Equal(t, "123", dataMap["abc"].Value)
	assert.Equal(t, "101", dataMap["lmn"].Value)
}

func TestFromMap(t *testing.T) {
	data := map[string]Data{
		"abc": {Key: "abc", Value: "123"},
		"def": {Key: "def", Value: "456"},
		"vja": {Key: "vja", Value: "789"},
		"lmn": {Key: "lmn", Value: "101"},
	}

	dataArray := array.FromMap(data)
	assert.Equal(t, 4, len(dataArray))

	keyArray := array.FromMapKeys(data)
	assert.Equal(t, 4, len(keyArray))
}

func TestDistinct(t *testing.T) {
	data := []string{"abc", "def", "124", "abc"}
	assert.Equal(t, 3, len(array.Distinct(
		data)))
}

func TestFilter(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	assert.EqualValues(t, "[2 4 6]", fmt.Sprintf("%v", array.Filter(data, func(item int, i int) bool { return item%2 == 0 })))

	type Data struct {
		Value string
	}

	data2 := []Data{
		{Value: "abc"},
		{Value: "def"},
		{Value: "vja"},
	}

	assert.EqualValues(t, "[{abc} {vja}]", fmt.Sprintf("%v", array.Filter(data2, func(item Data, i int) bool { return item.Value == "abc" || item.Value == "vja" })))
}

func TestAnyIn(t *testing.T) {
	data := []string{"abc", "def", "124", "abc"}
	assert.Equal(t, true, array.AnyIn([]string{"abc", "dx"}, data))
}

func TestIntersect(t *testing.T) {
	data1 := []string{"abc", "def", "123", "#85"}
	data2 := []string{"#86", "#85", "abc", "xxx", "yyy"}

	assert.EqualValues(t, "[#85 abc]", fmt.Sprintf("%v", array.Intersect(data1, data2)))
}

func TestDifferrnce(t *testing.T) {
	data1 := []string{"abc", "def", "123", "#85", "123"}
	data2 := []string{"#86", "#85", "abc", "xxx", "yyy"}

	assert.EqualValues(t, "[def 123]", fmt.Sprintf("%v", array.Difference(data1, data2)))
}

func TestMap(t *testing.T) {
	type Data struct {
		Value string
	}

	data := []string{"abc", "def", "oops"}
	for i, item := range array.Map(data, func(item string, _ int) Data { return Data{Value: item} }) {
		assert.EqualValues(t, data[i], item.Value)
	}
}

func TestReduce(t *testing.T) {
	{
		data := []int{1, 2, 3, 4, 5, 6, 7}
		assert.EqualValues(t, 28, array.Reduce(data, func(carry int, item int) int { return carry + item }, 0))
	}

	{
		data := []string{"abc", "def", "oops"}
		assert.EqualValues(t, 10, array.Reduce(data, func(carry int, item string) int {
			return carry + len(item)
		}, 0))
	}
}

func TestGroupBy(t *testing.T) {
	type Data struct {
		Value    string
		Category string
	}

	data := []Data{
		{Value: "abc", Category: "seq"},
		{Value: "def", Category: "seq"},
		{Value: "vja", Category: "random"},
		{Value: "lmn", Category: "seq"},
		{Value: "tux", Category: "random"},
	}

	result := array.GroupBy(data, func(item Data) string { return item.Category })
	assert.Equal(t, 2, len(result))
}

func TestEach(t *testing.T) {
	data := []string{"abc", "def", "oops"}

	count := 0
	array.Each(data, func(item string, i int) {
		count++
		assert.Equal(t, true, array.In(item, data))
	})

	assert.Equal(t, len(data), count)
}

func TestSort(t *testing.T) {
	data := []int{1, 5, 6, 63, 12, 0, 99, 184, 23}
	assert.EqualValues(t, "[184 99 63 23 12 6 5 1 0]", fmt.Sprintf("%v", array.Sort(data, func(item1, item2 int) bool { return item1 > item2 })))
}

func TestReverse(t *testing.T) {
	data := []int{1, 2, 3, 4}
	assert.EqualValues(t, "[4 3 2 1]", fmt.Sprintf("%v", array.Reverse(data)))
}

func TestChunks(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.EqualValues(t, "[[1 2 3 4] [5 6 7 8] [9 10]]", fmt.Sprintf("%v", array.Chunks(data, 4)))
}

func TestChunksEach(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	count := 0
	array.ChunksEach(data, 4, func(chunk []int) {
		count++
		assert.True(t, len(chunk) <= 4)
	})

	assert.Equal(t, 3, count)
}

func TestShuffle(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array.Shuffle(data))
	fmt.Println(array.Shuffle(data))
	fmt.Println(array.Shuffle(data))
	fmt.Println(array.Shuffle(data))
	fmt.Println(data)

	assert.EqualValues(t, 10, len(array.Shuffle(data)))
}

func TestDistinctBy(t *testing.T) {
	type Data struct {
		Value    string
		Category string
		Items    []string
	}

	data := []Data{
		{Value: "abc", Category: "seq"},
		{Value: "def", Category: "seq"},
		{Value: "vja", Category: "random"},
	}

	res := array.DistinctBy(data, func(item Data) string { return item.Category })
	assert.EqualValues(t, "[{abc seq []} {vja random []}]", fmt.Sprintf("%v", res))
}
