package main

import (
	"fmt"
	"net/http"
)

var items []string

func main() {
	http.HandleFunc("/", showList)
	http.HandleFunc("/add", addItem)

	fmt.Println("http://localhost:8080 で起動中")
	http.ListenAndServe(":8080", nil)
}

func showList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
		<h1>買い物リスト</h1>
		<form action="/add" method="POST">
			<input name="item" placeholder="商品名を入力">
			<button type="submit">追加</button>
		</form>
		<ul>
	`)

	for _, item := range items {
		fmt.Fprintf(w, "<li>%s</li>", item)
	}

	fmt.Fprint(w, `
		</ul>
	`)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	if item != "" {
		items = append(items, item)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}