package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
	"github.com/google/uuid"
)


type ProxyRequest struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}


type ProxyResponse struct {
	ID      string            `json:"id"`
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Length  int               `json:"length"`
}

var store sync.Map 

func handler(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()


	var req ProxyRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}


	if req.URL == "" || req.Method == "" {
		http.Error(w, "Missing 'method' or 'url'", http.StatusBadRequest)
		return
	}


	client := &http.Client{Timeout: 10 * time.Second}
	proxyReq, err := http.NewRequest(req.Method, req.URL, nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}


	for k, v := range req.Headers {
		proxyReq.Header.Add(k, v)
	}


	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("Request error: %v", err), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()


	respBody, _ := io.ReadAll(resp.Body)


	id := uuid.New().String()
	headers := make(map[string]string)
	for k, v := range resp.Header {
		headers[k] = v[0]
	}
	proxyResp := ProxyResponse{
		ID:      id,
		Status:  resp.StatusCode,
		Headers: headers,
		Length:  len(respBody),
	}


	store.Store(id, map[string]interface{}{
		"request":  req,
		"response": proxyResp,
	})


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(proxyResp)
}

func main() {
	http.HandleFunc("/proxy", handler)
	fmt.Println("🚀 Proxy server started on http://localhost:8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" 
	}
	http.ListenAndServe(":"+port, nil)

}
