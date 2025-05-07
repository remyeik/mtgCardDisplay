# Scryfall Card Search CLI

This is a simple CLI tool for looking up Magic: The Gathering cards using the [Scryfall API](https://scryfall.com/docs/api).

## Features

- Search for cards by a query string (fuzzy search).
- Display a list of matching cards with basic details such as the card name, set, and more.
- **Select a card** from the search results.
- **Display full card details**, including mana cost, card type, oracle text, set, image URL, and Scryfall URL.

## Usage

1. Run the CLI tool:

   After you have the code set up, run the application with a search query:

   ```bash
   go run main.go "lightning bolt"
   ```

   This will search for Magic: The Gathering cards matching "lightning bolt" and display a list of matching results.

2. Choose a card:

   The application will show a numbered list of matching cards. You will be prompted to enter the number corresponding to the card you want to see more details about.

   Example output:

   ```
   Found the following cards:
   1. Lightning Bolt - Alpha
   2. Lightning Bolt - Beta
   3. Lightning Bolt - Unlimited

   Enter the number of the card you want to view: 1

   🃏 Name:       Lightning Bolt
   💧 Mana Cost:  {R}
   📜 Type:       Instant
   📖 Oracle Text:Lightning Bolt deals 3 damage to any target.
   📦 Set:        Alpha
   🖼️ Image:      https://cards.scryfall.io/normal/front/4/0/40c01ff9-...
   🔗 URL:        https://scryfall.com/card/lea/163/lightning-bolt
   ```

## How It Works

- The application uses the Scryfall API to search for cards using a query string.
- Once the search is complete, it displays the cards that match your query.
- After selecting a card, it fetches detailed information about the selected card and displays it in the terminal.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
