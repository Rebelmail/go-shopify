package goshopify

import (
	"github.com/shopspring/decimal"
)

type LineItem struct {
	ID                 int              `json:"id"`
	ProductID          int              `json:"product_id"`
	VariantID          int              `json:"variant_id"`
	Quantity           int              `json:"quantity"`
	Price              *decimal.Decimal `json:"price"`
	TotalDiscount      *decimal.Decimal `json:"total_discount"`
	Title              string           `json:"title"`
	VariantTitle       string           `json:"variant_title"`
	Name               string           `json:"name"`
	SKU                string           `json:"sku"`
	Vendor             string           `json:"vendor"`
	GiftCard           bool             `json:"gift_card"`
	Taxable            bool             `json:"taxable"`
	FulfillmentService string           `json:"fulfillment_service"`
	FulfillmentStatus  string           `json:"fulfillment_status"`
	Grams              int              `json:"grams"`
	RequiresShipping   bool             `json:"requires_shipping"`
}
