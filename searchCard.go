package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Card struct {
	Name        string `json:"name"`
	ManaCost    string `json:"mana_cost"`
	TypeLine    string `json:"type_line"`
	OracleText  string `json:"oracle_text"`
	SetName     string `json:"set_name"`
	ScryfallURI string `json:"scryfall_uri"`
	ImageUris   struct {
		Normal string `json:"normal"`
	} `json:"image_uris"`
}

type SearchResponse struct {
	Data []Card `json:"data"`
}

func main() {
	// Ensure the user has input a search string
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <search query>")
		return
	}

	// Build the search query from the arguments
	query := strings.Join(os.Args[1:], " ")
	apiURL := "https://api.scryfall.com/cards/search?q=" + strings.ReplaceAll(query, " ", "+")

	// Fetch data from Scryfall API
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	// Check if the response is valid
	if resp.StatusCode != 200 {
		fmt.Printf("Error: Unable to fetch data. Status code: %d\n", resp.StatusCode)
		return
	}

	// Read and parse the JSON response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Parse the JSON data into the struct
	var searchResp SearchResponse
	err = json.Unmarshal(body, &searchResp)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Check if there are any results
	if len(searchResp.Data) == 0 {
		fmt.Println("No cards found for your search.")
		return
	}

	// Display search results
	fmt.Println("\nFound the following cards:")
	for i, card := range searchResp.Data {
		fmt.Printf("%d. %s - %s\n", i+1, card.Name, card.SetName)
	}

	// Prompt user to choose a card
	fmt.Print("\nEnter the number of the card you want to view: ")
	var choice int
	_, err = fmt.Scanf("%d", &choice)
	if err != nil || choice < 1 || choice > len(searchResp.Data) {
		fmt.Println("Invalid choice. Exiting.")
		return
	}

	// Display the full information of the selected card
	selectedCard := searchResp.Data[choice-1]
	fmt.Printf("\n🃏 Name:       %s\n", selectedCard.Name)
	fmt.Printf("💧 Mana Cost:  %s\n", selectedCard.ManaCost)
	fmt.Printf("📜 Type:       %s\n", selectedCard.TypeLine)
	fmt.Printf("📖 Oracle Text:%s\n", selectedCard.OracleText)
	fmt.Printf("📦 Set:        %s\n", selectedCard.SetName)
	fmt.Printf("🖼️ Image:      %s\n", selectedCard.ImageUris.Normal)
	fmt.Printf("🔗 URL:        %s\n", selectedCard.ScryfallURI)
}
