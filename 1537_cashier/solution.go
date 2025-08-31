package cashier

type Cashier struct {
	i        int
	n        int
	discount int
	priceMap map[int]int
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	priceMap := make(map[int]int, len(products))
	for i, product := range products {
		priceMap[product] = prices[i]
	}
	return Cashier{n: n, discount: discount, priceMap: priceMap}
}

func (this *Cashier) GetBill(product []int, amount []int) float64 {
	this.i++
	var total = 0.0
	for i, v := range product {
		total += float64(this.priceMap[v] * amount[i])
	}
	if this.i%this.n == 0 {
		total = total - float64(this.discount)*total/100.0
	}
	return total
}
