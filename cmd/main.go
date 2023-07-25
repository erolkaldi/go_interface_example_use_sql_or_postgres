package main

import "cmd/main.go/app"

func main() {
	a := app.App{}
	if a.InitializeDB() {
		a.Routes()
		a.Run()
	}
}
