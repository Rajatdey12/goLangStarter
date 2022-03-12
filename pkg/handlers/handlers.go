package handlers

import (
	"net/http"

	"github.com/Rajatdey12/goLangStarter/pkg/utils"
)

// Ws is the web service handler func..
func Ws(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "home.html")
}

// About is the about page for the application..
func About(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "about.html")
}
