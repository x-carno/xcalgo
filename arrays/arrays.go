package arrays

import "sort"

// just return any repeat number
func FindRepeatNumber(nums []int) int {
	ret := -1
	rep := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := rep[v]; ok {
			ret = v
			break
		}
		rep[v] = struct{}{}
	}
	return ret
}

// nums is a Non-decrement array, e.g. [2,3,3,5,7,8,8,8,10]
func SearchRepeatCount(nums []int, target int) int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return 0
	}
	// rightmost := sort.SearchInts(nums, target+1) - 1
	return sort.SearchInts(nums, target+1) - leftmost
}

// nums is a Non-decrement array, e.g. [2,3,3,5,7,8,8,8,10]
func SearchRepeatCountNormal(nums []int, target int) int {
	count := 0
	for _, v := range nums {
		if v == target {
			count++
		}
		if v > target {
			break
		}
	}
	return count
}

// nums is an natural number array with just one element less than nums[n-1] missing, e.g. [0,1,2,3,4,5,6,7,9] missing 8
func MissingNumber(nums []int) int {
	l, r := 0, len(nums)-1
	m := 0
	for l <= r {
		m = (r + l) >> 1
		if nums[m] == m {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}

func FindNumberIn2DArray(matrix [][]int, target int) bool {
	idxRow, idxCol := len(matrix)-1, 0
	for idxRow >= 0 && idxCol < len(matrix[0]) {
		if target == matrix[idxRow][idxCol] {
			return true
		}
		if target < matrix[idxRow][idxCol] {
			idxRow--
		} else {
			idxCol++
		}
	}
	return false
}

func AddToArrayForm(num []int, k int) []int {
	ret := make([]int, len(num)+1)
	ret[0] = k
	for i := len(num) - 1; i >= 0; i-- {
		c := num[i] + ret[len(num)-1-i]
		ret[len(num)-1-i] = c % 10
		ret[len(num)-i] = c / 10
	}

	lastRet := ret[len(ret)-1]
	if lastRet == 0 {
		ret = ret[:len(ret)-1]
	}
	for lastRet >= 10 {
		ret[len(ret)-1] = lastRet % 10
		ret = append(ret, lastRet/10)
		lastRet = lastRet / 10
	}
	i, j := 0, len(ret)-1
	for i < j {
		ret[i], ret[j] = ret[j], ret[i]
		i++
		j--
	}
	return ret
}
