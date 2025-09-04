package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting HTTP proxy server on :8888")

	server := &http.Server{
		Addr:    ":8888",
		Handler: http.HandlerFunc(proxyHandler),
	}

	log.Fatal(server.ListenAndServe())
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s %s", r.Method, r.Host, r.URL.Path)

	// We need the full URL for the outbound request
	if r.URL.Scheme == "" {
		r.URL.Scheme = "https" // Assume HTTPS if not specified
	}
	if r.URL.Host == "" {
		r.URL.Host = r.Host
	}

	// Create a new request to send to the destination server
	outReq, err := http.NewRequest(r.Method, r.URL.String(), r.Body)
	if err != nil {
		log.Printf("Error creating outbound request: %v", err)
		http.Error(w, "Error creating outbound request", http.StatusInternalServerError)
		return
	}

	// Copy headers from the client request to the outbound request
	for key, values := range r.Header {
		for _, value := range values {
			outReq.Header.Add(key, value)
		}
	}
	// Remove headers that are specific to the proxy connection
	outReq.Header.Del("Proxy-Connection")
	outReq.Header.Del("Proxy-Authorization")

	// Send the request to the destination
	client := &http.Client{}
	resp, err := client.Do(outReq)
	if err != nil {
		log.Printf("Error sending outbound request: %v", err)
		http.Error(w, "Error sending outbound request", http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	log.Printf("Received response %s for %s", resp.Status, r.URL.String())

	// Copy headers from the destination's response to our response for the client
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Write the status code and copy the body
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
