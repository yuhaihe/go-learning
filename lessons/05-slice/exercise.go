package slicex

// Unique 去掉重复元素，保留第一次出现的顺序。
// 提示见 README：range 遍历 + map 记录是否见过 + append 结果。
func Unique(nums []int) []int {
	arrs := []int{}
	// map写法
	// arrsMap := make(map[int]bool)

	// for _, num := range nums {
	// 	if arrsMap[num] {
	// 		continue
	// 	}
	// 	arrs = append(arrs, num)
	// 	arrsMap[num] = true
	// }

	// 双层for循环
	for _, num := range nums {
		var isExit = false
		for _, arr := range arrs {
			if arr == num {
				isExit = true
				break
			}
		}

		if !isExit {
			arrs = append(arrs, num)
		}
	}
	return arrs
}
