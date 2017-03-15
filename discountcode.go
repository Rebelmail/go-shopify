package goshopify

import (
	"github.com/shopspring/decimal"
)

type DiscountCode struct {
	Amount *decimal.Decimal `json:"amount"`
	Code   string           `json:"code"`
	Type   string           `json:"type"`
}
