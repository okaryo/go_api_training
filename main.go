package main

import (
	"github.com/gorilla/mux"
	"go_api_training/handlers"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler)
	r.HandleFunc("/article", handlers.PostArticleHandler)
	r.HandleFunc("/article/list", handlers.ArticleListHandler)
	r.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler)
	r.HandleFunc("/article/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
