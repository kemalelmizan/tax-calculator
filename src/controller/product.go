package controller

import (
	"strconv"

	"github.com/kemalelmizan/tax-calculator/src/model"
)

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

type productController struct {
	model model.ProductModel
}

// ProductController is interface of productModel.
//go:generate mockery -name=ProductModel
type ProductController interface {
	PostProduct([]ProductInput) (BillOutput, error)
}

// PostProduct ...
func (pc productController) PostProduct(inputs []ProductInput) (BillOutput, error) {
	var bill BillOutput
	var total float64
	var productsModel []model.Product

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

		productModel := model.Product{
			Name:    input.Name,
			Price:   int64(input.Price),
			TaxCode: input.TaxCode,
		}

		productsModel = append(productsModel, productModel)
		bill.Products = append(bill.Products, product)
	}
	bill.Total = strconv.FormatFloat(total, 'f', 2, 64)

	err := pc.model.InsertProducts(productsModel)
	if err != nil {
		return BillOutput{}, err
	}

	return bill, nil
}

// NewProductController is function creates a new instance of ProductModel
func NewProductController(model model.ProductModel) ProductController {
	return productController{
		model: model,
	}
}
