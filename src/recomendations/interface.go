package recomendations

//Writer REPO Needs to implement data layer reader
type reader interface {
	GetSongs() ([]string, error)
	GetSongById(id string) (string, error)
}

//Writer REPO Needs to implement data layer writer
type writer interface {
}

//Repository  Data layer INTERFACE
type repository interface {
	reader
	writer
}

//Recommendations Service will implemente this
type UseCase interface {
	repository
	GetTopTen() ([]string, error)
	GetSong(artist string) (string, error)
	GetAlbum(artis string) (string, error)
}
