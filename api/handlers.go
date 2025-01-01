package api

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/Stupnikjs/musicsite/util"
	"github.com/go-chi/chi/v5"
)

var pathToTemplates = "./static/templates/"

type TemplateData struct {
	Data map[string]any
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
	td.Data["songs"] = util.ListAllSong()
	_ = render(w, r, "/main.gohtml", &td)
}

func (app *Application) RenderPlaylist(w http.ResponseWriter, r *http.Request) {

	td := TemplateData{}
	td.Data = make(map[string]any)
	td.Data["songs"] = util.ListPlaylist()
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
		_, _ = io.Copy(w, file)
	}
}
