package main

import (
	"fmt"

	"github.com/Rhymond/go-money"
)

// Individual Product
type Product struct {
	ID string
	name string
	Quanyity uint
	// This Contains two fields 1. currency and 2. amount  -> $ 2
	UnitPrice *money.Money // WHAT IS THE CURRENCY AND WHAT IS THE ACTUAL PRICE 
}
// unitPrice.Current -> dollars
// unitPrice.amount --> 2
// products
// product categort
// category --> sports , cricket kit , boxing gloves, hockey.................. kit bag -> quantity : 2 id : 10, 100 EUR
type Products struct {
	ID string
	Products []  Product // wrong Product; --> product 1 : using currency USD , Second product using currency INR
	Currency string
}
// product { name: kit , price: EURO 100, quanity:2}
// product { name: gloves , price: EURO 200 }
// product { name: dumbells , price: EURO 400}
// create a function that returns the amount of product catregory
// sould return 700 Euro --> 80O EURO

func (productCategory Products) ComputeProductCategoryAmount() (*money.Money, error) {
	totalAmount := money.New(0, productCategory.Currency);

	for _,product := range productCategory.Products {
		// we are adding the price of each product to the totalAmountVariable....
		var err error;
		totalAmount ,err = totalAmount.Add(product.UnitPrice.Multiply(int64(product.Quanyity)));

		if err != nil {
			return nil , fmt.Errorf("Impossible to add Item");
		}
	}
	return totalAmount,nil;
}