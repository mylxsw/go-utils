package array

import "sort"

// Distinct remove duplicate elements from array
func Distinct[K comparable](input []K) []K {
	u := make([]K, 0, len(input))
	m := make(map[K]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

// In 判断元素是否在数组中
func In[T comparable](val T, items []T) bool {
	for _, item := range items {
		if item == val {
			return true
		}
	}

	return false
}

// Exclude exclude all items match excepts
func Exclude[T comparable](items []T, excepts ...T) []T {
	return Filter(items, func(item T) bool {
		return !In(item, excepts)
	})
}

// Filter 字符串数组过滤
func Filter[T comparable](items []T, filter func(item T) bool) []T {
	res := make([]T, 0)
	for _, item := range items {
		if filter(item) {
			res = append(res, item)
		}
	}

	return res
}

// Map 依次对每一个元素做 mapper 操作
func Map[T interface{}, K interface{}](items []T, mapper func(item T) K) []K {
	res := make([]K, 0)
	for _, item := range items {
		res = append(res, mapper(item))
	}

	return res
}

// Diff 提取 itemsA 中包含，但是 itemsB 中不存在的元素
func Diff[T comparable](itemsA []T, itemsB []T) []T {
	res := make([]T, 0)
	for _, item := range itemsA {
		if In(item, itemsB) {
			continue
		}

		res = append(res, item)
	}

	return res
}

// Union 两个字符串数组合并，去重复
func Union[T comparable](itemsA []T, itemsB []T) []T {
	return Distinct(append(itemsA, itemsB...))
}

// Reduce 对数组执行 reduce 操作
func Reduce[T interface{}, K interface{}](data []K, cb func(carry T, item K) T, initValue T) T {
	for _, dat := range data {
		initValue = cb(initValue, dat)
	}

	return initValue
}

// GroupBy 按照数组的某个值进行分组
func GroupBy[T interface{}, K comparable](data []T, cb func(item T) K) map[K][]T {
	results := make(map[K][]T)
	for _, dat := range data {
		k := cb(dat)
		if _, ok := results[k]; !ok {
			results[k] = make([]T, 0)
		}

		results[k] = append(results[k], dat)
	}

	return results
}

// Each 遍历data，依次执行 cb 函数
func Each[T interface{}](data []T, cb func(item T)) {
	for _, dat := range data {
		cb(dat)
	}
}

// Sort 对数组进行排序
func Sort[T interface{}](data []T, cb func(item1 T, item2 T) bool) []T {
	results := Map(data, func(item T) sortStruct {
		return sortStruct{Value: item, Compare: func(v1, v2 interface{}) bool {
			return cb(v1.(T), v2.(T))
		}}
	})
	sort.Sort(sortStructs(results))

	return Map(results, func(item sortStruct) T { return item.Value.(T) })
}

type sortStruct struct {
	Compare func(v1, v2 interface{}) bool
	Value   interface{}
}

type sortStructs []sortStruct

func (s sortStructs) Len() int {
	return len(s)
}

func (s sortStructs) Less(i, j int) bool {
	return s[i].Compare(s[i].Value, s[j].Value)
}

func (s sortStructs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Reverse 数组逆序
func Reverse[T interface{}](data []T) []T {
	length := len(data)
	for i := 0; i < length/2; i++ {
		data[length-1-i], data[i] = data[i], data[length-1-i]
	}

	return data
}
