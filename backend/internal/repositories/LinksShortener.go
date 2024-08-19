package repositories

import (
	"LinksShortener/internal/domain"
	_ "LinksShortener/internal/domain"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ShortenerRepository struct {
	db *pgxpool.Pool
}

func NewShortenerRepository(db *pgxpool.Pool) *ShortenerRepository {
	return &ShortenerRepository{
		db: db,
	}
}

func (r *ShortenerRepository) SetLink(fullLink, shortLink string) (string, error) {
	_, err := r.db.Exec(context.Background(), `INSERT INTO links (full_link, short_link) VALUES ($1, $2)`, fullLink, shortLink)
	if err != nil {
		return "", err
	}
	return shortLink, nil
}

func (r *ShortenerRepository) GetLinkIfExist(fullLink string) (shortLink string, isFound bool, err error) {
	linksOut := &domain.LinksOut{}
	rows, err := r.db.Query(context.Background(), `SELECT short_link FROM links WHERE full_link = $1`, fullLink)
	if err != nil {
		return "", false, err
	}
	defer rows.Close()

	if !rows.Next() {
		return "", false, nil
	}
	for rows.Next() {
		err = rows.Scan(&linksOut)
		if err != nil {
			return "", false, err
		}
	}
	shortLink = linksOut.ShortLink
	return shortLink, true, nil
}
