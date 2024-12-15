package practice

// 1480. Running Sum of 1d Array

func runningSum(nums []int) []int {
	var a []int
	for i := range nums {
		a = append(a, sumOfEle(nums, i))
	}
	return a
}

func sumOfEle(nums []int, a int) int {
	sum := 0
	for i := 0; i < a; i++ {
		sum += nums[i]
	}
	return sum
}
