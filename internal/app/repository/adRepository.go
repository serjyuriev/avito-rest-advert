package repository

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"
	e "github.com/serjyuriev/avito-rest-advert/internal/pkg/models/entity"
)

type AdRepository interface {
	Create(ctx context.Context, ad *e.Ad) error
	ReadOne(ctx context.Context, id string) (*e.Ad, error)
	ReadAll(ctx context.Context) ([]*e.Ad, error)
}

type adRepository struct {
	db  *sql.DB
	lgr zerolog.Logger
}

func NewAdRepository(ctx context.Context, logger zerolog.Logger) (AdRepository, error) {
	logger.Info().Msg("setting up ad repository")

	// TODO: replace nil with postgres db initialization
	repo := &adRepository{
		db:  nil,
		lgr: logger,
	}

	logger.Info().Msg("repository was successfully set up")
	return repo, nil
}

func (r *adRepository) Create(ctx context.Context, ad *e.Ad) error {
	r.lgr.Debug().Msg("inserting ad in the database")

	if _, err := r.db.ExecContext(
		ctx,
		`INSERT INTO ads(id, name, description, photos, price, added)
		VALUES ($1, $2, $3, $4, $5, $6);`,
		ad.Id,
		ad.Name,
		ad.Description,
		ad.AllPhotos,
		ad.Price,
		ad.Added,
	); err != nil {
		r.lgr.Error().Msg("unable to insert ad in the database")
		return err
	}

	r.lgr.Debug().Msg("ad was successfully inserted in the database")
	return nil
}

func (r *adRepository) ReadOne(ctx context.Context, id string) (*e.Ad, error) {
	r.lgr.Debug().Str("id", id).Msg("selecting ad from the database")

	row := r.db.QueryRowContext(
		ctx,
		`SELECT name, description, price
		FROM ads
		WHERE id = $1;`,
		id,
	)

	ad := &e.Ad{}
	d := sql.NullString{}
	if err := row.Scan(&ad.Name, &d, &ad.Price); err != nil {
		r.lgr.Error().Str("id", id).Msg("unable to select ad from the database")
		return nil, err
	}

	if d.Valid {
		ad.Description = d.String
	}

	rows, err := r.db.QueryContext(
		ctx,
		"SELECT photo FROM adPhotos WHERE ad_id = $1;",
		id,
	)
	if err != nil {
		r.lgr.Error().Str("id", id).Msg("unable to select ad photos from the database")
		return nil, err
	}

	ad.AllPhotos = make([]string, 0)
	for rows.Next() {
		var ph string
		if err = rows.Scan(&ph); err != nil {
			r.lgr.Error().Str("id", id).Msg("unable to scan ad photo to variable")
			return nil, err
		}
		ad.AllPhotos = append(ad.AllPhotos, ph)
	}

	r.lgr.Debug().Str("id", id).Msg("ad was successfully selected from the database")
	return ad, nil
}

func (r *adRepository) ReadAll(ctx context.Context) ([]*e.Ad, error) {
	r.lgr.Debug().Msg("selecting all ads from the database")

	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, name, price FROM ads;",
	)
	if err != nil {
		r.lgr.Error().Msg("unable to select ads from the database")
		return nil, err
	}

	ads := make([]*e.Ad, 0)
	for rows.Next() {
		ad := new(e.Ad)
		if err = rows.Scan(&ad.Id, &ad.Name, &ad.Price); err != nil {
			r.lgr.Error().Msg("unable to scan ad to variable")
			return nil, err
		}

		photoRows, err := r.db.QueryContext(
			ctx,
			"SELECT photo FROM adPhotos WHERE ad_id = $1;",
			ad.Id,
		)
		if err != nil {
			r.lgr.Error().Str("id", ad.Id).Msg("unable to select ad photos from the database")
			return nil, err
		}

		ad.AllPhotos = make([]string, 0)
		for photoRows.Next() {
			var ph string
			if err = photoRows.Scan(&ph); err != nil {
				r.lgr.Error().Str("id", ad.Id).Msg("unable to scan ad photo to variable")
				return nil, err
			}
			ad.AllPhotos = append(ad.AllPhotos, ph)
		}

		ads = append(ads, ad)
	}

	r.lgr.Debug().Int("amount", len(ads)).Msg("all ads where successfully selected from the database")
	return ads, nil
}
