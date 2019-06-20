package main

import (
    "fmt"
)

// func maxSubArray(nums []int) int {
// 	length := len(nums)
	
// 	if length == 0 {
// 		return 0
// 	}

// 	max, leftsum := nums[0], nums[0]
// 	for i := 1; i < length; i++ {
// 		sum := leftsum + nums[i]
// 		if nums[i] > max {
// 			max = nums[i]
// 		}
// 		if sum > max {
// 			max = sum
// 		}

// 		if leftsum <= 0 {
// 			leftsum = nums[i]
// 		} else {
// 			leftsum = sum
// 		}
// 	}

// 	return max
// }

func maxSubArray(nums []int) int {
	var leftsum int
	maxSub := nums[0]
	for _, v := range nums {
		if leftsum <= 0 {
			leftsum = v
		} else {
			leftsum += v
		}

		if leftsum > maxSub {
			maxSub = leftsum
		}
	}

	return maxSub
}

func main() {
	// nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	// nums := []int{-2,1,-3}
	nums := []int{-2,-1,-3}
	res := maxSubArray(nums)
	fmt.Println(res)
}
