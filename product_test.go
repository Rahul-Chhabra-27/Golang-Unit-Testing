package main

import (
	"testing"

	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
)

func TestComputeProductCategoryAmount(t *testing.T) {
	 
	t.Run("Nomical Case / Actual Data Checking", func(t *testing.T) {
		productCategory := Products {
			ID:"100",
			Currency: "EUR",
			Products: [] Product {
				{ ID: "1001",name: "Boxing Gloves" , Quanyity: 2, UnitPrice: money.New(100,"EUR") },
				{ ID: "1002",name: "Cricket kit" , Quanyity: 3, UnitPrice: money.New(200,"EUR") },
			},
		}
		amount,err := productCategory.ComputeProductCategoryAmount();

		// checing for error
		assert.NoError(t,err); // return and marked this case as failed case
		assert.Equal(t,int64(800), amount.Amount());	
		assert.Equal(t,"EUR", amount.Currency().Code)
	})
	t.Run("Currency Issues",func(t *testing.T) {
		productCategory := Products {
			ID:"100",
			Currency: "INR",
			Products: [] Product {
				{ ID: "1001",name: "Boxing Gloves" , Quanyity: 2, UnitPrice: money.New(100,"EUR") },
				{ ID: "1002",name: "Cricket kit" , Quanyity: 3, UnitPrice: money.New(200,"EUR") },
			},
		}
		_,err := productCategory.ComputeProductCategoryAmount();

		// checing for error
		assert.Error(t,err); // return and marked this case as failed case
	})
	

	// 	The assert package provides some helpful methods that allow you to write better test code in Go.
	// Prints friendly, easy to read failure descriptions
	// Allows for very readable code
	// Optionally annotate each assertion with a message
}