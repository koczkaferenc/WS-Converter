package models

import (
	"database/sql"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID            int
	Code          string
	Name          string
	Unit          sql.NullString
	Weight        decimal.NullDecimal
	Stock         decimal.NullDecimal
	WebPrice      decimal.NullDecimal
	ReSellerPrice decimal.NullDecimal
}
