package server

import (
	"https://github.com/olivierdg2/react-go-docker-app/go/pkg/cors/cors"
	"https://github.com/olivierdg2/react-go-docker-app/go/pkg/handler/handler"
	"https://github.com/olivierdg2/react-go-docker-app/go/pkg/types/types"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(CORS)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE", "OPTIONS")
	myRouter.HandleFunc("/article/{id}", putArticle).Methods("PUT", "OPTIONS")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
