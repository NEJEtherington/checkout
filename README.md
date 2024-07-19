# Supermarket Checkout

Implementation of a supermarket checkout that calculates the total price of a number of items, with individual and discounted prices as follows:

| SKU | Unit Price | Special Price |
|-----|------------|---------------|
| A   | 50         | 3 for 130     |
| B   | 30         | 2 for 45      |
| C   | 20         |               |
| D   | 15         |               |

The supermarket's inventory comprises only the above four items at present and entering any SKU other than A, B, C, or D will result in a error.

## Running the app

`go run internal/main.go`

You will then be prompted to scan an item, which is done by tpying an SKU (i.e. A, B, C, or D) and hitting return.  The total cost of your basket, net of any discounts,  will be displayed as each item is added.

## Testing

`go test ./cmd/checkout`