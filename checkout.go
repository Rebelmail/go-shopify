package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const checkoutsBasePath = "admin/checkouts"

// CheckoutService is an interface for interfacing with the checkouts endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/abandoned_checkouts
type CheckoutService interface {
	List(interface{}) ([]Checkout, error)
	Count(interface{}) (int64, error)
	Get(int64, interface{}) (*Checkout, error)
}

// CheckoutServiceOp handles communication with the checkout related methods of
// the Shopify API.
type CheckoutServiceOp struct {
	client *Client
}

// A struct for all available checkout list options.
// See: https://help.shopify.com/api/reference/abandoned_checkouts#index
type CheckoutListOptions struct {
	Page         int64     `url:"page,omitempty"`
	Limit        int64     `url:"limit,omitempty"`
	SinceID      int64     `url:"since_id,omitempty"`
	Status       string    `url:"status,omitempty"`
	CreatedAtMin time.Time `url:"created_at_min,omitempty"`
	CreatedAtMax time.Time `url:"created_at_max,omitempty"`
	UpdatedAtMin time.Time `url:"updated_at_min,omitempty"`
	UpdatedAtMax time.Time `url:"updated_at_max,omitempty"`
}

// Checkout represents a Shopify abandoned checkout
type Checkout struct {
	ID                    int64            `json:"id"`
	Name                  string           `json:"name"`
	Email                 string           `json:"email"`
	CreatedAt             *time.Time       `json:"created_at"`
	CompletedAt           *time.Time       `json:"completed_at"`
	UpdatedAt             *time.Time       `json:"updated_at"`
	ClosedAt              *time.Time       `json:"closed_at"`
	Customer              *Customer        `json:"customer"`
	BillingAddress        *Address         `json:"billing_address"`
	ShippingAddress       *Address         `json:"shipping_address"`
	Currency              string           `json:"currency"`
	TotalPrice            *decimal.Decimal `json:"total_price"`
	SubtotalPrice         *decimal.Decimal `json:"subtotal_price"`
	TotalDiscounts        *decimal.Decimal `json:"total_discounts"`
	TotalLineItemsPrice   *decimal.Decimal `json:"total_line_items_price"`
	TotalTax              *decimal.Decimal `json:"total_tax"`
	TotalWeight           int64            `json:"total_weight"`
	Token                 string           `json:"token"`
	CartToken             string           `json:"cart_token"`
	Note                  string           `json:"note"`
	BuyerAcceptsMarketing bool             `json:"buyer_accepts_marketing"`
	CancelReason          string           `json:"cancel_reason"`
	NoteAttributes        []NoteAttribute  `json:"note_attributes"`
	DiscountCodes         []DiscountCode   `json:"discount_codes"`
	LineItems             []LineItem       `json:"line_items"`
	AbandonedCheckoutURL  string           `json:"abandoned_checkout_url"`
	Gateway               string           `json:"gateway"`
	LandingSite           string           `json:"landing_site"`
	ReferringSite         string           `json:"referring_site"`
	ShippingLines         []ShippingLine   `json:"shipping_lines"`
	SourceName            string           `json:"source_name"`
	TaxLines              []TaxLine        `json:"tax_lines"`
	TaxesIncluded         bool             `json:"taxes_included"`
}

type ShippingLine struct {
	Code   string           `json:"code"`
	Price  *decimal.Decimal `json:"price"`
	Source string           `json:"source"`
	Title  string           `json:"title"`
}

type TaxLine struct {
	Price *decimal.Decimal `json:"price"`
	Rate  float64          `json:"rate"`
	Title string           `json:"title"`
}

// Represents the result from the checkouts/X.json endpoint
type CheckoutResource struct {
	Checkout *Checkout `json:"checkout"`
}

// Represents the result from the checkouts.json endpoint
type CheckoutsResource struct {
	Checkouts []Checkout `json:"checkouts"`
}

// List checkouts
func (s *CheckoutServiceOp) List(options interface{}) ([]Checkout, error) {
	path := fmt.Sprintf("%s.json", checkoutsBasePath)
	resource := new(CheckoutsResource)
	err := s.client.Get(path, resource, options)
	return resource.Checkouts, err
}

// Count checkouts
func (s *CheckoutServiceOp) Count(options interface{}) (int64, error) {
	path := fmt.Sprintf("%s/count.json", checkoutsBasePath)
	return s.client.Count(path, options)
}

// Get individual checkout
func (s *CheckoutServiceOp) Get(checkoutID int64, options interface{}) (*Checkout, error) {
	path := fmt.Sprintf("%s/%d.json", checkoutsBasePath, checkoutID)
	resource := new(CheckoutResource)
	err := s.client.Get(path, resource, options)
	return resource.Checkout, err
}
