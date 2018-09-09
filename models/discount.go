package models

import "fmt"

type Discount interface {

	// Apply takes the price for a single unit of a Product, and returns the price that
	// unit would have were this Discount applied.
	Apply(priceCents uint) uint
}

type discountFactory func(id int) (Discount, error)

var discountFactories = make(map[string]discountFactory)

func registerDiscountType(typeName string, factory discountFactory) {
	discountFactories[typeName] = factory
}

func getDiscount(discountType string, discountID int) (Discount, error) {
	discountFactory, ok := discountFactories[discountType]
	if !ok {
		// TODO Better errors
		return nil, fmt.Errorf("Discount factory not found for discount type %s", discountType)
	}
	return discountFactory(discountID)
}
