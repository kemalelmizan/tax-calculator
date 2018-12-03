package controller

import "strconv"

// ProductInput ...
type ProductInput struct {
	Name    string  `json:"name"`
	TaxCode int     `json:"tax_code"`
	Price   float64 `json:"price"`
}

// Product ...
type Product struct {
	Name       string `json:"name"`
	TaxCode    int    `json:"tax_code"`
	Type       string `json:"type"`
	Refundable string `json:"refundable"`
	Price      string `json:"price"`
	Tax        string `json:"tax"`
	Amount     string `json:"amount"`
}

// BillOutput ...
type BillOutput struct {
	Products []Product `json:"products"`
	Total    string    `json:"total"`
}

// PostProduct ...
func PostProduct(inputs []ProductInput) BillOutput {
	var bill BillOutput
	var total float64

	for _, input := range inputs {
		price := float64(input.Price)
		tax := (map[int]float64{
			1: (10.0 / 100) * price,
			2: (10 + (2.0 / 100 * price)),
			3: (map[bool]float64{
				true:  0,
				false: (1.0 / 100) * (price - 100),
			})[input.Price > 0 && input.Price < 100],
		})[input.TaxCode]
		amount := price + tax
		total += amount

		product := Product{
			Name:    input.Name,
			Price:   strconv.FormatFloat(price, 'f', 2, 64),
			TaxCode: input.TaxCode,
			Type: (map[int]string{
				1: "Food & Beverage",
				2: "Tobacco",
				3: "Entertainment",
			})[input.TaxCode],
			Refundable: (map[int]string{
				1: "yes",
				2: "no",
				3: "no",
			})[input.TaxCode],
			Tax:    strconv.FormatFloat(tax, 'f', 2, 64),
			Amount: strconv.FormatFloat(amount, 'f', 2, 64),
		}

		bill.Products = append(bill.Products, product)
	}
	bill.Total = strconv.FormatFloat(total, 'f', 2, 64)

	// TODO: save to db

	return bill
}
