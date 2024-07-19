package checkout

import (
	"fmt"
)

type ICheckout interface {
	Scan(SKU string) error
	GetTotalPrice() (int, error)
}

type Product struct {
	SKU              string
	UnitPrice        int
	DiscountQuantity int
	DiscountedPrice  int
}

type Checkout struct {
	catalogue map[string]Product
	Basket    map[string]int
}

// NewCheckout create a new Checkout instance and populates the catalogue with passed Product slice
func NewCheckout(products []Product) *Checkout {
	productMap := make(map[string]Product)
	for _, product := range products {
		productMap[product.SKU] = product
	}

	return &Checkout{
		catalogue: productMap,
		Basket:    make(map[string]int),
	}
}

// Scan adds an item to the Checkout Basket
func (c *Checkout) Scan(SKU string) error {

	// check if SKU exists in products slice
	// return an error if not
	_, ok := c.catalogue[SKU]
	if !ok {

		return fmt.Errorf("invalid SKU: %s", SKU)
	}
	// otherwise add the product to the basket

	return nil
}

// GetTotalPrice calculates the total price of the Checkout Basket, net of any discounts
func (c *Checkout) GetTotalPrice() (int, error) {
	totalPrice := 0

	// range over basket
	// check if SKU exists in checkout products slice
	// return an error if not
	// check whether product is subject to discount - use modulo
	// update the checkout totalPrice with the relevant amount
	// return the checkout totalPrice

	return totalPrice, nil
}
