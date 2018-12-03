package model

// Product ...
type Product struct {
	Name    string `db:"name"`
	TaxCode int    `db:"tax_code"`
	Price   int64  `db:"price"`
}
