package sessionmanager

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hi" {
		fmt.Fprintf(w, "hello")
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
