package main

import (
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`byte
			<html>
				<head>
					<meta charset="utf-8" />
					<title>Simple messenger</title>
				</head>
				<body>
					<h1>Simple messenger</h1>
				</body>
			</html>
		`))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}