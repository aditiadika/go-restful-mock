package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	data, _ := json.Marshal(authors)
	fmt.Println(string(data))
}
