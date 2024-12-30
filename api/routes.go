package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Application struct {
	Port int
}

func (app *Application) Routes() http.Handler {

	mux := chi.NewRouter()

	// register routes
	mux.Get("/", app.RenderAccueil)
	mux.Get("/playlist", app.RenderPlaylist)
	mux.Get("/song/{id}", app.GetSong)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}
