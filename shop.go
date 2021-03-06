package goshopify

import "time"

// ShopService is an interface for interfacing with the shop endpoint64of the
// Shopify API.
// See: https://help.shopify.com/api/reference/shop
type ShopService interface {
	Get(options interface{}) (*Shop, error)
}

// ShopServiceOp handles communication with the shop related methods of the
// Shopify API.
type ShopServiceOp struct {
	client *Client
}

// Shop represents a Shopify shop
type Shop struct {
	ID                      int64       `json:"id"`
	Name                    string     `json:"name"`
	ShopOwner               string     `json:"shop_owner"`
	Email                   string     `json:"email"`
	CreatedAt               *time.Time `json:"created_at"`
	UpdatedAt               *time.Time `json:"updated_at"`
	Address1                string     `json:"address1"`
	City                    string     `json:"city"`
	Country                 string     `json:"country"`
	CountryCode             string     `json:"country_code"`
	CountryName             string     `json:"country_name"`
	Currency                string     `json:"currency"`
	Domain                  string     `json:"domain"`
	Latitude                float64    `json:"latitude"`
	Longitude               float64    `json:"longitude"`
	Phone                   string     `json:"phone"`
	Province                string     `json:"province"`
	ProvinceCode            string     `json:"province_code"`
	Zip                     string     `json:"zip"`
	MoneyFormat             string     `json:"money_format"`
	MoneyWithCurrencyFormat string     `json:"money_with_currency_format"`
	MyshopifyDomain         string     `json:"myshopify_domain"`
	PlanName                string     `json:"plan_name"`
	PlanDisplayName         string     `json:"plan_display_name"`
	PasswordEnabled         bool       `json:"password_enabled"`
	PrimaryLocale           string     `json:"primary_locale"`
	Timezone                string     `json:"timezone"`
	IanaTimezone            string     `json:"iana_timezone"`
	ForceSSL                bool       `json:"force_ssl"`
	HasStorefront           bool       `json:"has_storefront"`
	HasDiscounts            bool       `json:"has_discounts"`
	HasGiftcards            bool       `json:"has_gift_cards"`
	SetupRequire            bool       `json:"setup_required"`
	CountyTaxes             bool       `json:"county_taxes"`
}

// Represents the result from the admin/shop.json endpoint
type ShopResource struct {
	Shop *Shop `json:"shop"`
}

// Get shop
func (s *ShopServiceOp) Get(options interface{}) (*Shop, error) {
	resource := new(ShopResource)
	err := s.client.Get("admin/shop.json", resource, options)
	return resource.Shop, err
}
