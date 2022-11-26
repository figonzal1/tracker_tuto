package types

import "time"

type Transaction struct {
	Id       string
	DateTime time.Time
	Coin     string
	Value    float64
}
