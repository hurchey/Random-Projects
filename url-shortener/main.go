package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB

const (
	bucketName  = "urls"
	codeLength  = 6
	listenAddr  = ":8080"
	baseURL     = "http://localhost:8080"
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// shortenRequest represents the JSON payload for a URL shortening request.
type shortenRequest struct {
	URL string `json:"url"`
}

// shortenResponse represents the JSON response after shortening a URL.
type shortenResponse struct {
	ShortURL string `json:"short_url"`
}

func main() {
	// Open or create the BoltDB database.
	var err error
	db, err = bolt.Open("urls.db", 0600, nil)
	if err != nil {
		log.Fatal("Error opening BoltDB:", err)
	}
	defer db.Close()

	// Ensure the bucket exists.
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	})
	if err != nil {
		log.Fatal("Error creating bucket:", err)
	}

	// Initialize the router.
	router := mux.NewRouter()

	// Root endpoint for a simple welcome message.
	router.HandleFunc("/", rootHandler).Methods("GET")

	// POST /shorten for creating a new shortened URL.
	router.HandleFunc("/shorten", shortenHandler).Methods("POST")

	// GET /{shortCode} for redirecting to the original URL.
	router.HandleFunc("/{shortCode}", redirectHandler).Methods("GET")

	fmt.Printf("Server running on %s\n", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, router))
}

// rootHandler responds with a welcome message.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the URL Shortener API. Use POST /shorten to create a shortened URL."))
}

// shortenHandler processes URL shortening requests.
func shortenHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("shortenHandler called")
	var req shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	if req.URL == "" {
		http.Error(w, "URL field is required", http.StatusBadRequest)
		return
	}

	// Generate a unique short code.
	shortCode := generateShortCode(codeLength)

	// Store the mapping in BoltDB.
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		return bucket.Put([]byte(shortCode), []byte(req.URL))
	})
	if err != nil {
		http.Error(w, "Error storing URL", http.StatusInternalServerError)
		return
	}

	// Create and send the JSON response.
	resp := shortenResponse{ShortURL: fmt.Sprintf("%s/%s", baseURL, shortCode)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// redirectHandler handles redirection using the short code.
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]
	log.Printf("redirectHandler called with shortCode: %s", shortCode)

	var originalURL string
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		value := bucket.Get([]byte(shortCode))
		if value == nil {
			return fmt.Errorf("URL not found")
		}
		originalURL = string(value)
		return nil
	})
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	// Redirect the user to the original URL.
	http.Redirect(w, r, originalURL, http.StatusFound)
}

// generateShortCode creates a random string of the specified length.
func generateShortCode(n int) string {
	// Seed the random number generator. In production, consider a more robust solution.
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
