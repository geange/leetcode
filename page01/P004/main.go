package main

import "fmt"

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

/**
4. 寻找两个正序数组的中位数
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。

进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？



示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
示例 3：

输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000
示例 4：

输入：nums1 = [], nums2 = [1]
输出：1.00000
示例 5：

输入：nums1 = [2], nums2 = []
输出：2.00000


提示：

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	array := make([]int, len(nums1)+len(nums2))
	size := len(array)

	cursor := 0
	i1, i2 := 0, 0
	for i1 < len(nums1) && i2 < len(nums2) {
		if nums1[i1] < nums2[i2] {
			array[cursor] = nums1[i1]
			i1++
		} else {
			array[cursor] = nums2[i2]
			i2++
		}
		cursor++
	}

	for i1 < len(nums1) {
		array[cursor] = nums1[i1]
		i1++
		cursor++
	}

	for i2 < len(nums2) {
		array[cursor] = nums2[i2]
		i2++
		cursor++
	}

	fmt.Println(array)

	if size%2 == 0 {
		index1, index2 := size/2-1, size/2
		return float64(array[index1]+array[index2]) / 2
	}
	return float64(array[size/2])
}
