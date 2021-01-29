package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article is
type Article struct {
	ArticleID          string `json:"ArticleID"`
	ArticleTitle       string `json:"ArticleTitle"`
	ArticleDescription string `json:"ArticleDescription"`
	ArticleText        string `json:"ArticleText"`
	ArticlePostDate    string `json:"ArticlePostDate"`
}

var articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint called:homePage()")
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	_deleteArticleAtArticleid(params["articleid"])
	json.NewEncoder(w).Encode(articles)

}

func _deleteArticleAtArticleid(articleid string) {
	for index, article := range articles {
		if article.ArticleID == articleid {
			// delete item from slice
			articles = append(articles[:index], articles[index+1:]...)
			break
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	params := mux.Vars(r)
	_deleteArticleAtArticleid(params["articleid"])
	articles = append(articles, article)
	json.NewEncoder(w).Encode(articles)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Function Called: getArticle()")
	json.NewEncoder(w).Encode(articles)
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	articles = append(articles, article)
	json.NewEncoder(w).Encode(article)
}

func handlerequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/article", getArticle).Methods("GET")
	router.HandleFunc("/article", createArticle).Methods("POST")
	router.HandleFunc("/article/{articleid}", deleteArticle).Methods("DELETE")
	router.HandleFunc("/article/{articleid}", updateArticle).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func main() {
	articles = append(articles, Article{
		ArticleID:          "0",
		ArticleTitle:       "traveling and exposure",
		ArticleDescription: "Traveling makes person mind relax and comfortable.",
		ArticleText:        "Traveling makes person mind relax and comfortable.It also makes person more interactive and rejuvinates person mind.",
		ArticlePostDate:    "04/05/2020",
	})

	articles = append(articles, Article{
		ArticleID:          "1",
		ArticleTitle:       "Cricket",
		ArticleDescription: "Cricket is a game played with bat and bowl.",
		ArticleText:        "Cricket is a game played with bat and bowl.This game is played between two teams.It includes 11 players in each teams. ",
		ArticlePostDate:    "04/05/2020",
	})

	handlerequests()
}
