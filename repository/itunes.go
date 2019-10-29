package repository

import "github.com/pkg/errors"

type itunes struct {
}

func (s *itunes) GetTopTen() ([]string, error) {
	return []string{}, errors.New("not implemented")
}

func (s *itunes) GetSong(artist string) (string, error) {
	return "", errors.New("not implemented")
}

func (s *itunes) GetAlbum() (string, error) {
	return "", errors.New("not implemented")
}

