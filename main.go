package main

import (
	"fmt"
	"recomendations"
)

func main() {
	itunesRepo := recomendations.NewItunesRepo()
	spotifyRepo := recomendations.NewSpotifyRepo()

	recSvc := recomendations.NewSvc(itunesRepo)
	spotifySvc := recomendations.NewSvc(spotifyRepo)
	fmt.Println(recSvc.GetTopTen())
	fmt.Println(spotifySvc.GetTopTen())
}
