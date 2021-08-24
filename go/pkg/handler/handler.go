package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	types "github.com/olivierdg2/react-go-docker-app/go/pkg/types/articles"
	"go.etcd.io/etcd/clientv3"
)

var cli clientv3.Client
var kv clientv3.KV

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	// Retrieve all articles
	articles, err := kv.Get(context.TODO(), "/articles", clientv3.WithPrefix())
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	var Articles types.Article
	for _, article := range articles.Kvs {
		var a types.Article
		json.Unmarshal(article.Value, &a)
		Articles = append(Articles, a)
	}
	json.NewEncoder(w).Encode(Articles)
}

func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
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

func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var data types.new_Article
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
		var last types.Article
		json.Unmarshal(last_data, &last)

		id, _ = strconv.Atoi(last.Id)
	}
	var new types.Article
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
	json.NewEncoder(w).Encode(string(article.Kvs[0].Value))
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]
	_, err := kv.Delete(context.TODO(), "/articles/"+id)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	// Retrieve all articles
	articles, err2 := kv.Get(context.TODO(), "/articles", clientv3.WithPrefix())
	if err2 != nil {
		fmt.Printf("Error: %v", err2)
	}
	var Articles []types.Article
	for _, article := range articles.Kvs {
		var a types.Article
		json.Unmarshal(article.Value, &a)
		Articles = append(Articles, a)
	}
	json.NewEncoder(w).Encode(Articles)

}

func PutArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, _ := ioutil.ReadAll(r.Body)
	var modified_article types.Article
	json.Unmarshal(reqBody, &modified_article)
	_, err := kv.Put(context.TODO(), "/articles/"+id, modified_article.toString())
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	json.NewEncoder(w).Encode(modified_article)
}
