package controller

import (
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

type productController struct {
	model model.ProductModel
}

// ProductController is interface of productModel.
//go:generate mockery -name=ProductModel
type ProductController interface {
	PostProduct([]ProductInput) error
}

// PostProduct ...
func (pc productController) PostProduct(inputs []ProductInput) error {
	var productsModel []model.Product

	for _, input := range inputs {
		productModel := model.Product{
			Name:    input.Name,
			Price:   int64(input.Price),
			TaxCode: input.TaxCode,
		}
		productsModel = append(productsModel, productModel)
	}

	err := pc.model.InsertProducts(productsModel)
	if err != nil {
		return err
	}
	return nil
}

// NewProductController is function creates a new instance of ProductModel
func NewProductController(model model.ProductModel) ProductController {
	return productController{
		model: model,
	}
}
