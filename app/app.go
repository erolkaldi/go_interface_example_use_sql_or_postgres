package app

import (
	"cmd/main.go/database"
	"cmd/main.go/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
	Db     *gorm.DB
	Config *models.Config
}

func (app *App) InitializeDB() bool {
	app.Config = &models.Config{}
	app.Config.GetConfigValues()
	var db models.BaseDb
	db = &database.Postgress{}
	var err error
	app.Db, err = db.Connect(app.Config)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true

}
func (a *App) Routes() {
	a.Router = mux.NewRouter()
}
func (a *App) Run() {
	fmt.Printf("Server started at %s\n", a.Config.Api.Port)
	log.Fatal(http.ListenAndServe(a.Config.Api.Port, a.Router))
}
