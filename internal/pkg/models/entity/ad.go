package models

import (
	"errors"
	"time"
)

var (
	ErrAdNameIsTooLong        = errors.New("ad name length cannot exceed 200 characters")
	ErrAdDescriptionIsTooLong = errors.New("ad description length cannot exceed 1000 characters")
	ErrTooManyPhotos          = errors.New("ad cannot have more than 3 photos")
)

type Ad struct {
	Id          string
	Name        string
	Description string
	MainPhoto   string
	AllPhotos   []string
	Price       Rubles
	Added       time.Time
}

func NewAd(name string, description string, photos []string, price Rubles) (*Ad, error) {
	if len(name) > 200 {
		return nil, ErrAdNameIsTooLong
	}

	if len(description) > 1000 {
		return nil, ErrAdDescriptionIsTooLong
	}

	if len(photos) > 3 {
		return nil, ErrTooManyPhotos
	}

	ad := &Ad{
		Name:        name,
		Description: description,
		AllPhotos:   photos,
		Price:       price,
		Added:       time.Now(),
	}

	if len(photos) > 0 {
		ad.MainPhoto = photos[0]
	}

	return ad, nil
}
