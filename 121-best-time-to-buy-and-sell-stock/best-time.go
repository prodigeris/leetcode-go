package _21_best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	l, r := 0, 0
	max := 0
	for r < len(prices)-1 {
		r++
		p1, p2 := prices[l], prices[r]
		if p2 < p1 {
			l = r
			continue
		}
		if p2-p1 > max {
			max = p2 - p1
		}
	}
	return max
}
