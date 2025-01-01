package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Stupnikjs/musicsite/api"
	"github.com/Stupnikjs/musicsite/database"

	_ "google.golang.org/api/option"
)

func main() {

	uri := os.Getenv("DATABASE_URL")
	conn, err := api.ConnectDB(uri)

	if err != nil {
		fmt.Println(err)
	}
	app := api.Application{
		Port: 8080,
		DB: &database.PostgresRepo{
			DB: conn,
		},
	}
	http.ListenAndServe(fmt.Sprintf(":%d", app.Port), app.Routes())

}
