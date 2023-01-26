package main

import (
	"database/sql"
	"fmt"
	"github.com/getground/tech-tasks/backend/cmd/internal/handlers"
	"github.com/getground/tech-tasks/backend/cmd/internal/repositories"
	"github.com/getground/tech-tasks/backend/cmd/internal/usecases"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// init mysql.
	db, err := sql.Open("mysql", "root:Lanusla14.@tcp(127.0.0.1:3306)/getground")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	tableRepo := repositories.NewTableRepo(db)
	guestRepo := repositories.NewGuestRepo(db)
	guestUseCase := usecases.NewGuestUseCase(*guestRepo, *tableRepo)
	tableUseCase := usecases.NewTableUseCase(*tableRepo)
	guestHandler := handlers.NewGuestHandler(guestUseCase)
	tableHandler := handlers.NewTableHandler(tableUseCase)

	// routes
	r := mux.NewRouter()
	r.HandleFunc("/ping", handlerPing).Methods("GET")
	r.HandleFunc("/tables", tableHandler.CreateTable).Methods("POST")
	r.HandleFunc("/guest_list/{name}", guestHandler.AddGuestToGuestList).Methods("POST")
	r.HandleFunc("/guest_list", guestHandler.GetGuestsList).Methods("GET")
	r.HandleFunc("/guests/{name}", guestHandler.UpdateAccompanyingGuests).Methods("PUT")
	r.HandleFunc("/guests/{name}", guestHandler.DeleteGuest).Methods("DELETE")
	r.HandleFunc("/guests", guestHandler.GetArrivedGuests).Methods("GET")
	r.HandleFunc("/seats_empty", tableHandler.GetEmptySeats).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}

func handlerPing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong\n")
}
