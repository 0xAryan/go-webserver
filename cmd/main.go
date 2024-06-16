package main

import (
	"fmt"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("static/"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			fmt.Fprint(w, "Hello world")
		} else {
			fmt.Fprintf(w, "Hello, %s\n", name)
		}

	})

	http.Handle("/cat/", http.StripPrefix("/cat/", fs))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Errorf("Cannot listen on port 8080 %s", err.Error())
	}

	var input string
	fmt.Println("Press Enter to exit...")
	fmt.Scanln(&input)
}
