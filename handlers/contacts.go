package handlers

import (
	"fontstudios/utils"
	"net/http"
)

func ContactsHandler(w http.ResponseWriter, r *http.Request) {
	utils.MainRenderTemplate(w, "contacts", nil)
}
