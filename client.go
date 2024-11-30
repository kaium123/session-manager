package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", httpCall)
	fmt.Println("Server is listening on port 8085...")
	if err := http.ListenAndServe(":8085", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func httpCall(w http.ResponseWriter, r *http.Request) {
	// Call the other service
	resp, err := http.Get("http://localhost:8080/hi")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error making request: %s", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading response: %s", err), http.StatusInternalServerError)
		return
	}

	// Prepare the JSON response
	response := map[string]string{"message": string(body)}

	// Set the Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Encode and send the JSON response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %s", err), http.StatusInternalServerError)
	}
}
