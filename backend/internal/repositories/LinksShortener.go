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

func (r *ShortenerRepository) SetLink(fullLink, shortLink string) (string, error) {
	return "", nil
}

func (r *ShortenerRepository) GetLinkIfExist(fullLink string) (shortLink string, isFound bool, err error) {
	return "", true, nil
}
