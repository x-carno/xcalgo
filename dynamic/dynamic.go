package dynamic

// find max margin(>=0) between idx i and j for j > i
// input: [7,1,5,3,6,4]
// output: 5 (i=1,j=4)
func MaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	maxProfix := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		m := prices[i] - minPrice
		if m > maxProfix {
			maxProfix = m
		}
	}
	return maxProfix
}

// given 2-d grid, find the path of the max sum from up-left corner to bottom-down corner
// for a coordinate (i,j), next step should just go bottom or righ: (i+1,j) or (i,j+1)
// e.g.
//grid: [][]int{
// 	{1, 3, 1},
// 	{1, 5, 1},
// 	{4, 2, 1},
// },
//}
// path is : 1-->3-->5-->2-->1, sum is 12
func MaxValue(grid [][]int) int {
	if len(grid) == 1 && len(grid[0]) == 1 {
		return grid[0][0]
	}
	row, col := len(grid), len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 && j > 0 {
				grid[i][j] += grid[i][j-1]
			} else if i > 0 && j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				if grid[i-1][j] > grid[i][j-1] {
					grid[i][j] += grid[i-1][j]
				} else {
					grid[i][j] += grid[i][j-1]
				}
			}
		}
	}
	return grid[row-1][col-1]
}

// translate a number to a string
// for 0->a, 1->b, 2->c ...... 25->z
// e.g. 12258 can be translate to "bccfi", "bwfi", "bczi", "mcfi" or "mzi", so return 5
func TranslateNum(num int) int {
	// kind := 1
	// rem1, rem2 := 0, 0
	// for num > 0 {
	// 	rem1 = num % 100
	// 	if rem1 > 9 && rem1 < 26 {

	// 	}
	// }

	// return kind

	if num <= 9 {
		return 1
	}
	rem := num % 100
	if rem <= 9 || rem >= 26 {
		return TranslateNum(num / 10)
	} else {
		return TranslateNum(num/10) + TranslateNum(num/100)
	}

}

// find the longest sub string that has no repeat character,return the length of sub string
func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	maxSubLength := 1
	maxEndWithI := 1 // indicate the longest substring's length that ends with s[i]
	charAtIdx := make(map[byte]int)
	charAtIdx[s[0]] = 0
	for i := 1; i < len(s); i++ {
		if preIdx, ok := charAtIdx[s[i]]; !ok {
			maxEndWithI++
		} else {
			if i-preIdx > maxEndWithI {
				maxEndWithI++
			} else {
				maxEndWithI = i - preIdx
			}
		}
		charAtIdx[s[i]] = i
		if maxEndWithI > maxSubLength {
			maxSubLength = maxEndWithI
		}
	}

	return maxSubLength
}
