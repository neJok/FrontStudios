package handlers

import (
	"fontstudios/utils"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	utils.MainRenderTemplate(w, "home", nil)
}
