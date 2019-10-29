package repository

type Repository interface {
	GetTopTen() ([]string, error)
	GetSong(artist string) (string, error)
	GetAlbum(artis string) (string, error)
}
