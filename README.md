# go-utils 一些常用的 Go 工具函数

使用方式

```bash
go get -u github.com/mylxsw/go-utils
```

API 文档看这里: https://pkg.go.dev/github.com/mylxsw/go-utils

- `array` Slice 系列函数，包含 map, reduce, filter, distinct 等常用类型转换函数
- `assert`  用于在代码测试中使用，断言条件是否符合
- `chunk` 用于对数据进行分批次，目前支持对 channels 中的数据分批次获取
- `debug`  调试函数
- `diff` 对两个文本进行差异对比，输出类似于 git diff 的结果
- `failover` 用于对执行失败的函数自动重试
- `file` 一些常用的文件操作函数
- `maps` Map 系列函数，包含 map, reduct, filter 等常见的类型转换函数
- `must` 用于消除函数返回值中的 error 参数，将多返回值函数转换为单返回值
- `ternary` 模拟三元操作符

## 部分包代码示例

## array

- `array.Distinct([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6})` 

      []int{10, 2, 3, 8, 9, 4, 6}

- `array.Map([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6}, func(item int, _ int) string { return fmt.Sprintf("col_%d", item) })` 

      []string{"col_10", "col_2", "col_3", "col_8", "col_9", "col_10", "col_3", "col_4", "col_10", "col_6"}

- `array.Filter([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6}, func(item int, _ int) bool { return item > 5 })` 

      []int{10, 8, 9, 10, 10, 6}

- `array.Sort([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6}, func(v1, v2 int) bool { return v1 < v2 })` 

      []int{2, 3, 3, 4, 6, 8, 9, 10, 10, 10}

- `array.Chunks([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6}, 3)` 

      [][]int{[]int{10, 2, 3}, []int{8, 9, 10}, []int{3, 4, 10}, []int{6}}

- `array.BuildMap([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6}, func(v int, i int) (string, int) { return "#" + strconv.Itoa(i), v })` 

      map[string]int{"#0":10, "#1":2, "#2":3, "#3":8, "#4":9, "#5":10, "#6":3, "#7":4, "#8":10, "#9":6}

- `array.FromMap(map[string]int{"a": 1, "b": 3, "c": 10})` 

      []int{1, 3, 10}

- `array.FromMapKeys(map[string]int{"a": 1, "b": 3, "c": 10})` 

      []string{"a", "b", "c"}

- `array.GroupBy([]int{10, 2, 3, 8, 9, 4, 6}, func(v int) string { return strconv.Itoa(v % 2) })` 

      map[string][]int{"0":[]int{10, 2, 8, 4, 6}, "1":[]int{3, 9}}

- `array.In("c", []string{"a", "b", "c"})` 

      true

- `array.Reduce([]int{1, 2, 3, 4}, func(carry int, v int) int { return carry + v }, 0)` 

      10

- `array.Intersect([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6}, []int{4, 5, 6, 7})` 

      []int{4, 6}

- `array.Diff([]int{10, 2, 3, 8, 9, 10, 3, 4, 10, 6}, []int{4, 5, 6, 7})` 

      []int{10, 2, 3, 8, 9}

- `array.Union([]int{1, 2, 3}, []int{3, 4, 5})` 

      []int{1, 2, 3, 4, 5}

- `array.Repeat("?", 5)` 

      []string{"?", "?", "?", "?", "?"}

- `array.Reverse([]int{1, 2, 3})` 

      []int{3, 2, 1}

- `array.Shuffle([]int{1, 2, 3, 4, 5})` 

      []int{3, 1, 2, 5, 4}


## maps

- `maps.Filter(map[string]int{"a": 1, "b": 5, "c": 10}, func(v int, k string) bool { return v > 5 })` 

      map[string]int{"c":10}

- `maps.Keys(map[string]int{"a": 1, "b": 5, "c": 10})` 

      []string{"b", "c", "a"}

- `maps.Values(map[string]int{"a": 1, "b": 5, "c": 10})` 

      []int{1, 5, 10}

- `maps.Map(map[string]int{"a": 1, "b": 5, "c": 10}, func(v int, k string) string { if k == "b" {return "bbb"} else {return k + strconv.Itoa(v)} })` 

      map[string]string{"a":"a1", "b":"bbb", "c":"c10"}


## ternary

- `ternary.If(true, "yes", "no")` 

      "yes"

- `ternary.IfLazy(false, func() string { return "yes" }, func() string { return "no" })` 

      "no"

- `ternary.IfLazy2(true, func() (string, int) { return "yes", 1 }, func() (string, int) { return "no", 0 })` 

      "return (yes, 1)"
