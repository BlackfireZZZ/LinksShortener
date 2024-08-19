package handlers

import (
	"LinksShortener/internal/domain"
	_ "LinksShortener/internal/domain"
	"encoding/json"
	"net/http"
)

type ShortenerService interface {
	SetLink(fullLink string) (string, error)
}

type ShortenerHandler struct {
	service ShortenerService
}

func NewShortenerHandler(service ShortenerService) *ShortenerHandler {
	return &ShortenerHandler{
		service: service,
	}
}

func (s ShortenerHandler) Shortener(w http.ResponseWriter, r *http.Request) {
	var linkIn domain.LinksIn
	err := json.NewDecoder(r.Body).Decode(&linkIn)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	shortLink, err := s.service.SetLink(linkIn.FullLink)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortLink))
}
