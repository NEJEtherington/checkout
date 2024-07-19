package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/NEJEtherington/checkout/cmd/checkout"
)

func StartApp() {
	// initialise new Checkout
	checkout := checkout.NewCheckout(checkout.Inventory)
	// prompt the user for input
	fmt.Println("Please scan an item or type q and hit enter to quit")
	// call scan function with input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "q" {

			return
		}

		err := checkout.Scan(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
		// calculate and log total price
		price, err := checkout.GetTotalPrice()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Basket total is %d\n", price)
		fmt.Println("Please scan the next item or type q and hit enter to quit")
	}
}
