package repository

import "github.com/pkg/errors"

type spotify struct {
}

func (s *spotify) GetTopTen() ([]string, error) {
	return []string{}, errors.New("not implemented")
}

func (s *spotify) GetSong(artist string) (string, error) {
	return "", errors.New("not implemented")
}

func (s *spotify) GetAlbum() (string, error) {
	return "", errors.New("not implemented")
}
