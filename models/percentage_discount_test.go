package models

import "testing"

func TestPercentageDiscountApply(t *testing.T) {
	percentageDiscount := PercentageDiscount{Amount: 0.75}
	originalPriceCents := uint(303)
	discountedPriceCents := percentageDiscount.Apply(originalPriceCents)
	expectedDiscountedPriceCents := uint(76)
	if discountedPriceCents != expectedDiscountedPriceCents {
		t.Fatalf(
			"Result %d cents does not match expected result of %d",
			discountedPriceCents,
			expectedDiscountedPriceCents)
	}
}
