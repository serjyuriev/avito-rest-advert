package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRubles(t *testing.T) {
	type want struct {
		currency Rubles
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
				currency: Rubles(162),
				hasError: false,
				err:      nil,
			},
		},
		{
			name:  "negative value",
			value: -6123612.77,
			want: want{
				currency: 0,
				hasError: true,
				err:      ErrNegativePrice,
			},
		},
		{
			name:  "zero",
			value: 0,
			want: want{
				currency: Rubles(0),
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
		currency Rubles
		want     want
	}{
		{
			name:     "rubles - positive value",
			currency: Rubles(31566),
			want:     want{value: 315.66},
		},
		{
			name:     "rubles - zero",
			currency: Rubles(0),
			want:     want{value: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.currency.Float64()
			assert.Equal(t, tt.want.value, actual)
		})
	}
}
