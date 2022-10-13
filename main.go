package main

import (
	"log"
	"net/http"
	"go_api_training/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/article/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
