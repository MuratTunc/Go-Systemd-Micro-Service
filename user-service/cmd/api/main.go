package main

import (
	"fmt"
	"log"
	"net/http"
)

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow all origins (adjust as needed)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Allow specific headers and methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// Health check handler
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("✅ /healthcheck called — sending 'success' to client")
	fmt.Fprintln(w, "success")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", healthCheckHandler)

	// Wrap with CORS middleware
	handler := corsMiddleware(mux)

	port := "8080"
	log.Printf("🚀 Server is running on port %s...", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
