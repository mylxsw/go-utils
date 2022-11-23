package maps

import "github.com/mylxsw/go-utils/array"

// Keys 将 map 的 key 转换为数组
func Keys[T any, K comparable](input map[K]T) []K {
	keys := make([]K, 0, len(input))
	for key := range input {
		keys = append(keys, key)
	}

	return keys
}

// OrderedKeys 将 map 的 key 转换为数组，并按照 key 的顺序进行排序
func OrderedKeys[K Ordered, V any](items map[K]V) []K {
	return array.Sort(Keys(items), func(k1, k2 K) bool { return k1 < k2 })
}

// OrderedKeys 将 map 的 key 转换为数组，并按照 key 的顺序进行排序（倒序）
func OrderedKeysReverse[K Ordered, V any](items map[K]V) []K {
	return array.Sort(Keys(items), func(k1, k2 K) bool { return k1 > k2 })
}

// OrderedKeysBy 将 map 的 key 转换为数组，并按照 key 使用 orderBy 函数排序
func OrderedKeysBy[K Ordered, V any](items map[K]V, orderBy func(k1, k2 K) bool) []K {
	return array.Sort(Keys(items), orderBy)
}

// FromMap 从 map 中提取数组
func Values[T any, K comparable](input map[K]T) []T {
	output := make([]T, 0, len(input))
	for _, val := range input {
		output = append(output, val)
	}

	return output
}

// Map 依次对每一个元素做 mapper 操作
func Map[T any, K comparable, V any](items map[K]V, mapper func(item V) T) map[K]T {
	res := make(map[K]T, len(items))
	for k, v := range items {
		res[k] = mapper(v)
	}

	return res
}

// MapWithKey 依次对每一个元素做 mapper 操作
func MapWithKey[T any, K comparable, V any](items map[K]V, mapper func(item V, key K) T) map[K]T {
	res := make(map[K]T, len(items))
	for k, v := range items {
		res[k] = mapper(v, k)
	}

	return res
}

// Filter 对每一个元素做 filter 操作
func Filter[K comparable, V any](items map[K]V, predicate func(item V) bool) map[K]V {
	res := make(map[K]V, 0)
	for k, v := range items {
		if predicate(v) {
			res[k] = v
		}
	}

	return res
}

// Filter 对每一个元素做 filter 操作
func FilterWithKey[K comparable, V any](items map[K]V, predicate func(item V, key K) bool) map[K]V {
	res := make(map[K]V, 0)
	for k, v := range items {
		if predicate(v, k) {
			res[k] = v
		}
	}

	return res
}

// Each 对每一个元素做 consumer 操作
func Each[K comparable, V any](items map[K]V, consumer func(item V)) {
	for _, v := range items {
		consumer(v)
	}
}

// EachWithKey 对每一个元素做 consumer 操作
func EachWithKey[K comparable, V any](items map[K]V, consumer func(item V, key K)) {
	for k, v := range items {
		consumer(v, k)
	}
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// OrderedEach 对每一个元素做 consumer 操作，对 Map 的 Key 进行排序后再遍历
func OrderedEach[K Ordered, V any](items map[K]V, consumer func(item V)) {
	sortedKeys := array.Sort(Keys(items), func(k1, k2 K) bool { return k1 < k2 })
	for _, k := range sortedKeys {
		consumer(items[k])
	}
}

// OrderedEachWithKey 对每一个元素做 consumer 操作，对 Map 的 Key 进行排序后再遍历
func OrderedEachWithKey[K Ordered, V any](items map[K]V, consumer func(item V, key K)) {
	sortedKeys := array.Sort(Keys(items), func(k1, k2 K) bool { return k1 < k2 })
	for _, k := range sortedKeys {
		consumer(items[k], k)
	}
}

// OrderedEachBy 对每一个元素做 consumer 操作，对 Map 的 Key 进行排序后再遍历，通过 orderBy 指定排序规则
func OrderedEachBy[K Ordered, V any](items map[K]V, consumer func(item V), orderBy func(k1, k2 K) bool) {
	sortedKeys := array.Sort(Keys(items), orderBy)
	for _, k := range sortedKeys {
		consumer(items[k])
	}
}

// OrderedEachWithKeyBy 对每一个元素做 consumer 操作，对 Map 的 Key 进行排序后再遍历，通过 orderBy 指定排序规则
func OrderedEachWithKeyBy[K Ordered, V any](items map[K]V, consumer func(item V, key K), orderBy func(k1, k2 K) bool) {
	sortedKeys := array.Sort(Keys(items), orderBy)
	for _, k := range sortedKeys {
		consumer(items[k], k)
	}
}

// OrderedEachByValue 对每一个元素做 consumer 操作，对 Map 的 Key 进行排序后再遍历，通过 orderBy 指定排序规则
func OrderedEachByValue[K Ordered, V any](items map[K]V, consumer func(item V), orderBy func(v1, v2 V) bool) {
	sortedKeys := array.Sort(Keys(items), func(k1, k2 K) bool { return orderBy(items[k1], items[k2]) })
	for _, k := range sortedKeys {
		consumer(items[k])
	}
}

// OrderedEachWithKeyByValue 对每一个元素做 consumer 操作，对 Map 的 Key 进行排序后再遍历，通过 orderBy 指定排序规则
func OrderedEachWithKeyByValue[K Ordered, V any](items map[K]V, consumer func(item V, key K), orderBy func(v1, v2 V) bool) {
	sortedKeys := array.Sort(Keys(items), func(k1, k2 K) bool { return orderBy(items[k1], items[k2]) })
	for _, k := range sortedKeys {
		consumer(items[k], k)
	}
}

// Reduce 对 map 执行 reduce 操作
func Reduce[K comparable, V any, R any](items map[K]V, reducer func(result R, item V) R, init R) R {
	res := init
	for _, v := range items {
		res = reducer(res, v)
	}

	return res
}

// ReduceWithKey 对 map 执行 reduce 操作
func ReduceWithKey[K comparable, V any, R any](items map[K]V, reducer func(result R, item V, key K) R, init R) R {
	res := init
	for k, v := range items {
		res = reducer(res, v, k)
	}

	return res
}

// OrderedReduce 对 map 执行 reduce 操作
func OrderedReduce[K Ordered, V any, R any](items map[K]V, reducer func(result R, item V) R, init R) R {
	sortedKeys := array.Sort(Keys(items), func(k1, k2 K) bool { return k1 < k2 })
	res := init
	for _, k := range sortedKeys {
		res = reducer(res, items[k])
	}

	return res
}

// OrderedReduceWithKey 对 map 执行 reduce 操作
func OrderedReduceWithKey[K Ordered, V any, R any](items map[K]V, reducer func(result R, item V, key K) R, init R) R {
	sortedKeys := array.Sort(Keys(items), func(k1, k2 K) bool { return k1 < k2 })
	res := init
	for _, k := range sortedKeys {
		res = reducer(res, items[k], k)
	}

	return res
}

// OrderedReduceBy 对 map 执行 reduce 操作
func OrderedReduceBy[K Ordered, V any, R any](items map[K]V, reducer func(result R, item V) R, init R, orderBy func(k1, k2 K) bool) R {
	sortedKeys := array.Sort(Keys(items), orderBy)
	res := init
	for _, k := range sortedKeys {
		res = reducer(res, items[k])
	}

	return res
}

// OrderedReduceReverse 对 map 执行 reduce 操作
func OrderedReduceReverse[K Ordered, V any, R any](items map[K]V, reducer func(result R, item V) R, init R) R {
	sortedKeys := array.Sort(Keys(items), func(k1, k2 K) bool { return k1 > k2 })
	res := init
	for _, k := range sortedKeys {
		res = reducer(res, items[k])
	}

	return res
}

// OrderedReduceWithKeyReverse 对 map 执行 reduce 操作
func OrderedReduceWithKeyReverse[K Ordered, V any, R any](items map[K]V, reducer func(result R, item V, key K) R, init R) R {
	sortedKeys := array.Sort(Keys(items), func(k1, k2 K) bool { return k1 > k2 })
	res := init
	for _, k := range sortedKeys {
		res = reducer(res, items[k], k)
	}

	return res
}
