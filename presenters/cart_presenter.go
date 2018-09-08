package presenters

import (
	"github.com/matt-harvey/cart/db"
	"github.com/matt-harvey/cart/models"
)

type CartPresenter struct {
	ID              int                `json:"id"`
	TotalPriceCents uint               `json:"total_price_cents"`
	Items           CartItemPresenters `json:"items"`
}

func NewCartPresenter(cartID int) (*CartPresenter, error) {
	cartPresenter := CartPresenter{ID: cartID}
	cart := models.Cart{}
	err := db.Conn().Eager("Items.Product").Find(&cart, cartID)
	if err != nil {
		return &CartPresenter{}, err
	}
	for _, cartItem := range cart.Items {
		cartItemPresenter := CartItemPresenter{
			ProductName:          cartItem.Product.Name,
			Quantity:             cartItem.Quantity,
			StandardPriceCents:   cartItem.StandardPriceCents,
			DiscountCents:        cartItem.StandardPriceCents - cartItem.DiscountedPriceCents,
			DiscountedPriceCents: cartItem.DiscountedPriceCents,
		}
		cartPresenter.Items = append(cartPresenter.Items, cartItemPresenter)
		cartPresenter.TotalPriceCents += cartItem.DiscountedPriceCents
	}
	return &cartPresenter, nil
}
