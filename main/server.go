package main

//go run main/server.go
// This import statement allows us to use the "fmt" package, which provides
// functions for formatted I/O operations. It is commonly used for printing
// output to the console and reading input from the user.
import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//prints hello world to the console
	//fmt.Println("Hello world")
	http.HandleFunc("/api/thumbnail", thumbnailHandler)

	// to read .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Serve static files from the frontend/dist directory.
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	//start server
	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

// Right now there is a a major security flaw. It is that the screenshot API token is visible in our frontend code in the console.
type thumbnailRequest struct {
	URL string `json:"url"`
}

type screenshotAPIRequest struct {
	Token          string `json:"token"`
	Url            string `json:"url"`
	Output         string `json:"output"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	ThumbnailWidth int    `json:"thumbnail_width"`
}

type screenshotAPIResponse struct {
	Screenshot string `json:"screenshot"`
}

// thumbnailHandler handles the HTTP request for generating a thumbnail.
// It decodes the request body, makes a screenshot API request, and returns the generated thumbnail.
func thumbnailHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a thumbnailRequest struct
	var decoded thumbnailRequest
	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a screenshot API request with the required parameters
	apiRequest := screenshotAPIRequest{
		Token:          os.Getenv("API_KEY"),
		Url:            decoded.URL,
		Output:         "json",
		Width:          1920,
		Height:         1080,
		ThumbnailWidth: 300,
	}
	fmt.Println("API key:", os.Getenv("API_KEY"))
	// Convert the API request to JSON
	jsonString, err := json.Marshal(apiRequest)
	checkError(err)

	// Create a new HTTP POST request to the screenshot API endpoint
	req, err := http.NewRequest("POST", "https://shot.screenshotapi.net/screenshot", bytes.NewBuffer(jsonString))
	checkError(err) // Fix: Check the error value

	// Set the Content-Type header to indicate JSON data
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request and get the response
	client := &http.Client{}
	response, err := client.Do(req)
	checkError(err)

	defer response.Body.Close()

	// Decode the response body into a screenshotAPIResponse struct
	var apiResponse screenshotAPIResponse
	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	checkError(err)

	// Write the generated thumbnail URL as a JSON response
	_, err = fmt.Fprintf(w, `{ "screenshot": "%s" }`, apiResponse.Screenshot)
	checkError(err)

	// Print the URL of the website to the console
	fmt.Println("Website URL:", decoded.URL)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
