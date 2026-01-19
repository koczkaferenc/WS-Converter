package models

import (
	"database/sql"

	"github.com/shopspring/decimal"
)

// KS Termék adatok
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

// Gyártók
var Manufacturers = map[int]string{
    0:  "Prémium",      // Renold
    1:  "Távol-keleti", // MSC
    2:  "Távol-keleti", // Lovas
    3:  "Távol-keleti", // TEC
    4:  "Európai",      // Codex
    5:  "Európai",      // Vamberk
    6:  "Európai",      // Strakonice
    7:  "Prémium",      // Rexnord
    8:  "Európai",      // Link-Belt
    9:  "Európai",      // Retezarna
    10: "Európai",      // Reiter
    11: "Távol-keleti",
    12: "AL",
}
