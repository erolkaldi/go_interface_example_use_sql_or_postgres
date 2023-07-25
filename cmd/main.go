package main

import (
	"cmd/main.go/app"
	"cmd/main.go/database"
)

func main() {
	a := app.App{}
	if a.InitializeDB(&database.Postgress{}) {
		a.Routes()
		a.Run()
	}
}
