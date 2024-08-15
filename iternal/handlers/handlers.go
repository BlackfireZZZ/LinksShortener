package handlers

type Handlers struct {
	Shortener ShortenerHandler
}

func InitHandlers(Shortener ShortenerService) *Handlers {
	return &Handlers{
		Shortener: *NewShortenerHandler(Shortener),
	}
}
