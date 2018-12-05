package controller

import (
	"strconv"

	"github.com/kemalelmizan/tax-calculator/src/model"
)

// BillOutput ...
type BillOutput struct {
	Products []Product `json:"products"`
	Total    string    `json:"total"`
}

type billController struct {
	model model.ProductModel
}

// BillController is interface of billModel.
//go:generate mockery -name=BillController
type BillController interface {
	GetBill([]string) (BillOutput, error)
}

// GetBill ...
func (bc billController) GetBill(productNames []string) (BillOutput, error) {
	var bill BillOutput
	var total float64

	var inputs []ProductInput
	p, err := bc.model.SelectProductsFromNames(productNames)
	if err != nil {
		return bill, err
	}

	for _, product := range p {
		inputs = append(inputs, ProductInput{
			Name:    product.Name,
			TaxCode: product.TaxCode,
			Price:   float64(product.Price) / 100,
		})
	}

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

	return bill, nil
}

// NewBillController is function creates a new instance of BillModel
func NewBillController(model model.ProductModel) BillController {
	return billController{
		model: model,
	}
}
