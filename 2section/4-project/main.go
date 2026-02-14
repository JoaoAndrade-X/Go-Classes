package main

import (
	"fmt"
	"strings"
)

var productPrizes = map[string]float64{
	"TSHIRT": 20.00,
	"MUG":    12.50,
	"HAT":    18.00,
	"BOOK":   25.99,
}

func calculateItemPrice(itemCode string) (float64, bool) {
	basePrice, found := productPrizes[itemCode]
	if !found {
		if strings.HasSuffix(itemCode, "_SALE") {
			originalItemCode := strings.TrimSuffix(itemCode, "_SALE")
			basePrice, found = productPrizes[originalItemCode]
			if found {
				salePrice := basePrice * 0.9
				fmt.Printf(
					" - Item %s (Sale! Original: $%.2f, Sale Price: $%.2f)\n",
					originalItemCode,
					basePrice,
					salePrice,
				)
				return salePrice, true
			}
		}

		fmt.Printf(" - Item: %s (Product not found)\n", itemCode)
		return 0.0, false
	}
	return basePrice, true
}

func main() {
	fmt.Println("-------- Simple Sales Order Processor --------")

	orderItems := []string{
		"TSHIRT", "MUG_SALE", "HAT", "BOOK",
	}

	fmt.Println("-------- Processing Order Items --------")
	var subtotal float64
	for _, itemCode := range orderItems {
		price, found := calculateItemPrice(itemCode)
		if found {
			subtotal += price
		}
	}

	fmt.Println("Total Price: ", subtotal)
}
