package recomendations

import "errors"

type itunes struct {
}

// create new repository
func NewItunesRepo() *itunes {
	return &itunes{}
}

func (s *itunes) GetSongs() ([]string, error) {
	return []string{}, errors.New(" itunes GetSongs not implemented")
}

func (s *itunes) GetSongById(id string) (string, error) {
	return "", errors.New("itunes GetSongById not implemented")
}
