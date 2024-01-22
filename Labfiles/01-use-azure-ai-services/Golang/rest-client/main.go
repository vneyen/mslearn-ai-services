package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	fmt.Println("Enter some text (\"quit\" to stop)")
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		text := reader.Text()
		if strings.ToLower(text) == "quit" {
			break
		}
		getLanguage(text)
		fmt.Println("Enter some text (\"quit\" to stop)")
	}
}

func getLanguage(text string) {
	ai_endpoint := os.Getenv("AI_SERVICE_ENDPOINT")
	ai_key := os.Getenv("AI_SERVICE_KEY")

	url := ai_endpoint + "/text/analytics/v3.1/languages?"

	jsonBody := []byte(`{
   "documents": [
       {"id": 1, "text": "` + text + `"}
    ]
}`)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Ocp-Apim-Subscription-Key", ai_key)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
