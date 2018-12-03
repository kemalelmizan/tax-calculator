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

type model struct {
	db     *sql.DB
	logger *log.Logger
}

// Model is interface of model.
//go:generate mockery -name=Model
type Model interface {
	InsertProducts([]Product) error
}

// InsertProducts ...
func (m model) InsertProducts(products []Product) (err error) {
	var id int64

	for _, product := range products {
		err = m.db.QueryRow(`
INSERT INTO products (name, tax_code, price)
VALUES ($1, $2, $3)
RETURNING id`,
			product.Name,
			product.TaxCode,
			product.Price,
		).Scan(&id)
		if err != nil {
			return err
		}
		m.logger.Println("inserted to DB: ", id)
	}

	return nil
}

// NewModel is function creates a new instance of Model
func NewModel() Model {
	db, err := initDB()
	if err != nil {
		panic(err)
	}

	return model{
		db:     db,
		logger: nil,
	}
}
