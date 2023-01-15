package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRubles(t *testing.T) {
	type want struct {
		currency Currency
		hasError bool
		err      error
	}

	tests := []struct {
		name  string
		value float64
		want  want
	}{
		{
			name:  "positive value",
			value: 1.62,
			want: want{
				currency: &rubles{
					name:  "руб.",
					value: 162,
				},
				hasError: false,
				err:      nil,
			},
		},
		{
			name:  "negative value",
			value: -6123612.77,
			want: want{
				currency: &rubles{
					name:  "руб.",
					value: -612361277,
				},
				hasError: true,
				err:      ErrNegativePrice,
			},
		},
		{
			name:  "zero",
			value: 0,
			want: want{
				currency: &rubles{
					name:  "руб.",
					value: 0,
				},
				hasError: false,
				err:      nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := NewRubles(tt.value)
			if tt.want.hasError {
				assert.ErrorIs(t, err, tt.want.err)
			} else {
				assert.Equal(t, tt.want.currency, actual)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	type want struct {
		value float64
	}

	tests := []struct {
		name     string
		currency Currency
		want     want
	}{
		{
			name: "rubles - positive value",
			currency: &rubles{
				value: 31566,
			},
			want: want{
				value: 315.66,
			},
		},
		{
			name: "rubles - zero",
			currency: &rubles{
				value: 0,
			},
			want: want{
				value: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.currency.Float64()
			assert.Equal(t, tt.want.value, actual)
		})
	}
}

func TestRaw(t *testing.T) {
	type want struct {
		value int64
	}

	tests := []struct {
		name     string
		currency Currency
		want     want
	}{
		{
			name: "positive value",
			currency: &rubles{
				name:  "руб.",
				value: 32199,
			},
			want: want{value: 32199},
		},
		{
			name: "zero",
			currency: &rubles{
				name:  "руб.",
				value: 32199,
			},
			want: want{value: 32199},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.currency.Raw()
			assert.Equal(t, tt.want.value, actual)
		})
	}
}

func TestString(t *testing.T) {
	type want struct {
		printedCurrency string
	}

	tests := []struct {
		name     string
		currency Currency
		want     want
	}{
		{
			name: "positive value",
			currency: &rubles{
				name:  "руб.",
				value: 6100,
			},
			want: want{printedCurrency: "61.00 руб."},
		},
		{
			name: "zero",
			currency: &rubles{
				name:  "руб.",
				value: 0,
			},
			want: want{printedCurrency: "0.00 руб."},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.currency.String()
			assert.Equal(t, tt.want.printedCurrency, actual)
		})
	}
}
