package checkout

import (
	"testing"
)

func TestScan(t *testing.T) {
	checkout := NewCheckout(Inventory)

	tests := []struct {
		name    string
		SKU     string
		wantErr bool
	}{
		{
			name:    "Product does not exist",
			SKU:     "Z",
			wantErr: true,
		},
		{
			name:    "Product exists",
			SKU:     "A",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := checkout.Scan(tt.SKU)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scan error = %v, expected %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				_, ok := checkout.Basket[tt.SKU]
				if !ok {
					t.Errorf("Failed to add %s to  basket", tt.SKU)
				}
			}
		})
	}
}

func TestGetTotalPrice(t *testing.T) {

	tests := []struct {
		name      string
		scanItems []string // items to be scanned to populate shopping basket
		want      int
		wantErr   bool
	}{
		{
			name:      "Single item",
			scanItems: []string{"A"},
			want:      50,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		checkout := NewCheckout(Inventory)

		t.Run(tt.name, func(t *testing.T) {

			for _, item := range tt.scanItems {
				err := checkout.Scan(item)
				if (err != nil) != tt.wantErr {
					t.Errorf("Scan error = %v, expected %v", err, tt.wantErr)
				}
			}

			got, err := checkout.GetTotalPrice()
			if (err != nil) && !tt.wantErr {
				t.Errorf("GetTotalPrice error = %v, expected %v", err, tt.wantErr)

				return
			}
			if got != tt.want {
				t.Errorf("GetTotalPrice returned %v, expected %v", got, tt.want)
			}
		})
	}
}
