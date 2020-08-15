package main

func main() {

}

func maxProfit(prices []int) {
	sum := 0

	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			sum += prices[i] - prices[i-1]
		}
	}
}
