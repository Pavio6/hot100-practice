package hot100practice

import "math"

// 121. 买卖股票的最佳时机

func MaxProfit(prices []int) int {
	// 遍历prices
	// 1.比较当前currentPrice和minPrice，取最小赋值给minPrice
	// 2.比较maxProfit和currentPrice-minPrice，取最大值给maxPrice
	// 这样就可以保证maxPrice在每次遍历中都是最大差值
	minPrice := math.MaxInt32
	maxProfit := 0

	for _, currentPrice := range prices {
		if currentPrice < minPrice {
			minPrice = currentPrice
		}
		tmp := currentPrice - minPrice
		if tmp > maxProfit {
			maxProfit = tmp
		}
	}
	return maxProfit
}
