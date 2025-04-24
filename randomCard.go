package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func main() {
	resp, err := http.Get("https://api.scryfall.com/cards/random")
	if err != nil {
		fmt.Println("Error fetching:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	var card Card
	err = json.Unmarshal(body, &card)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
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
