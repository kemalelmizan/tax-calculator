package model

import (
	"database/sql"
	"log"
)

// Product ...
type Product struct {
	Name    string `db:"name"`
	TaxCode int    `db:"tax_code"`
	Price   int64  `db:"price"`
}

type productModel struct {
	db     *sql.DB
	logger *log.Logger
}

// ProductModel is interface of productModel.
//go:generate mockery -name=ProductModel
type ProductModel interface {
	InsertProducts([]Product) error
}

// InsertProducts ...
func (pm productModel) InsertProducts(products []Product) (err error) {
	var id int64

	for _, product := range products {
		err = pm.db.QueryRow(`
INSERT INTO products (name, tax_code, price)
VALUES ($1, $2, $3)
RETURNING id`,
			product.Name,
			product.TaxCode,
			product.Price,
		).Scan(&id)
		if err != nil {
			pm.logger.Println(err)
			continue
		}
		pm.logger.Println("New product inserted to DB:", id)
	}

	return nil
}

// NewProductModel is function creates a new instance of ProductModel
func NewProductModel(db *sql.DB, logger *log.Logger) ProductModel {
	return productModel{
		db:     db,
		logger: logger,
	}
}
