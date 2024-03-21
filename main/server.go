package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var inputString string

type Prompt struct {
	Text string `json:"text"`
}

type GPT4Request struct {
	Model     string `json:"model"`
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
}

type GPT4Response struct {
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type ParaphraseResponse struct {
	Sent  string `json:"sent"`
	Paras []struct {
		Alt string `json:"alt"`
	} `json:"paras"`
}

func main() {
	http.HandleFunc("/api/prompt", gptHandler)

	// to read .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
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

func gptHandler(w http.ResponseWriter, r *http.Request) {
	//to see if gptHandler is even working
	log.Println("Received a request")
	// Parse the prompt from the request body
	var p Prompt
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	log.Println("Decoded request body:", p)

	prompt := p.Text
	log.Println("Prompt:", prompt)

	// Make a GPT-4 API call with the prompt
	inputString, err = callGPT4API(p.Text)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to call GPT-4 API: %s", err), http.StatusInternalServerError)
		return
	}

	// Paraphrase the GPT-4 API response
	paraphrasedString, err := paraphrase(inputString)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to paraphrase text: %s", err), http.StatusInternalServerError)
		return
	}

	// Send a response back to the client
	fmt.Fprintf(w, "Paraphrased text: %s", paraphrasedString)
}

// functions works
func callGPT4API(prompt string) (string, error) {
	log.Println("ABout to call GPT-4 API with prompt:", prompt)
	// Get the API key from .env file
	apiKey := os.Getenv("GPT4_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("GPT4_API_KEY is not set in .env file")
	}
	log.Println("Got API key from .env file")

	// Prepare the API request
	// gpt4Request := GPT4Request{
	// 	Model:     "text-davinci-004",
	// 	Prompt:    prompt,
	// 	MaxTokens: 2500, // Adjust this value as needed
	// }
	// Your API request setup and execution
	requestBody, err := json.Marshal(map[string]interface{}{
		"model":      "gpt-4",
		"messages":   []map[string]string{{"role": "user", "content": prompt}},
		"max_tokens": 2500,
	})
	if err != nil {
		return "", fmt.Errorf("Failed to marshal GPT-4 request: %s", err)
	}
	// log.Println("Marshalled GPT-4 request: ", bytes.NewBuffer(requestBody))
	log.Println("Request body:", string(requestBody))

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("Failed to create GPT-4 API request: %s", err)
	}
	log.Println("Created GPT-4 API request")
	// Set the request headers
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	//print request before it is sent
	log.Println("API Request: ", req)
	// Send the request
	log.Println("Sending API request...")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending API request:", err)
		return "", fmt.Errorf("Failed to send GPT-4 API request: %s", err)
	}
	log.Println("API request sent, response status code:", resp.StatusCode)
	// log.Println(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}
	log.Println("Received GPT-4 API response")

	// Parse the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to read GPT-4 API response: %s", err)
	}
	// log.Println("Body: ", body)
	defer resp.Body.Close()
	log.Println("Read GPT-4 API response")

	var gpt4Response GPT4Response
	err = json.Unmarshal(body, &gpt4Response)
	if err != nil {
		return "", fmt.Errorf("Failed to unmarshal GPT-4 API response: %s", err)
	}
	log.Println("Unmarshalled GPT-4 API response")
	// Extract the generated text
	// Extract the generated text
	generatedText := ""
	if len(gpt4Response.Choices) > 0 {
		generatedText = gpt4Response.Choices[0].Message.Content
	}
	//not printing generatedText
	log.Println("Extracted generated text:", generatedText)
	return generatedText, nil
}

// currently using paraphrase-genius API
// 5 request limit per day
// Either switch to a cheaper API or find a way to build your own paraphraser using Langchain or transoformers
// ***THIS API ONLY REDUCES AI DETECTION BY 50%MAX***
func paraphrase(input string) (string, error) {
	fmt.Println("Paraphrasing.....")

	// Prepare the API request
	apiURL := "https://paraphrase-genius.p.rapidapi.com/dev/paraphrase/"
	payload := map[string]string{
		"text":        input,
		"result_type": "single",
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("Failed to encode payload: %s", err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", apiURL, body)
	if err != nil {
		return "", fmt.Errorf("Failed to create paraphrase API request: %s", err)
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Rapidapi-Key", "402aebc953msh92a4014881dca61p1df56fjsn64b5430865c7")
	req.Header.Set("X-Rapidapi-Host", "paraphrase-genius.p.rapidapi.com")

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("Failed to send paraphrase API request: %s", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Paraphrase API request failed with status code: %d", resp.StatusCode)
	}

	// Parse the response
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to read paraphrase API response: %s", err)
	}

	var paraphraseResponse []string
	err = json.Unmarshal(bodyBytes, &paraphraseResponse)
	if err != nil {
		return "", fmt.Errorf("Failed to unmarshal paraphrase API response: %s", err)
	}

	// Extract the paraphrased text
	if len(paraphraseResponse) > 0 {
		return paraphraseResponse[0], nil
	}

	return "", nil
}
