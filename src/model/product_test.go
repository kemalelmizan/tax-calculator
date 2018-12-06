package model

import (
	"database/sql/driver"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func Test_productModel_SelectProductsFromNames(t *testing.T) {
	type args struct {
		productNames []string
	}
	tests := []struct {
		name         string
		args         args
		wantProducts []Product
		wantErr      bool
	}{
		{
			name: "Happy path",
			args: args{
				productNames: []string{"Lucky Stretch"},
			},
			wantErr: false,
			wantProducts: []Product{
				Product{
					Name:    "Lucky Stretch",
					TaxCode: 2,
					Price:   110000,
				},
			},
		},
		{
			name: "Two products",
			args: args{
				productNames: []string{"Big Mac", "Lucky Stretch"},
			},
			wantErr: false,
			wantProducts: []Product{
				Product{
					Name:    "Big Mac",
					TaxCode: 1,
					Price:   102000,
				},
				Product{
					Name:    "Lucky Stretch",
					TaxCode: 2,
					Price:   110000,
				},
			},
		},
		{
			name: "0 length product names",
			args: args{
				productNames: []string{},
			},
			wantErr:      true,
			wantProducts: []Product{},
		},
		{
			name: "0 length product names",
			args: args{
				productNames: []string{},
			},
			wantErr:      true,
			wantProducts: []Product{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			productNamesInterface := make([]driver.Value, len(tt.args.productNames))
			for i, productName := range tt.args.productNames {
				productNamesInterface[i] = strings.ToLower(strings.TrimSpace(productName))
			}

			if len(tt.args.productNames) == 1 {
				mock.ExpectQuery("SELECT (.+) FROM products (.+)").
					WithArgs("lucky stretch").
					WillReturnRows(sqlmock.NewRows([]string{"name", "tax_code", "price"}).
						AddRow("Lucky Stretch", 2, 110000))
			} else if len(tt.args.productNames) == 2 {
				mock.ExpectQuery("SELECT (.+) FROM products (.+) OR (.+)").
					WithArgs("big mac", "lucky stretch").
					WillReturnRows(sqlmock.NewRows([]string{"name", "tax_code", "price"}).
						AddRow("Big Mac", 1, 102000).
						AddRow("Lucky Stretch", 2, 110000))
			}

			pm := productModel{
				db:     db,
				logger: log.New(os.Stdout, "test: ", log.LstdFlags),
			}
			gotProducts, err := pm.SelectProductsFromNames(tt.args.productNames)
			if (err != nil) != tt.wantErr {
				t.Errorf("productModel.SelectProductsFromNames() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotProducts, tt.wantProducts) {
				t.Errorf("productModel.SelectProductsFromNames() = %v, want %v", gotProducts, tt.wantProducts)
			}
		})
	}
}

func Test_productModel_InsertProducts(t *testing.T) {
	type args struct {
		products []Product
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy Path",
			args: args{
				products: []Product{
					Product{
						Name:    "a",
						TaxCode: 1,
						Price:   100,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			productInterface := make([]driver.Value, len(tt.args.products))
			for i, product := range tt.args.products {
				productInterface[i] = product
			}

			mock.ExpectQuery("INSERT (.+) INTO products (.+)").
				WithArgs("a", 1, 100).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).
					AddRow(1))

			pm := productModel{
				db:     db,
				logger: log.New(os.Stdout, "test: ", log.LstdFlags),
			}

			if err := pm.InsertProducts(tt.args.products); (err != nil) != tt.wantErr {
				t.Errorf("productModel.InsertProducts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewProductModel(t *testing.T) {
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	got := NewProductModel(db, log.New(os.Stdout, "test: ", log.LstdFlags))
	require.NotNil(t, got)
}
