package main

import (
	"fmt"
	"html/template"
	"net/http"
	"portfolio/database"
)

type Players struct {
	name     string
	team     string
	position string
	country  string
	age      int
	score    int
	assist   int
	foot     string
	height   int
}

var SearchResult []Players

/*
検索画面遷移時の処理
*/
func handleSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("./assets/html/search.html")
		if err != nil {
			panic(err.Error())
		}
		if err := t.Execute(w, nil); err != nil {
			panic(err.Error())
		}
	}
	if r.Method == "POST" {
		db := database.Connect()
		defer db.Close()

		rows, err := db.Query(database.CreateQuery(r))

		if err != nil {
			fmt.Println("データベース接続失敗")
			panic(err.Error())
		} else {
			fmt.Println("データベース接続成功")
		}

		defer rows.Close()

		SearchResult = []Players{}

		for rows.Next() {
			var players Players
			err := rows.Scan(
				&players.name,
				&players.team,
				&players.position,
				&players.country,
				&players.age,
				&players.score,
				&players.assist,
				&players.foot,
				&players.height,
			)

			if err != nil {
				panic(err.Error())
			}

			SearchResult = append(SearchResult, players)

		}

		err = rows.Err()

		if err != nil {
			panic(err.Error())
		}

		http.Redirect(w, r, "/searchResult/", 301)
	}
}
