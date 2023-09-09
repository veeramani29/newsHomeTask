package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Article struct {
	Title    string `json:"title"`
	Description  string `json:"description"`
	Content string `json:"content"`
	Url   string    `json:"url"`
	Publishedat   string    `json:"publishedAt"`
	Image   string    `json:"image"`
	Id   string    `json:"image"`
}

type Response struct {
	StatusCode int    `json:"statuscode"`
	Message    string `json:"message"`
}

var articles []Article

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveServer).Methods("get")
	r.HandleFunc("/news", NewsList).Methods("get")
	r.HandleFunc("/news/{id}", DeleteNewsList).Methods("delete")
	r.HandleFunc("/news", AddNewsList).Methods("post")
	r.HandleFunc("/news/{id}", GetNewsList).Methods("get")
	r.HandleFunc("/news/{id}", UpdateNewsList).Methods("put")
	log.Fatal((http.ListenAndServe(":1000", r)))

}

func serveServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Home Assigmane</>"))
}

func DeleteNewsList(w http.ResponseWriter, r *http.Request) {
	rsp := Response{200, "done"}
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, v := range articles {
		if v.Id == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(rsp)
}
func AddNewsList(w http.ResponseWriter, r *http.Request) {
	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	article.Id = strconv.Itoa(rand.Intn(100))
	articles = append(articles, article)
	rsp := Response{200, "done"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rsp)
}
func UpdateNewsList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)
	for index, v := range articles {
		if v.Id == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)
		}
	}
	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	article.Id = params["id"]
	articles = append(articles, article)
	rsp := Response{200, "done"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rsp)
}
func GetNewsList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params)
	for _, v := range articles {
		if v.Id == params["id"] {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode("NNot valiud id")
	return

}
func NewsList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}
