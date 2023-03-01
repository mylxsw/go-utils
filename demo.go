package main

import (
	"fmt"
	"strconv"

	"github.com/mylxsw/go-utils/array"
	"github.com/mylxsw/go-utils/maps"
	"github.com/mylxsw/go-utils/ternary"
)

func main() {
	demoForArray()
	demoForMaps()
	demoForTernary()
}

func demoForTernary() {
	fmt.Printf("\n## ternary\n\n")
	print(`ternary.If(true, "yes", "no")`, ternary.If(true, "yes", "no"))
	print(`ternary.IfLazy(false, func() string { return "yes" }, func() string { return "no" })`, ternary.IfLazy(false, func() string { return "yes" }, func() string { return "no" }))
	{
		v1, v2 := ternary.IfLazy2(true, func() (string, int) { return "yes", 1 }, func() (string, int) { return "no", 0 })
		print(`ternary.IfLazy2(true, func() (string, int) { return "yes", 1 }, func() (string, int) { return "no", 0 })`, fmt.Sprintf("return (%s, %d)", v1, v2))
	}
}

func demoForMaps() {
	fmt.Printf("\n## maps\n\n")
	print(`maps.Filter(map[string]int{"a": 1, "b": 5, "c": 10}, func(v int, k string) bool { return v > 5 })`, maps.Filter(map[string]int{"a": 1, "b": 5, "c": 10}, func(v int, k string) bool { return v > 5 }))
	print(`maps.Keys(map[string]int{"a": 1, "b": 5, "c": 10})`, maps.Keys(map[string]int{"a": 1, "b": 5, "c": 10}))
	print(`maps.Values(map[string]int{"a": 1, "b": 5, "c": 10})`, maps.Values(map[string]int{"a": 1, "b": 5, "c": 10}))
	print(`maps.Map(map[string]int{"a": 1, "b": 5, "c": 10}, func(v int, k string) string { if k == "b" {return "bbb"} else {return k + strconv.Itoa(v)} })`, maps.Map(map[string]int{"a": 1, "b": 5, "c": 10}, func(v int, k string) string {
		if k == "b" {
			return "bbb"
		} else {
			return k + strconv.Itoa(v)
		}
	}))
}

func demoForArray() {
	fmt.Printf("\n## array\n\n")
	print(`array.Distinct([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6})`, array.Distinct([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6}))
	print(`array.Map([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6}, func(item int, _ int) string { return fmt.Sprintf("col_%d", item) })`, array.Map([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6}, func(item int, _ int) string { return fmt.Sprintf("col_%d", item) }))
	print(`array.Filter([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6}, func(item int, _ int) bool { return item > 5 })`, array.Filter([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6}, func(item int, _ int) bool { return item > 5 }))
	print(`array.Sort([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6}, func(v1, v2 int) bool { return v1 < v2 })`, array.Sort([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6}, func(v1, v2 int) bool { return v1 < v2 }))
	print(`array.Chunks([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6}, 3)`, array.Chunks([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6}, 3))
	print(`array.BuildMap([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6}, func(v int, i int) (string, int) { return "#" + strconv.Itoa(i), v })`, array.BuildMap([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6}, func(v int, i int) (string, int) { return "#" + strconv.Itoa(i), v }))
	print(`array.FromMap(map[string]int{"a": 1, "b": 3, "c": 10})`, array.FromMap(map[string]int{"a": 1, "b": 3, "c": 10}))
	print(`array.FromMapKeys(map[string]int{"a": 1, "b": 3, "c": 10})`, array.FromMapKeys(map[string]int{"a": 1, "b": 3, "c": 10}))
	print(`array.GroupBy([]int{10, 2, 3, 8, 9, 4, 6}, func(v int) string { return strconv.Itoa(v % 2) })`, array.GroupBy([]int{10, 2, 3, 8, 9, 4, 6}, func(v int) string { return strconv.Itoa(v % 2) }))
	print(`array.In("c", []string{"a", "b", "c"})`, array.In("c", []string{"a", "b", "c"}))
	print(`array.Reduce([]int{1, 2, 3, 4}, func(carry int, v int) int { return carry + v }, 0)`, array.Reduce([]int{1, 2, 3, 4}, func(carry int, v int) int { return carry + v }, 0))
	print(`array.Intersect([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6}, []int{4, 5, 6, 7})`, array.Intersect([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6}, []int{4, 5, 6, 7}))
	print(`array.Diff([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6}, []int{4, 5, 6, 7})`, array.Diff([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6}, []int{4, 5, 6, 7}))
	print(`array.Union([]int{1, 2, 3}, []int{3, 4, 5})`, array.Union([]int{1, 2, 3}, []int{3, 4, 5}))
	print(`array.Repeat("?", 5)`, array.Repeat("?", 5))
	print(`array.Reverse([]int{1, 2, 3})`, array.Reverse([]int{1, 2, 3}))
	print(`array.Shuffle([]int{1, 2, 3, 4, 5})`, array.Shuffle([]int{1, 2, 3, 4, 5}))
}

func print(title string, data interface{}) {
	fmt.Printf("- `%s` \n\n      %#v\n\n", title, data)
}
