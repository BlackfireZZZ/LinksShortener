package handlers

import (
	"LinksShortener/internal/domain"
	_ "LinksShortener/internal/domain"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
)

type ShortenerService interface {
	SetLink(fullLink string) (string, bool, error)
	GetLink(shortLink string) (string, error)
}

type ShortenerHandler struct {
	service ShortenerService
	domain  string
}

func NewShortenerHandler(service ShortenerService) *ShortenerHandler {
	return &ShortenerHandler{
		service: service,
		domain:  os.Getenv("DOMAIN"),
	}
}

func (s ShortenerHandler) SetLink(w http.ResponseWriter, r *http.Request) {
	var linkIn domain.LinksIn
	err := json.NewDecoder(r.Body).Decode(&linkIn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	shortLink, existed, err := s.service.SetLink(linkIn.FullLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := json.Marshal(&domain.SetLinkResponse{
		ShortLink: s.domain + "/" + shortLink,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !existed {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(response)
}

func (s ShortenerHandler) GetLink(w http.ResponseWriter, r *http.Request) {
	shortLink := chi.URLParam(r, "shortLink")
	fullLink, err := s.service.GetLink(shortLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//response, err := json.Marshal(&domain.GetLinkResponse{
	//	FullLink: fullLink,
	//})
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//w.WriteHeader(http.StatusOK)
	//w.Write(response)
	http.Redirect(w, r, fullLink, http.StatusTemporaryRedirect)
}
