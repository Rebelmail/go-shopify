package goshopify

import (
	"github.com/shopspring/decimal"
)

// Address represents one of the many addresses a customer may have.
// See: https://help.shopify.com/api/reference/order#shipping-address-property
type Address struct {
	ID           int     `json:"id"`
	Address1     string  `json:"address1"`
	Address2     string  `json:"address2"`
	City         string  `json:"city"`
	Company      string  `json:"company"`
	Country      string  `json:"country"`
	CountryCode  string  `json:"country_code"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Name         string  `json:"name"`
	Phone        string  `json:"phone"`
	Province     string  `json:"province"`
	ProvinceCode string  `json:"province_code"`
	Zip          string  `json:"zip"`
	Default      bool    `json:"default"`
}

// NoteAttribute is extra information that is added to an order.
// See: https://help.shopify.com/api/reference/order#note-attributes-property
type NoteAttribute struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

// LineItem is information about an item in the order.
// See: https://help.shopify.com/api/reference/order#line-items-property
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

// DiscountCode is applicable discount codes that can be applied to an order.
// See: https://help.shopify.com/api/reference/order#discount-codes-property
type DiscountCode struct {
	Amount *decimal.Decimal `json:"amount"`
	Code   string           `json:"code"`
	Type   string           `json:"type"`
}
