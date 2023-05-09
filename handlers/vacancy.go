package handlers

import (
	"fontstudios/utils"
	"net/http"
)

func VacancyHandler(w http.ResponseWriter, r *http.Request) {
	utils.MainRenderTemplate(w, "vacancy", nil)
}
