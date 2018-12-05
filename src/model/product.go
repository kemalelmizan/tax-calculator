package model

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
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
	SelectProductsFromNames([]string) ([]Product, error)
}

// SelectProductsFromNames ...
func (pm productModel) SelectProductsFromNames(productNames []string) (products []Product, err error) {

	if len(productNames) <= 0 {
		return []Product{}, errors.New("invalid productNames input")
	}

	qs := "SELECT name, tax_code, price FROM products WHERE lower(name)=$1"

	// convert []string to []interface{}
	productNamesInterface := make([]interface{}, len(productNames))
	for i, productName := range productNames {
		productNamesInterface[i] = strings.ToLower(strings.TrimSpace(productName))
		if i > 0 {
			qs += fmt.Sprintf(" OR lower(name)=$%d", i+1)
		}
	}

	pm.logger.Println("Incoming query: ", qs, productNamesInterface)

	rows, err := pm.db.Query(qs, productNamesInterface...)
	if err != nil {
		return []Product{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var taxCode int
		var price int64
		if err := rows.Scan(&name, &taxCode, &price); err != nil {
			return products, err
		}
		products = append(products, Product{
			Name:    name,
			TaxCode: taxCode,
			Price:   price,
		})
	}

	return products, nil
}

// InsertProducts ...
func (pm productModel) InsertProducts(products []Product) (err error) {
	var id int64

	for _, product := range products {
		err = pm.db.QueryRow("INSERT INTO products (name, tax_code, price) VALUES ($1, $2, $3) RETURNING id",
			product.Name, product.TaxCode, product.Price).Scan(&id)
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
