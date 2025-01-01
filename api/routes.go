package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Application struct {
	Port int
 DB *database.Repo
}

func (app *Application) Routes() http.Handler {

	mux := chi.NewRouter()

	// register routes
	mux.Get("/", app.RenderAccueil)
	mux.Get("/playlist", app.RenderPlaylist)
	mux.Get("/song/compo/{id}", app.GetHandlerSong(true))
	mux.Get("/song/free/{id}", app.GetHandlerSong(false))
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}
