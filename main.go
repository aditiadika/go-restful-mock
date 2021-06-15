package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var authors []Author = []Author{
	Author{
		Id:        "author-1",
		Firstname: "Jhon",
		Lastname:  "Doe",
		Username:  "JhonDoe",
		Password:  "password",
	},
	Author{
		Id:        "author-2",
		Firstname: "Budi",
		Lastname:  "Anduk",
		Username:  "BudiAnduk",
		Password:  "password",
	},
}

var articles []Article = []Article{
	Article{
		Id:      "article-1",
		Author:  "author-1",
		Title:   "Title One",
		Content: "lorem ipsum simply dolor si amet",
	},
	Article{
		Id:      "article-2",
		Author:  "author-2",
		Title:   "Title Two",
		Content: "lorem ipsum simply dolor si amet from article two",
	},
}

func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	response.Write([]byte(`{ "message" : "hello world"}`))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("Get")
	http.ListenAndServe(":12345", router)
}
