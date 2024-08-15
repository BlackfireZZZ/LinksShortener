package handlers

import "net/http"

type ShortenerService interface {
	SetLink(FullLink string) (string, error)
}

type ShortenerHandler struct {
	service ShortenerService
}

func NewShortenerHandler(service ShortenerService) *ShortenerHandler {
	return &ShortenerHandler{
		service: service,
	}
}

func (s ShortenerHandler) Shortener(r http.ResponseWriter, w *http.Request) {
	return
}
