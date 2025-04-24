
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
	Name       string `json:"name"`
	ManaCost   string `json:"mana_cost"`
	TypeLine   string `json:"type_line"`
	OracleText string `json:"oracle_text"`
	SetName    string `json:"set_name"`
	ScryfallURI string `json:"scryfall_uri"`
	ImageUris  struct {
		Normal string `json:"normal"`
	} `json:"image_uris"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <card name>")
		return
	}

	cardName := strings.Join(os.Args[1:], " ")
	apiURL := "https://api.scryfall.com/cards/named?fuzzy=" + strings.ReplaceAll(cardName, " ", "+")

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error fetching card:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("Card not found or error (%d)\n", resp.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var card Card
	err = json.Unmarshal(body, &card)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

fmt.Printf("🃏 Name:       %s\n", card.Name)
fmt.Printf("💧 Mana Cost:  %s\n", card.ManaCost)
fmt.Printf("📜 Type:       %s\n", card.TypeLine)
fmt.Printf("📖 Oracle Text:%s\n", card.OracleText)
fmt.Printf("📦 Set:        %s\n", card.SetName)
fmt.Printf("🖼️ Image:      %s\n", card.ImageUris.Normal)
fmt.Printf("🔗 URL:        %s\n", card.ScryfallURI)
}
