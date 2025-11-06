package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const ollamaURL = "http://localhost:11434/api/generate"

// Request payload for Ollama
type GenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

// Response structure (for non-streamed responses)
type GenerateResponse struct {
	Model    string `json:"model"`
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

func main() {
	// Check if user passed a prompt argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go \"<your prompt>\"")
		return
	}

	prompt := os.Args[1]

	// Prepare JSON payload
	requestBody := GenerateRequest{
		Model:  "llama3", // you can change this to any model you have locally
		Prompt: prompt,
		Stream: false, // set true if you want streamed responses
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Send POST request to Ollama API
	resp, err := http.Post(ollamaURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var result GenerateResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		fmt.Println("Error parsing response:", err)
		fmt.Println("Raw response:", string(respBody))
		return
	}

	fmt.Println("\n🦙 Ollama response:\n")
	fmt.Println(result.Response)
}
