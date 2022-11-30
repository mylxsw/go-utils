package array

import (
	"math/rand"
	"sort"
)

// Repeat 生成 count 个包含 item 值的数组
func Repeat[T any](item T, count int) []T {
	result := make([]T, count)
	for i := 0; i < count; i++ {
		result[i] = item
	}
	return result
}

// RepeatFunc 重复执行 count 次 fn，返回结果为数组
func RepeatFunc[T any](fn func() T, count int) []T {
	result := make([]T, count)
	for i := 0; i < count; i++ {
		result[i] = fn()
	}
	return result
}

// BuildMap 遍历 input 数组，使用 mapBuilder 函数（返回值为 Key, Value）创建 map
func BuildMap[T any, M any, K comparable](input []T, mapBuilder func(value T, i int) (K, M)) map[K]M {
	result := make(map[K]M)
	for i, item := range input {
		key, value := mapBuilder(item, i)
		result[key] = value
	}

	return result
}

// ToMap 将数组转换为 map
func ToMap[T any, K comparable](input []T, keyFunc func(T, int) K) map[K]T {
	m := make(map[K]T)
	for i, val := range input {
		m[keyFunc(val, i)] = val
	}

	return m
}

// FromMapKeys 将 map 的 key 转换为数组
func FromMapKeys[T any, K comparable](input map[K]T) []K {
	keys := make([]K, 0, len(input))
	for key := range input {
		keys = append(keys, key)
	}

	return keys
}

// FromMap 从 map 中提取数组
func FromMap[T any, K comparable](input map[K]T) []T {
	output := make([]T, 0, len(input))
	for _, val := range input {
		output = append(output, val)
	}

	return output
}

// Uniq remove duplicate elements from array
func Uniq[K comparable](input []K) []K {
	return Distinct(input)
}

// UniqBy remove duplicate elements from array, use keyFunc to compare
func UniqBy[K any, V comparable](input []K, keyFunc func(item K) V) []K {
	return DistinctBy(input, keyFunc)
}

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

// DistinctBy remove duplicate elements from array, use keyFunc to compare
func DistinctBy[K any, V comparable](input []K, keyFunc func(item K) V) []K {
	u := make([]K, 0, len(input))
	m := make(map[V]bool)

	for _, val := range input {
		vv := keyFunc(val)
		if _, ok := m[vv]; !ok {
			m[vv] = true
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

// AnyIn 判断 vals 中的任意元素是否在 items 中
func AnyIn[T comparable](vals []T, items []T) bool {
	for _, val := range vals {
		if In(val, items) {
			return true
		}
	}

	return false
}

// Intersect 两个数组取交集
func Intersect[T comparable](arr1 []T, arr2 []T) []T {
	tm1 := make(map[T]bool)
	for _, a1 := range arr1 {
		if _, ok := tm1[a1]; !ok {
			tm1[a1] = true
		}
	}

	res := make([]T, 0)
	for _, a2 := range arr2 {
		if _, ok := tm1[a2]; ok {
			res = append(res, a2)
		}
	}

	return res
}

// Difference 取数据 a 和 b 的差集（返回在数据 a 中，但是不在 b 中的元素）
func Difference[T comparable](a, b []T) []T {
	bm1 := make(map[T]bool)
	for _, b1 := range b {
		if _, ok := bm1[b1]; !ok {
			bm1[b1] = true
		}
	}

	res := make([]T, 0)
	for _, a1 := range a {
		if _, ok := bm1[a1]; !ok {
			res = append(res, a1)
		}
	}

	return Distinct(res)
}

// Exclude exclude all items match excepts
func Exclude[T comparable](items []T, excepts ...T) []T {
	return Filter(items, func(item T, i int) bool {
		return !In(item, excepts)
	})
}

// Filter 字符串数组过滤
func Filter[T any](items []T, predicate func(item T, index int) bool) []T {
	res := make([]T, 0)
	for i, item := range items {
		if predicate(item, i) {
			res = append(res, item)
		}
	}

	return res
}

// Map 依次对每一个元素做 mapper 操作
func Map[T any, K any](items []T, mapper func(item T, index int) K) []K {
	res := make([]K, len(items))
	for i, item := range items {
		res[i] = mapper(item, i)
	}

	return res
}

// Diff 提取 itemsA 中包含，但是 itemsB 中不存在的元素
func Diff[T comparable](itemsA []T, itemsB []T) []T {
	return Difference(itemsA, itemsB)
}

// Union 两个字符串数组合并，去重复
func Union[T comparable](itemsA []T, itemsB []T) []T {
	return Distinct(append(itemsA, itemsB...))
}

// Reduce 对数组执行 reduce 操作
func Reduce[T any, K any](data []K, cb func(carry T, item K) T, initValue T) T {
	for _, dat := range data {
		initValue = cb(initValue, dat)
	}

	return initValue
}

// ReduceWithIndex 对数组执行 reduce 操作
func ReduceWithIndex[T any, K any](data []K, cb func(carry T, item K, index int) T, initValue T) T {
	for index, dat := range data {
		initValue = cb(initValue, dat, index)
	}

	return initValue
}

// GroupBy 按照数组的某个值进行分组
func GroupBy[T any, K comparable](data []T, cb func(item T) K) map[K][]T {
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
func Each[T any](data []T, cb func(item T, index int)) {
	for index, dat := range data {
		cb(dat, index)
	}
}

// Sort 对数组进行排序
func Sort[T any](data []T, cb func(item1 T, item2 T) bool) []T {
	results := Map(data, func(item T, _ int) sortStruct {
		return sortStruct{Value: item, Compare: func(v1, v2 any) bool {
			return cb(v1.(T), v2.(T))
		}}
	})
	sort.Sort(sortStructs(results))

	return Map(results, func(item sortStruct, _ int) T { return item.Value.(T) })
}

type sortStruct struct {
	Compare func(v1, v2 any) bool
	Value   any
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
func Reverse[T any](data []T) []T {
	length := len(data)
	for i := 0; i < length/2; i++ {
		data[length-1-i], data[i] = data[i], data[length-1-i]
	}

	return data
}

// Chunks 将数组分割成多个指定长度的数组
func Chunks[T any](data []T, size int) [][]T {
	var result [][]T
	for i := 0; i < len(data); i += size {
		end := i + size
		if end > len(data) {
			end = len(data)
		}

		result = append(result, data[i:end])
	}

	return result
}

// ChunksEach 将数组分割成多个指定长度的数组，依次执行 cb 函数
func ChunksEach[T any](data []T, size int, cb func(items []T)) {
	for i := 0; i < len(data); i += size {
		end := i + size
		if end > len(data) {
			end = len(data)
		}

		cb(data[i:end])
	}
}

// Shuffle 随机打乱数组，使用前请注意需要使用 rand.Seed(time.Now().UnixNano()) 方法来初始化随机数种子
func Shuffle[T any](data []T) []T {

	res := make([]T, len(data))
	copy(res, data)
	rand.Shuffle(len(data), func(i, j int) { res[i], res[j] = res[j], res[i] })

	return res
}
