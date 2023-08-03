package main

import (
	"encoding/json"
	"fmt"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
	"log"
	"net/http"
	"os"
	"strconv"
)

type QueryResult struct {
	WhoisInfo whoisparser.WhoisInfo `json:"whois_info"`
	WhoisRaw  string                `json:"whois_raw"`
	QueryItem string                `json:"item"`
}

func query(queryItem string) (string, whoisparser.WhoisInfo, error) {
	result, err := whois.Whois(queryItem)
	if err != nil {
		panic(err)
	}
	whoisParsed, err := whoisparser.Parse(result)
	if err != nil {
		panic(err)
	}
	// print in json
	return result, whoisParsed, err
}
func queryHandler(w http.ResponseWriter, r *http.Request) {
	// Get the query parameter "query_item" from the request URL
	queryItem := r.URL.Query().Get("item")

	// Call the query function with the provided query_item
	result, whoisParsed, err := query(queryItem)
	if err != nil {
		// return a 500 Internal Server Response
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return json of QueryResult
		w.Header().Set("Content-Type", "application/json")
		queryResult := QueryResult{whoisParsed, result, queryItem}
		json.NewEncoder(w).Encode(queryResult)
	}

}

// run as a http api
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run app.go <port>")
		return
	}
	portStr := os.Args[1]

	// Convert portStr to an integer
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/whois", queryHandler)

	// Start the HTTP server
	log.Printf("Server started on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
