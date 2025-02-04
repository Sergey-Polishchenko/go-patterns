package factorymethod

import (
	"reflect"
	"testing"
)

func TestNewProduct(t *testing.T) {
	tests := []struct {
		name     string
		actor    actor
		wantType Product
		wantErr  bool
	}{
		{
			name:     "Actor A should create ProductA",
			actor:    A,
			wantType: &ProductA{},
			wantErr:  false,
		},
		{
			name:     "Actor B should create ProductB",
			actor:    B,
			wantType: &ProductB{},
			wantErr:  false,
		},
		{
			name:     "Actor C should create ProductC",
			actor:    C,
			wantType: &ProductC{},
			wantErr:  false,
		},
		{
			name:     "Unknown actor should return error",
			actor:    actor(42),
			wantType: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				got, err := NewProduct(tt.actor)
				if (err != nil) != tt.wantErr {
					t.Errorf("NewProduct() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if tt.wantErr {
					return
				}

				if reflect.TypeOf(got) != reflect.TypeOf(tt.wantType) {
					t.Errorf("NewProduct() = %T, want %T", got, tt.wantType)
				}

				got.DoSmt()
			},
		)
	}
}
