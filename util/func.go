package util

import (
	"fmt"
	"os"
	"path"
)

type Song struct {
	Artist string
	Name   string
	YtLink string
	Path   string
}

func ListAllSong() []string {
	var trackNames []string
	cur, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	trackPath := path.Join(cur, "static", "track", "compo")
	entries, err := os.ReadDir(trackPath)
	if err != nil {
		fmt.Println(err)
	}
	for _, e := range entries {
		trackNames = append(trackNames, e.Name())
	}

	return trackNames

}

func ListPlaylist() []Song {

	return []Song{
		{Artist: "Elektronomia",
			Name:   "Limitless",
			YtLink: "https://www.youtube.com/watch?v=cNcy3J4x62M",
			Path:   "Limitless.mp3",
		},
		{
			Artist: "Dandy Warhols",
			Name:   "Godless",
			YtLink: "https://www.youtube.com/watch?v=LduipA_XUJ8",
			Path:   "",
		},
		{
			Artist: "Martin Garrix and MOTi",
			Name:   "Virus",
			YtLink: "https://www.youtube.com/watch?v=iXIDtf1wP0g",
			Path:   "",
		},
		{
			Artist: "The Glitch Mob",
			Name:   "We can make the world stop",
			YtLink: "https://www.youtube.com/watch?v=H-k_Eg7zXuc",
			Path:   "",
		},
	}
}
