package utils

import (
	"net/http"
	"text/template"
)

func MainRenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	// Загружаем указанный шаблон и подставляем переданные данные.
	tmpl = "templates/" + tmpl + ".html"
	t, err := template.New("").ParseFiles(tmpl, "templates/base.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
