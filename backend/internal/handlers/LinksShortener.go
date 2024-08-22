package handlers

import (
	"LinksShortener/internal/domain"
	_ "LinksShortener/internal/domain"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"os"
)

type ShortenerService interface {
	SetLink(fullLink string) (string, bool, error) // returns shortLink, existed, error
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

var getShortLinkCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_get_shortLink_count", // metric name
		Help: "Number of get_shortLink requests.",
	},
	[]string{"status"}, // labels
)

var redirectCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_redirect_count", // metric name
		Help: "Number of redirect requests.",
	},
	[]string{"status"}, // labels
)

func init() {
	prometheus.MustRegister(getShortLinkCounter)
	prometheus.MustRegister(redirectCounter)
}

func (s ShortenerHandler) SetLink(w http.ResponseWriter, r *http.Request) {
	var linkIn domain.LinksIn
	var status string

	defer func() {
		getShortLinkCounter.WithLabelValues(status).Inc()
	}()

	err := json.NewDecoder(r.Body).Decode(&linkIn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		status = "error"
		return
	}
	shortLink, existed, err := s.service.SetLink(linkIn.FullLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		status = "error"
		return
	}
	response, err := json.Marshal(&domain.SetLinkResponse{
		ShortLink: s.domain + "/" + shortLink,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		status = "error"
		return
	}
	if !existed {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	status = "success"
	w.Write(response)
}

func (s ShortenerHandler) GetLink(w http.ResponseWriter, r *http.Request) {
	var status string
	defer func() {
		redirectCounter.WithLabelValues(status).Inc()
	}()
	shortLink := chi.URLParam(r, "shortLink")
	fullLink, err := s.service.GetLink(shortLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		status = "error"
		return
	}
	status = "success"
	http.Redirect(w, r, fullLink, http.StatusTemporaryRedirect)
}
