package recomendations

func NewSvc(r repository) *RecommendationsSvc {
	return &RecommendationsSvc{
		repo: r,
	}
}

//Service  interface, here we will do all the business Logic
type RecommendationsSvc struct {
	repo repository
}

func (a *RecommendationsSvc) GetTopTen() ([]string, error) {
	return a.repo.GetSongs()
}

func (a *RecommendationsSvc) GetSong(artist string) (string, error) {
	return a.repo.GetSongById(artist)
}

func (a *RecommendationsSvc) GetAlbum(artist string) ([]string, error) {
	return a.repo.GetSongs()
}
