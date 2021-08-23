package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/olivierdg2/react-go-docker-app/go/pkg/cors/cors"
	"github.com/olivierdg2/react-go-docker-app/go/pkg/handler/handler"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(cors.CORS)
	myRouter.HandleFunc("/", handler.HomePage)
	myRouter.HandleFunc("/articles", handler.ReturnAllArticles)
	myRouter.HandleFunc("/article", handler.CreateNewArticle).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/article/{id}", handler.DeleteArticle).Methods("DELETE", "OPTIONS")
	myRouter.HandleFunc("/article/{id}", handler.PutArticle).Methods("PUT", "OPTIONS")
	myRouter.HandleFunc("/article/{id}", handler.ReturnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
