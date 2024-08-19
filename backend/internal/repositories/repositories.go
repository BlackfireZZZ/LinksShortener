package repositories

import "github.com/jackc/pgx/v4/pgxpool"

type Repositories struct {
	Shortener ShortenerRepository
}

func InitRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		Shortener: *NewShortenerRepository(db),
	}
}
