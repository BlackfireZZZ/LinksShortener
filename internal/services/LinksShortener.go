package services

import (
	"fmt"
	"golang.org/x/crypto/blake2b"
	"os"
	"strconv"
)

type ShortenerRepository interface {
	SetLink(FullLink, ShortLink string) (string, error)
	GetLinkIfExist(fullLink string) (shortLink string, isFound bool, err error)
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

func (s *ShortenerService) SetLink(fullLink string) (string, error) {
	link, exists, err := s.repo.GetLinkIfExist(fullLink)
	if err != nil {
		return "", err
	} else if !exists {
		linkLengthString := os.Getenv("LINK_LENGTH")
		linkLength, err := strconv.Atoi(linkLengthString)
		if err != nil {
			return "", err
		}
		shortLink, err := GenerateShortLink(fullLink, linkLength)
		if err != nil {
			return "", err
		}
		shortLink, err = s.repo.SetLink(fullLink, shortLink)
		if err != nil {
			return "", err
		}
		return shortLink, nil
	} else {
		return link, nil
	}
}
