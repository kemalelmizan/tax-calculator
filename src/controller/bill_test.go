package controller

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/kemalelmizan/tax-calculator/src/model"
	"github.com/kemalelmizan/tax-calculator/src/model/mocks"
)

func Test_billController_GetBill(t *testing.T) {
	type args struct {
		productNames []string
	}
	tests := []struct {
		name    string
		args    args
		want    BillOutput
		wantErr bool
	}{
		{
			name: "Happy path",
			args: args{
				productNames: []string{"a"},
			},
			want: BillOutput{
				Products: []Product{
					Product{
						Name:   "a",
						Price:  "0.00",
						Tax:    "0.00",
						Amount: "0.00",
					},
				},
				Total: "0.00",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProductModel := new(mocks.ProductModel)
			mockProductModel.On("SelectProductsFromNames", []string{"a"}).
				Return([]model.Product{
					model.Product{
						Name: "a",
					},
				}, nil)

			bc := billController{
				model: mockProductModel,
			}
			got, err := bc.GetBill(tt.args.productNames)
			if (err != nil) != tt.wantErr {
				t.Errorf("billController.GetBill() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("billController.GetBill() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBillController(t *testing.T) {
	mockProductModel := new(mocks.ProductModel)
	got := NewBillController(mockProductModel)
	require.NotNil(t, got)
}
