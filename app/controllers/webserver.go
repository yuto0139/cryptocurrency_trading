package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"cryptocurrency_trading/app/models"
	"cryptocurrency_trading/config"
)

var templates = template.Must(template.ParseFiles("app/views/google.html"))

// viewChartHandler ...
func viewChartHandler(w http.ResponseWriter, r *http.Request) {
	limit := 100
	duration := "1m"
	durationTime := config.Config.Durations[duration]
	df, _ := models.GetAllCandle(config.Config.ProductCode, durationTime, limit)

	err := templates.ExecuteTemplate(w, "google.html", df.Candles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// StartWebServer ...
func StartWebServer() error {
	http.HandleFunc("/chart/", viewChartHandler)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
