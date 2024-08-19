package services

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/blake2b"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type ShortenerRepository interface {
	SetLink(fullLink, shortLink string) (string, error)
	GetLinkIfExist(fullLink string) (shortLink string, isFound bool, err error)
}

type ShortenerService struct {
	repo       ShortenerRepository
	linkLength int
}

func NewShortenerService(repo ShortenerRepository) *ShortenerService {
	linkLengthString := os.Getenv("LINK_LENGTH")
	linkLength, err := strconv.Atoi(linkLengthString)
	if err != nil {
		return nil
	}
	return &ShortenerService{
		repo:       repo,
		linkLength: linkLength,
	}
}

func (s *ShortenerService) generateShortLink(fullLink string, length int) (string, error) {
	hash, err := blake2b.New(length, nil)
	if err != nil {
		return "", err
	}
	hash.Write([]byte(fullLink))
	return fmt.Sprintf("%x", hash.Sum(nil)), err
}

func isValidURL(str string) bool {
	u, err := url.Parse(str)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	if !strings.HasPrefix(u.Scheme, "http") {
		return false
	}

	return true
}

func (s *ShortenerService) SetLink(fullLink string) (string, error) {
	if !isValidURL(fullLink) {
		return "", errors.New("invalid URL")
	}
	link, exists, err := s.repo.GetLinkIfExist(fullLink)
	if err != nil {
		return "", err
	} else if !exists {
		shortLink, err := s.generateShortLink(fullLink, s.linkLength)
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
