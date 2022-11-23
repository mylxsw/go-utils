package str

import "strings"

// Distinct remove duplicate elements from array
func Distinct(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

// In 判断元素是否在字符串数组中
func In(val string, items []string) bool {
	for _, item := range items {
		if item == val {
			return true
		}
	}

	return false
}

// InIgnoreCase 判断元素是否在字符串数组中
func InIgnoreCase(val string, items []string) bool {
	for _, item := range items {
		if strings.EqualFold(val, item) {
			return true
		}
	}

	return false
}

// HasPrefixes 判断字符串是否以指定的前缀开始
func HasPrefixes(val string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(val, prefix) {
			return true
		}
	}

	return false
}

// HasSuffixes 判断字符串是否以指定的后缀结束
func HasSuffixes(val string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(val, suffix) {
			return true
		}
	}

	return false
}

// Exclude exclude all items match excepts
func Exclude(items []string, excepts ...string) []string {
	return Filter(items, func(item string) bool {
		return !In(item, excepts)
	})
}

// Filter 字符串数组过滤
func Filter(items []string, filter func(item string) bool) []string {
	res := make([]string, 0)
	for _, item := range items {
		if filter(item) {
			res = append(res, item)
		}
	}

	return res
}

// Map 依次对字符串数组中每一个元素做 mapper 操作
func Map(items []string, mapper func(item string) string) []string {
	res := make([]string, len(items))
	for i, item := range items {
		res[i] = mapper(item)
	}

	return res
}

// FilterEmpty 过滤掉字符串数组中的空元素
func FilterEmpty(items []string) []string {
	return Filter(items, func(item string) bool {
		return item != ""
	})
}

// Diff 提取 itemsA 中包含，但是 itemsB 中不存在的元素
func Diff(itemsA []string, itemsB []string) []string {
	res := make([]string, 0)
	for _, item := range itemsA {
		if In(item, itemsB) {
			continue
		}

		res = append(res, item)
	}

	return res
}

// Union 两个字符串数组合并，去重复
func Union(itemsA []string, itemsB []string) []string {
	return Distinct(append(itemsA, itemsB...))
}

// Cutoff 字符串截取，如果截取了部分，后面自动追加 ...
func Cutoff(maxLen int, val string) string {
	valRune := []rune(strings.Trim(val, " \n"))

	valLen := len(valRune)
	if valLen <= maxLen {
		return string(valRune)
	}

	return string(valRune[0:maxLen]) + "..."
}
