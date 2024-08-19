package handlers

import (
	"LinksShortener/internal/domain"
	_ "LinksShortener/internal/domain"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type ShortenerService interface {
	SetLink(fullLink string) (string, error)
	GetLink(shortLink string) (string, error)
}

type ShortenerHandler struct {
	service ShortenerService
}

func NewShortenerHandler(service ShortenerService) *ShortenerHandler {
	return &ShortenerHandler{
		service: service,
	}
}

func (s ShortenerHandler) SetLink(w http.ResponseWriter, r *http.Request) {
	var linkIn domain.LinksIn
	err := json.NewDecoder(r.Body).Decode(&linkIn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	shortLink, err := s.service.SetLink(linkIn.FullLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("shortLink: ", shortLink)
	w.Write([]byte(shortLink))
	w.WriteHeader(http.StatusCreated)
}

func (s ShortenerHandler) GetLink(w http.ResponseWriter, r *http.Request) {
	shortLink := chi.URLParam(r, "shortLink")
	fullLink, err := s.service.GetLink(shortLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fullLink))
}
