package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Card struct {
	Name       string `json:"name"`
	ManaCost   string `json:"mana_cost"`
	TypeLine   string `json:"type_line"`
	OracleText string `json:"oracle_text"`
	SetName    string `json:"set_name"`
	ImageUris  struct {
		Normal string `json:"normal"`
	} `json:"image_uris"`
	ScryfallURI string            `json:"scryfall_uri"`
	Legalities  map[string]string `json:"legalities"`
}

type SearchResponse struct {
	Data []Card `json:"data"`
}

func searchCard(query string) ([]Card, error) {
	encodedQuery := url.QueryEscape(query)
	apiURL := fmt.Sprintf("https://api.scryfall.com/cards/search?q=%s", encodedQuery)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result SearchResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// displayCardDetails function  
func displayCardDetails(card Card) {
	fmt.Printf("🃏 Name:       %s\n", card.Name)
	fmt.Printf("🌄 Mana Cost:  %s\n", card.ManaCost)
	fmt.Printf("📜 Type:       %s\n", card.TypeLine)
	fmt.Printf("📖 Oracle Text: %s\n", card.OracleText)
	fmt.Printf("📦 Set:        %s\n", card.SetName)
	fmt.Printf("🖼️ Image:      %s\n", card.ImageUris.Normal)
	fmt.Printf("🔗 URL:        %s\n", card.ScryfallURI)

	// Check legality for Commander format
	if commanderLegality, ok := card.Legalities["commander"]; ok && commanderLegality != "legal" {
		fmt.Printf("⚖️ Commander Legal: %s\n", commanderLegality)
	}

	// Check legality for Standard format
	if standardLegality, ok := card.Legalities["standard"]; ok && standardLegality != "legal" {
		fmt.Printf("⚖️ Standard Legal: %s\n", standardLegality)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: Search for a card by name or keyword. Example: mtg lightning bolt")
		return
	}

	query := strings.Join(os.Args[1:], " ")
	cards, err := searchCard(query)
	if err != nil {
		fmt.Println("Error searching cards:", err)
		return
	}

	if len(cards) == 0 {
		fmt.Println("No cards found.")
		return
	}

	// If exactly one card is found, display it immediately
	if len(cards) == 1 {
		displayCardDetails(cards[0])
		return
	}

	// Otherwise, show the list of matching cards and let the user choose
	fmt.Println("Found the following cards:")
	for i, card := range cards {
		fmt.Printf("%d. %s - %s\n", i+1, card.Name, card.SetName)
	}

	var choice int
	fmt.Print("Enter the number of the card you want to view: ")
	_, err = fmt.Scan(&choice)
	if err != nil || choice < 1 || choice > len(cards) {
		fmt.Println("Invalid choice.")
		return
	}

	displayCardDetails(cards[choice-1])
}
