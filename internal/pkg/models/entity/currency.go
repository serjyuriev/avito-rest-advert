package models

import (
	"errors"
)

var (
	ErrNegativePrice = errors.New("ad price cannot be less than zero")
)

type Rubles int64

func NewRubles(f float64) (Rubles, error) {
	if f < 0 {
		return 0, ErrNegativePrice
	}

	return Rubles((f * 100) + 0.5), nil
}

func (r Rubles) Float64() float64 {
	return float64(r) / 100
}
