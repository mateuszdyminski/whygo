package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// list of backend workers
var backendServers = map[int]string{
	0: "http://dummy1:9001?q=",
	1: "http://dummy2:9002?q=",
	2: "http://dummy3:9003?q=",
}

func main() {
	// configure http paths
	http.HandleFunc("/search", searchForDocker)

	// start http server
	fmt.Println("Listening on port 9000")
	http.ListenAndServe(":9000", nil)
}

// endpoint which will be triggered when user create request for /search path
func searchForDocker(w http.ResponseWriter, r *http.Request) {
	// get search term from url query string
	query := r.URL.Query().Get("q")
	fmt.Printf("Start searching for string=%s\n", query)

	// create a channel for responses
	resultChan := make(chan response)

	// run queries to all backend services
	for k := range backendServers { // HL
		go func() { // HL
			resultChan <- searchInBackendDocker(k, query) // HL
		}() // HL
	} // HL

	// get only first result and write answer to response
	res := <-resultChan // HL
	fmt.Printf("Got response from server with id: %d\n", res.serverID)
	fmt.Fprint(w, string(res.payload)) // HL
}

// make a query to a backend service
func searchInBackendDocker(id int, query string) response {
	// make GET request
	resp, err := http.Get(backendServers[id] + query) // HL
	if err != nil {
		log.Fatalf("can't make request to: %s, err: %v", backendServers[id]+query, err)
	}

	if resp.Status != "200 OK" {
		log.Fatalf("wrong http status: %s", resp.Status)
	}

	// read response and return bytes
	bytes, _ := ioutil.ReadAll(resp.Body) // HL
	return response{id, bytes}
}

type response struct {
	serverID int
	payload  []byte
}
