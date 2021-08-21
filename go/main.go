package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.etcd.io/etcd/clientv3"
)

var cli clientv3.Client
var kv clientv3.KV

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

type new_Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

func (a Article) toString() string {
	var s string
	s = "{\"Id\":\"" + a.Id + "\",\"Title\":\"" + a.Title + "\",\"Desc\":\"" + a.Desc + "\",\"Content\":\"" + a.Content + "\"}"
	return s
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", putArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	// Retrieve all articles
	articles, err := kv.Get(context.TODO(), "/articles", clientv3.WithPrefix())
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	var Articles []Article
	for _, article := range articles.Kvs {
		var a Article
		json.Unmarshal(article.Value, &a)
		Articles = append(Articles, a)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	//Get url variables then get the id
	vars := mux.Vars(r)
	key := vars["id"]
	//Retrieve the article from the given id
	article, err := kv.Get(context.TODO(), "/articles/"+key)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	json.NewEncoder(w).Encode(string(article.Kvs[0].Value))
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var data new_Article
	json.Unmarshal(reqBody, &data)
	articles, err := kv.Get(context.TODO(), "/articles", clientv3.WithPrefix())
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	var id int
	if len(articles.Kvs) == 0 {
		id = 0
	} else {
		last_data := articles.Kvs[len(articles.Kvs)-1].Value
		var last Article
		json.Unmarshal(last_data, &last)

		id, _ = strconv.Atoi(last.Id)
	}
	var new Article
	new.Id = strconv.Itoa(id + 1)
	new.Title = data.Title
	new.Desc = data.Desc
	new.Content = data.Content
	_, err2 := kv.Put(context.TODO(), "/articles/"+new.Id, new.toString())
	if err2 != nil {
		fmt.Printf("Error: %v", err)
	}
	article, err3 := kv.Get(context.TODO(), "/articles/"+new.Id)
	if err3 != nil {
		fmt.Printf("Error: %v", err2)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(string(article.Kvs[0].Value))
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]
	_, err := kv.Delete(context.TODO(), "/articles/"+id)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Fprintf(w, "Article deleted")

}

func putArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, _ := ioutil.ReadAll(r.Body)
	var modified_article Article
	json.Unmarshal(reqBody, &modified_article)
	_, err := kv.Put(context.TODO(), "/articles/"+id, modified_article.toString())
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	json.NewEncoder(w).Encode(modified_article)
}

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	kv = clientv3.NewKV(cli)
	if err != nil {
		fmt.Printf("%v", err)
	}
	handleRequests()
	defer cli.Close()
}
