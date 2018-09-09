package presenters

type CartItemPresenter struct {
	ProductID            int    `json:"product_id"`
	ProductName          string `json:"product_name"`
	Quantity             uint   `json:"quantity"`
	StandardPriceCents   uint   `json:"standard_price_cents"`
	DiscountCents        uint   `json:"discount_cents"`
	DiscountedPriceCents uint   `json:"discounted_price_cents"`
}

type CartItemPresenters []CartItemPresenter
