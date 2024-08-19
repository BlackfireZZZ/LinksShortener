package services

type Services struct {
	Shortener ShortenerService
}

func InitServices(repository ShortenerRepository) *Services {
	return &Services{
		Shortener: *NewShortenerService(repository),
	}
}
