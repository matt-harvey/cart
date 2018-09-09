package models

import "testing"

func TestPinnedDiscountApply(t *testing.T) {

	pinnedDiscount := PinnedDiscount{AmountCents: 7000}

	// When original price is greater than pinned price.
	originalPriceCents := uint(10000)
	discountedPriceCents := pinnedDiscount.Apply(originalPriceCents)
	expectedDiscountedPriceCents := uint(7000)
	if discountedPriceCents != expectedDiscountedPriceCents {
		t.Fatalf(
			"Result %d cents does not match expected result of %d",
			discountedPriceCents,
			expectedDiscountedPriceCents)
	}

	// When original price is less than pinned price.
	originalPriceCents = uint(6000)
	discountedPriceCents = pinnedDiscount.Apply(originalPriceCents)
	expectedDiscountedPriceCents = uint(6000)
	if discountedPriceCents != expectedDiscountedPriceCents {
		t.Fatalf(
			"Result %d cents does not match expected result of %d",
			discountedPriceCents,
			expectedDiscountedPriceCents)
	}
}
