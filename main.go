package main

import (
	"log"
	"net/http"
	"database/sql"
	
     "fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	 _ "github.com/lib/pq"
	"github.com/mandloiabhi/PG_3/internal/database" 
	
	
)
const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    dbname   = "blogator"
)
type apiConfig struct {
	DB *database.Queries
}
func main(){
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
         
        // open database
	
	//var db *sql.DB

	// opens connection to a database

    db, err := sql.Open("postgres", psqlconn)
	if err != nil{
        log.Printf("Error connecting: %s", err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
		}
	dbQueries := database.New(db)	
	apiCfg := apiConfig{
		DB: dbQueries,
	}

    defer db.Close()

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	
	srv := &http.Server{
		Addr:    ":" + "8080",
		Handler: router,
	}

	
	router.Mount("/v1", v1Router)
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/users", apiCfg.handlerUsersCreate)

	mux.HandleFunc("GET /v1/healthz", handlerReadiness)
	
	log.Fatal(srv.ListenAndServe())



}