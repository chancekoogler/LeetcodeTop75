package BuSellStock1

func maxProfit(prices []int) int {
	var profit int
	min := prices[0]
	for i := range prices {
		if prices[i] < min {
			min = prices[i]
		} else {
			tmp := prices[i] - min
			if tmp > profit {
				profit = tmp
			}
		}
	}
	return profit
}
