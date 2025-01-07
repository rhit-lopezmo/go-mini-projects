package main

import (
	"net/http"
)

func main() {
	// serve files from static dir
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// start server on port 3000
	port := ":3000"
	println("Server is running at http://localhost" + port)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		panic(err)
	}
}
