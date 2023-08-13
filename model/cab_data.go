package model

import "github.com/shopspring/decimal"

type CabData struct {
	Time    string
	Mileage string
	Price   decimal.Decimal
}
