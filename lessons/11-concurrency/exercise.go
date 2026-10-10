package parallel

// Sum 把切片切成两半，各自用 goroutine 求和，再把结果加起来。空切片返回 0。
func Sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	mid := len(nums) / 2
	ch := make(chan int, 2)
	go func() {
		n := 0
		for i := 0; i < mid; i++ {
			n += nums[i]
		}
		ch <- n
	}()
	go func() {
		n := 0
		for i := mid; i < len(nums); i++ {
			n += nums[i]
		}
		ch <- n
	}()

	return <-ch + <-ch
}
