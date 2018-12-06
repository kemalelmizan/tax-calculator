package controller

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/kemalelmizan/tax-calculator/src/model/mocks"
	"github.com/stretchr/testify/mock"
)

func Test_productController_PostProduct(t *testing.T) {
	type args struct {
		inputs []ProductInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy path",
			args: args{
				inputs: []ProductInput{
					ProductInput{
						Name: "a",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProductModel := new(mocks.ProductModel)
			mockProductModel.On("InsertProducts", mock.Anything).Return(nil)

			pc := productController{
				model: mockProductModel,
			}
			if err := pc.PostProduct(tt.args.inputs); (err != nil) != tt.wantErr {
				t.Errorf("productController.PostProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewProductController(t *testing.T) {
	mockProductModel := new(mocks.ProductModel)
	got := NewProductController(mockProductModel)
	require.NotNil(t, got)
}
