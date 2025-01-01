package api

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/go-chi/chi/v5"
)

var pathToTemplates = "./static/templates/"

type TemplateData struct {
	Data map[string]any
}

type Song struct {
	Artist string
	Name   string
	YtLink string
	Path   string
}

func render(w http.ResponseWriter, r *http.Request, t string, td *TemplateData) error {
	_ = r.Method

	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, t), path.Join(pathToTemplates, "/base.layout.gohtml"))
	if err != nil {
		return err
	}
	err = parsedTemplate.Execute(w, td)
	if err != nil {
		return err
	}
	return nil

}

// template rendering

func (app *Application) RenderAccueil(w http.ResponseWriter, r *http.Request) {

	td := TemplateData{}
	td.Data = make(map[string]any)
	td.Data["songs"] = listAllSong()
	_ = render(w, r, "/main.gohtml", &td)
}
func (app *Application) RenderPlaylist(w http.ResponseWriter, r *http.Request) {

	td := TemplateData{}
	td.Data = make(map[string]any)
	td.Data["songs"] = listPlaylist()
	_ = render(w, r, "/playlist.gohtml", &td)
}

func (app *Application) GetHandlerSong(isCompo bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		var songPath string
		if isCompo {
			songPath = path.Join(cwd, "static", "track", "compo", id)
		} else {
			songPath = path.Join(cwd, "static", "track", "free", id)
		}

		file, err := os.Open(songPath)
		if err != nil {
			fmt.Println(err)
		}
		_, err = io.Copy(w, file)
	}
}

func listAllSong() []string {
	var trackNames []string
	cur, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	trackPath := path.Join(cur, "static", "track", "compo")
	entries, err := os.ReadDir(trackPath)

	for _, e := range entries {
		trackNames = append(trackNames, e.Name())
	}

	return trackNames

}

func listPlaylist() []Song {

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
