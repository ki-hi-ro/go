package main

import (
	"fmt"
	"net/http"
	"os"
)

var items []string

func main() {
	http.HandleFunc("/", showList)
	http.HandleFunc("/add", addItem)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}	

	fmt.Println("http://localhost:8080 で起動中")
	http.ListenAndServe(":8080", nil)
}

func showList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
		<!DOCTYPE html>
		<html lang="ja">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>買い物リスト</title>
			<style>
				body {
					font-family: sans-serif;
					background: #f5f5f5;
					margin: 0;
					padding: 24px;
				}

				.container {
					max-width: 480px;
					margin: 0 auto;
					background: white;
					padding: 24px;
					border-radius: 16px;
					box-shadow: 0 4px 16px rgba(0,0,0,0.08);
				}

				h1 {
					font-size: 24px;
					margin-bottom: 20px;
				}

				form {
					display: flex;
					gap: 8px;
				}

				input {
					flex: 1;
					padding: 12px;
					font-size: 16px;
				}

				button {
					padding: 12px 16px;
					font-size: 16px;
				}

				ul {
					padding-left: 20px;
				}

				li {
					font-size: 18px;
					margin: 10px 0;
				}

				@media (max-width: 600px) {
					body {
						padding: 12px;
					}

					.container {
						padding: 20px;
					}

					form {
						flex-direction: column;
					}

					button {
						width: 100%;
					}
				}
			</style>
		</head>
		<body>
			<div class="container">
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
			</div>
		</body>
		</html>
	`)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	if item != "" {
		items = append(items, item)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}