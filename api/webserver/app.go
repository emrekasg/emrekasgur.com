package webserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func RunApp() {
	app := new(App)
	app.Run()
}

func (app *App) Run() {
	staticPath := "/"
	r := mux.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "https://emrekasgur.com")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	r.HandleFunc(staticPath, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	}).Methods("GET")

	r.HandleFunc("/healthcheck", HealthCheck).Methods("GET")

	r.HandleFunc("/posts", GetPosts).Methods("GET")
	r.HandleFunc("/posts/{postLink}", GetPost).Methods("GET")
	r.HandleFunc("/tags", GetTags).Methods("GET")

	port := fmt.Sprintf(":%d", 80)

	if err := http.ListenAndServe(port, r); err != nil {
		fmt.Println(err)
	}
}
