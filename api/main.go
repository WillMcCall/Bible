package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Response struct {
	Word         string          `json:"word"`
	PageNumber   int             `json:"page"`
	TotalResults int             `json:"total_results"`
	TotalPages   int             `json:"total_pages"`
	Content      []WordReference `json:"results"`
}

type WordReference struct {
	Reference string `json:"reference"`
	Verse     string `json:"content"`
}

func GetWordJSON(word string, pageNum ...int) Response {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file")
	}

	apiKey := os.Getenv("ESV_KEY")

	page := 1 // Default value
	if len(pageNum) > 0 && pageNum[0] > 0 {
		page = pageNum[0]
	}

	url := "https://api.esv.org/v3/passage/search/?q=" + word + "&page-size=100&page=" + strconv.Itoa(page)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("request error: %v", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("response error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("request failed with status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}

	var response Response

	response.Word = word

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("error unmarshalling json: %v", err)
	}
	return response
}
