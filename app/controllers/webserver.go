package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"cryptocurrency_trading/config"
)

var templates = template.Must(template.ParseFiles("app/views/google.html"))

// viewChartHandler ...
func viewChartHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "google.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// StartWebServer ...
func StartWebServer() error {
	http.HandleFunc("/chart/", viewChartHandler)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
