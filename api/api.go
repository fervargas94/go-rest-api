package api

import(
	"github.com/fervargas94/go-rest-api/repository"
	"github.com/pkg/errors"
)

type Api interface {
	GetTopTen() ([]string, error)
	GetSong(artist string) (string, error)
	GetAlbum(artis string) (string, error)
}

type api struct {
	repository repository.Repository
}

func (a *api) GetTopTen() ([]string, error){
	return []string{}, errors.New("not implemented")
}

func (a *api) GetSong(artist string) (string, error) {
	return "", errors.New("not implemented")
}

func (a *api) GetAlbum(artist string) (string, error) {
	return "", errors.New("not implemented")
}