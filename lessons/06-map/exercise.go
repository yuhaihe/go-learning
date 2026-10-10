package mapsx

// Count 统计每个单词出现的次数。
func Count(words []string) map[string]int {
	countMap := make(map[string]int)

	for _, v := range words {
		if _, ok := countMap[v]; ok {
			countMap[v]++
		} else {
			countMap[v] = 1
		}
	}

	return countMap
}
