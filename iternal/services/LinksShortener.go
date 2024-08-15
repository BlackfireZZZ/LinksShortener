package services

import (
	"fmt"
	"golang.org/x/crypto/blake2b"
)

type ShortenerRepository interface {
	SetLink(FullLink, ShortLink string) (string, error)
	CheckLinkExists(FullLink string) (string, error)
}

type ShortenerService struct {
	repo ShortenerRepository
}

func NewShortenerService(repo ShortenerRepository) *ShortenerService {
	return &ShortenerService{
		repo: repo,
	}
}

func GenerateShortLink(FullLink string, length int) (string, error) {
	hash, err := blake2b.New(length, nil)
	if err != nil {
		return "", err
	}
	hash.Write([]byte(FullLink))
	return fmt.Sprintf("%x", hash.Sum(nil)), err
}

func (s *ShortenerService) SetLink(FullLink string) (string, error) {
	Link, err := s.repo.CheckLinkExists(FullLink)
	if err != nil {
		return "", err
	} else if Link == "" {
		ShortLink, err := GenerateShortLink(FullLink, 10)
		if err != nil {
			return "", err
		}
		ShortLink, err = s.repo.SetLink(FullLink, ShortLink)
		if err != nil {
			return "", err
		}
		return ShortLink, nil
	} else {
		return Link, nil
	}
}
