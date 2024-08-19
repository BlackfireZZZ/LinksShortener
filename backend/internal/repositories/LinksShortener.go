package repositories

import (
	"LinksShortener/internal/domain"
	_ "LinksShortener/internal/domain"
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
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

func (r *ShortenerRepository) GetShortLinkIfExist(fullLink string) (shortLink string, isFound bool, err error) {
	linksOut := &domain.LinksOut{}
	err = r.db.QueryRow(context.Background(), `SELECT full_link, short_link FROM links WHERE full_link = $1`, fullLink).Scan(&linksOut.FullLink, &linksOut.ShortLink)
	if errors.Is(err, pgx.ErrNoRows) || errors.Is(err, sql.ErrNoRows) {
		return "", false, nil
	} else if err != nil {
		return "", false, err
	}
	shortLink = linksOut.ShortLink
	return shortLink, true, nil
}

func (r *ShortenerRepository) GetFullLinkIfExist(shortLink string) (fullLink string, isFound bool, err error) {
	linksOut := &domain.LinksOut{}
	err = r.db.QueryRow(context.Background(), `SELECT full_link, short_link FROM links WHERE short_link = $1`, shortLink).Scan(&linksOut.FullLink, &linksOut.ShortLink)
	if errors.Is(err, pgx.ErrNoRows) || errors.Is(err, sql.ErrNoRows) {
		return "", false, nil
	} else if err != nil {
		log.Println("Repo: ", err)
		return "", false, err
	}

	fullLink = linksOut.FullLink
	return fullLink, true, nil
}
