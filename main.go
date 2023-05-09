package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"fontstudios/handlers"
)

func main() {
	fmt.Println("Starting Server")
	// Создаем новый маршрутизатор Gorilla.
	r := mux.NewRouter()

	// Привязывает хендлеры
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/order", handlers.OrderHandler)
	r.HandleFunc("/vacancy", handlers.VacancyHandler)
	r.HandleFunc("/contacts", handlers.ContactsHandler)

	// Настраиваем директорию для статических файлов.
	staticDir := "/static/"
	staticHandler := http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir)))
	r.PathPrefix(staticDir).Handler(staticHandler)

	// Запускаем сервер на порту 80.
	err := http.ListenAndServe(":80", r)
	if err != nil {
		panic(err)
	}
}
