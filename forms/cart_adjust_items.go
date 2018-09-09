package forms

// CartAdjustItems respresents a request to adjust the quantity of a given Product as it
// appears in a Cart.
type CartAdjustItems struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
