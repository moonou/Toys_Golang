package code

import (
	"fmt"
	"sort"
)

func Median() {
	fmt.Println(findMedianSortedArrays([]int{3, 4, 5}, []int{1, 2}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := make([]int, len(nums1)+len(nums2))
	copy(arr, nums1)
	copy(arr[len(nums1):], nums2)

	sort.Ints(arr)

	if len(arr)%2 == 0 {
		midl := len(arr) / 2
		return (float64(arr[midl]) + float64(arr[midl-1])) / 2
	} else {
		return float64(arr[len(arr)/2])
	}
	return 0
}

func mysort(nums []int) {
	for index, _ := range nums {
		for index2, _ := range nums[index+1:] {
			trueIndex := index + index2 + 1
			if nums[index] > nums[trueIndex] {
				temp := nums[trueIndex]
				nums[trueIndex] = nums[index]
				nums[index] = temp
			}
		}
	}
}
