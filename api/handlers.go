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
	fmt.Println(td)
	_ = render(w, r, "/main.gohtml", &td)
}

func (app *Application) GetSong(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	songPath := path.Join(cwd, "static", "track", id)
	fmt.Println(songPath)
	file, err := os.Open(songPath)
	fmt.Println(err, id)
	if err != nil {
		fmt.Println(err)
	}
	_, err = io.Copy(w, file)
}

func listAllSong() []string {
	var trackNames []string
	cur, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	trackPath := path.Join(cur, "static", "track")
	entries, err := os.ReadDir(trackPath)

	for _, e := range entries {
		trackNames = append(trackNames, e.Name())
	}

	return trackNames

}
