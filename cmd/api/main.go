package main

import (
	"fmt"
	"net/http"

	"github.com/Luke-Frazer/Personal-Projects/tree/main/Go_Basic_API/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus" // aliased as "log"
)

func main() {
	// turns on logger
	log.SetReportCaller(true)

	// pointer to a mux type, basically a struct used to setup api
	var r *chi.Mux = chi.NewRouter()

	// pass into handler function
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")
	fmt.Println(`
 ______     ______        ______     ______   __    
/\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \   
\ \ \__/\  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \  
 \ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\ 
  \/_____/   \/_____/      \/_/\/_/   \/_/     \/_/ `)

	// capture any error and handle them
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
