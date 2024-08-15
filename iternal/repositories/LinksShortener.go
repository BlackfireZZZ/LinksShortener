package repositories

import "github.com/jackc/pgx/v4/pgxpool"

type ShortenerRepository struct {
	db *pgxpool.Pool
}

func NewShortenerRepository(db *pgxpool.Pool) *ShortenerRepository {
	return &ShortenerRepository{
		db: db,
	}
}

func (r *ShortenerRepository) SetLink(FullLink, ShortLink string) (string, error) {
	return "", nil
}

func (r *ShortenerRepository) CheckLinkExists(FullLink string) (string, error) {
	return "", nil
}
