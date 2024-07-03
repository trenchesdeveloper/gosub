package main

import "net/http"

func (app *Config) HomePage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.gohtml", nil)
}

func (app *Config) Login(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.gohtml", nil)
}

func (app *Config) PostLogin(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// email := r.PostForm.Get("email")
	// password := r.PostForm.Get("password")
	// fmt.Println(email, password)
	// http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Config) Logout(w http.ResponseWriter, r *http.Request) {
	// app.Session.Remove(r.Context(), "user_id")
	// app.Session.Put(r.Context(), "flash", "You've been logged out successfully!")
	// http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Config) RegisterPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "register.page.gohtml", nil)
}

func (app *Config) PostRegister(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// email := r.PostForm.Get("email")
	// password := r.PostForm.Get("password")
	// fmt.Println(email, password)
	// http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Config) ActivateAccount(w http.ResponseWriter, r *http.Request) {
	// email := r.URL.Query().Get("email")
	// token := r.URL.Query().Get("token")
	// fmt.Println(email, token)
	// http.Redirect(w, r, "/", http.StatusSeeOther)
}