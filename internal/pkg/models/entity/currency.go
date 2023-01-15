package models

import (
	"errors"
	"fmt"
)

var (
	ErrNegativePrice = errors.New("ad price cannot be less than zero")
)

type Currency interface {
	Float64() float64
	Raw() int64
	String() string
}

type rubles struct {
	name  string
	value int64
}

func NewRubles(f float64) (Currency, error) {
	if f < 0 {
		return nil, ErrNegativePrice
	}

	return &rubles{
		name:  "руб.",
		value: int64((f * 100) + 0.5),
	}, nil
}

func (r *rubles) Float64() float64 {
	return float64(r.value) / 100
}

func (r *rubles) Raw() int64 {
	return r.value
}

func (r *rubles) String() string {
	return fmt.Sprintf(
		"%.2f %s",
		r.Float64(),
		r.name,
	)
}
