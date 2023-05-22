package main

import (
	"fmt"
	o "github.com/adetunjiakintundeakinde/terraform-provider-orders/orders"
)

func main() {

	orders := o.ImportState()

	order := o.Order{
		Id:        "3904712d-baa9-4d6f-b0ab-a2209af5a830",
		Recipient: "Colorado New Jersey",
		Date:      "2009-06-12",
		Cost:      "12",
	}

	orders, err := o.ResourceAdd(order, *orders)
	if err != nil {
		panic(err)
	}

	fmt.Println(orders)
}
