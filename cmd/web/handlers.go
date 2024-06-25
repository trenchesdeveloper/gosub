package main

import "net/http"

func (app *Config) HomePage(w http.ResponseWriter, r *http.Request) {
	app.InfoLog.Println("home")
	w.Write([]byte("Home"))
}
