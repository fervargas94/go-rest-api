package recomendations

import "errors"

type spotify struct {
}

// create new repository
func NewSpotifyRepo() *spotify {
	return &spotify{}
}

func (s *spotify) GetSongs() ([]string, error) {
	return []string{}, errors.New(" spotify GetSongs not implemented")
}

func (s *spotify) GetSongById(id string) (string, error) {
	return "", errors.New("spotify GetSongById not implemented")
}
