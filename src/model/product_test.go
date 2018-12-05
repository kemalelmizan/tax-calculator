package model

import (
	"database/sql/driver"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
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

			mock.ExpectQuery("SELECT (.+) FROM products (.+)").
				WithArgs("lucky stretch").
				WillReturnRows(sqlmock.NewRows([]string{"name", "tax_code", "price"}).
					AddRow("Lucky Stretch", 2, 110000))

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
