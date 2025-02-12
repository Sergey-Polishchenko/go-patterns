package prototype

import (
	"reflect"
	"testing"
)

func TestNewConcreteProduct(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "Create product with name",
			args: "TestProduct",
			want: "TestProduct",
		},
		{
			name: "Create product with empty name",
			args: "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewConcreteProduct(tt.args)
			if !reflect.DeepEqual(got.GetName(), tt.want) {
				t.Errorf("NewConcreteProduct() = %v, want %v", got.GetName(), tt.want)
			}
		})
	}
}

func TestConcreteProduct_GetName(t *testing.T) {
	tests := []struct {
		name   string
		fields string
		want   string
	}{
		{
			name:   "Get existing name",
			fields: "TestProduct",
			want:   "TestProduct",
		},
		{
			name:   "Get empty name",
			fields: "",
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ConcreteProduct{
				name: tt.fields,
			}
			if got := p.GetName(); got != tt.want {
				t.Errorf("ConcreteProduct.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcreteProduct_Clone(t *testing.T) {
	tests := []struct {
		name    string
		product *ConcreteProduct
		want    string
	}{
		{
			name:    "Clone with data",
			product: &ConcreteProduct{name: "Prototype"},
			want:    "Prototype",
		},
		{
			name:    "Clone empty",
			product: &ConcreteProduct{name: ""},
			want:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.product.Clone()

			// Проверка равенства значений
			if clone.GetName() != tt.want {
				t.Errorf("Clone name = %v, want %v", clone.GetName(), tt.want)
			}

			// Проверка, что это разные объекты
			if clone == tt.product {
				t.Error("Clone should be a different object")
			}

			// Типовая проверка через reflect (опционально)
			if !reflect.DeepEqual(clone.GetName(), tt.product.GetName()) {
				t.Errorf("Clone data not equal: %v vs %v", clone, tt.product)
			}
		})
	}
}
