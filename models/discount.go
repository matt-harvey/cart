package models

import "errors"

type Discount interface {

	// Apply accepts a given price and returns a new, discounted price.
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
		return nil, errors.New("Discount factory not found for discount type")
	}
	return discountFactory(discountID)
}
