package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const ordersBasePath = "admin/orders"

// OrderService is an interface for interfacing with the orders endpoints of
// the Shopify API.
// See: https://help.shopify.com/api/reference/order
type OrderService interface {
	List(interface{}) ([]Order, error)
	Count(interface{}) (int, error)
	Get(int, interface{}) (*Order, error)
}

// OrderServiceOp handles communication with the order related methods of the
// Shopify API.
type OrderServiceOp struct {
	client *Client
}

// A struct for all available order list options.
// See: https://help.shopify.com/api/reference/order#index
type OrderListOptions struct {
	Page              int       `url:"page,omitempty"`
	Limit             int       `url:"limit,omitempty"`
	SinceID           int       `url:"since_id,omitempty"`
	Status            string    `url:"status,omitempty"`
	FinancialStatus   string    `url:"financial_status,omitempty"`
	FulfillmentStatus string    `url:"fulfillment_status,omitempty"`
	CreatedAtMin      time.Time `url:"created_at_min,omitempty"`
	CreatedAtMax      time.Time `url:"created_at_max,omitempty"`
	UpdatedAtMin      time.Time `url:"updated_at_min,omitempty"`
	UpdatedAtMax      time.Time `url:"updated_at_max,omitempty"`
	ProcessedAtMin    time.Time `url:"processed_at_min,omitempty"`
	ProcessedAtMax    time.Time `url:"processed_at_max,omitempty"`
	Fields            string    `url:"fields,omitempty"`
}

// Order represents a Shopify order
type Order struct {
	ID                    int              `json:"id"`
	Name                  string           `json:"name"`
	Email                 string           `json:"email"`
	CreatedAt             *time.Time       `json:"created_at"`
	UpdatedAt             *time.Time       `json:"updated_at"`
	ClosedAt              *time.Time       `json:"closed_at"`
	ProcessedAt           *time.Time       `json:"processed_at"`
	Customer              *Customer        `json:"customer"`
	BillingAddress        *Address         `json:"billing_address"`
	ShippingAddress       *Address         `json:"shipping_address"`
	Currency              string           `json:"currency"`
	TotalPrice            *decimal.Decimal `json:"total_price"`
	SubtotalPrice         *decimal.Decimal `json:"subtotal_price"`
	TotalDiscounts        *decimal.Decimal `json:"total_discounts"`
	TotalLineItemsPrice   *decimal.Decimal `json:"total_line_items_price"`
	TotalTax              *decimal.Decimal `json:"total_tax"`
	TotalWeight           int              `json:"total_weight"`
	FinancialStatus       string           `json:"financial_status"`
	FulfillmentStatus     string           `json:"fulfillment_status"`
	Token                 string           `json:"token"`
	CartToken             string           `json:"cart_token"`
	Number                int              `json:"number"`
	OrderNumber           int              `json:"order_number"`
	Note                  string           `json:"note"`
	Test                  bool             `json:"test"`
	BrowserIp             string           `json:"browser_ip"`
	BuyerAcceptsMarketing bool             `json:"buyer_accepts_marketing"`
	CancelReason          string           `json:"cancel_reason"`
	NoteAttributes        []NoteAttribute  `json:"note_attributes"`
	DiscountCodes         []DiscountCode   `json:"discount_codes"`
	LineItems             []LineItem       `json:"line_items"`
}

// Represents the result from the orders/X.json endpoint
type OrderResource struct {
	Order *Order `json:"order"`
}

// Represents the result from the orders.json endpoint
type OrdersResource struct {
	Orders []Order `json:"orders"`
}

// List orders
func (s *OrderServiceOp) List(options interface{}) ([]Order, error) {
	path := fmt.Sprintf("%s.json", ordersBasePath)
	resource := new(OrdersResource)
	err := s.client.Get(path, resource, options)
	return resource.Orders, err
}

// Count orders
func (s *OrderServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", ordersBasePath)
	return s.client.Count(path, options)
}

// Get individual order
func (s *OrderServiceOp) Get(orderID int, options interface{}) (*Order, error) {
	path := fmt.Sprintf("%s/%d.json", ordersBasePath, orderID)
	resource := new(OrderResource)
	err := s.client.Get(path, resource, options)
	return resource.Order, err
}
