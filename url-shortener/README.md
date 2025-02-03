# URL Shortener

A basic URL shortener web service built in Go. This project demonstrates how to create a REST API using Gorilla Mux for routing and BoltDB (bbolt) for storage.

## Getting Started

### Initialize the Go Module

If not already initialized, run:

```bash
go mod init github.com/yourusername/url-shortener
```

### Install Dependencies

Install Gorilla Mux and BoltDB (bbolt) with:

```bash
go get -u github.com/gorilla/mux
go get go.etcd.io/bbolt
```

## Running the Application

### Start the Server

From your project directory, run:

```bash
go run main.go
```

The server will start and listen on http://localhost:8080.

## Testing the Endpoints

### Root Endpoint

* **URL:** `http://localhost:8080/`
* **Method:** GET

**Example:**
```bash
curl http://localhost:8080/
```

**Response:** A welcome message with instructions.

### Create a Short URL

* **URL:** `http://localhost:8080/shorten`
* **Method:** POST

**Example using curl:**
```bash
curl -X POST -d '{"url": "https://www.example.com"}' -H "Content-Type: application/json" http://localhost:8080/shorten
```

**Response:** A JSON object containing the shortened URL.

### Redirect to the Original URL

* **URL:** `http://localhost:8080/{shortCode}`
* **Method:** GET

**Example:** Visit the returned short URL in your browser to be redirected to the original URL.

## Code Overview

* **main.go:** Contains the full source code for:
  * Setting up the HTTP server
  * Defining the API endpoints
  * Handling URL shortening
  * Redirection
  * Interacting with BoltDB for storage

## Future Improvements

* **Collision Handling:** Check for and handle potential collisions when generating short codes
* **URL Validation:** Ensure that submitted URLs are valid
* **Custom Short Codes:** Allow users to specify custom short codes
* **Error Handling:** Enhance error responses and logging for production use

## License

This project is licensed under the MIT License.

---

### How to Use This README

1. **Create a New File:** In your project directory, create a file named `README.md`
2. **Paste the Content:** Copy the entire content above and paste it into your `README.md` file
3. **Customize as Needed:** Replace placeholders like `github.com/yourusername/url-shortener` with your actual repository URL or project details
4. **Save and Commit:** Save the file and commit it to your repository