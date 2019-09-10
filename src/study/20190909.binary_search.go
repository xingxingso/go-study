// xGolang Binary Search Algorithm
// https://flaviocopes.com/golang-algorithms-binary-search/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 6
	fmt.Println(arr, target)
	fmt.Println(LinearSearch(arr, target))
	fmt.Println(MyBinarySearch(arr, target))

	i := SortSearch(len(arr)-1, func (i int) bool {return arr[i] >= target})
	fmt.Println(i)
	if i >= 0 && i < len(arr) {
		fmt.Println(arr[i] == target)
	} else {
		fmt.Println(false)
	}
}

func LinearSearch(arr []int, target int) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

// my test
func MyBinarySearch(arr []int, target int) int {
	left,right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left) / 2
		if target == arr[mid] {
			return mid
		}

		if target > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
	
// sort.Search
func SortSearch(n int, f func(int) bool) int {
	i, j := 0, n
	for i < j {
		h := i + (j-i)/2
		if !f(h) {
			i = h+1
		} else {
			j = h
		}
	}
	return i
}
// sort.Search
// func Search(n int, f func(int) bool) int {
//     // Define f(-1) == false and f(n) == true.
//     // Invariant: f(i-1) == false, f(j) == true.
//     i, j := 0, n
//     for i < j {
//         h := i + (j-i)/2 // avoid overflow when computing h
//         // i ≤ h < j
//         if !f(h) {
//             i = h + 1 // preserves f(i-1) == false
//         } else {
//             j = h // preserves f(j) == true
//         }
//     }
//     // i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
//     return i
// }
