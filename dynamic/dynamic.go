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
