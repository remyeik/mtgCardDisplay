# Scryfall Card Search CLI

This is a simple CLI tool for looking up Magic: The Gathering cards using the [Scryfall API](https://scryfall.com/docs/api).

## Features

- Search for cards by a query string (fuzzy search).
- Display a list of matching cards with basic details such as the card name, set, and more.
- **Select a card** from the search results.
- **Display full card details**, including mana cost, card type, oracle text, set, image URL, and Scryfall URL.

## Installation

1. **Clone the repository** (or download the source code):

   ```bash
   git clone https://github.com/remyeik/mtgCardDisplay.git
   cd scryfall-cli
   ```

2. **Build and install the application**:

   ```bash
   go build -o mtg searchCard.go
   sudo mv mtg /usr/local/bin/
   sudo chmod +x /usr/local/bin/mtg
   ```

3. **Verify installation**:

   ```bash
   which mtg
   # Should output: /usr/local/bin/mtg
   ```

## Usage

Once installed, you can use the `mtg` command from anywhere in your terminal:

```bash
mtg "lightning bolt"
```

This will search for Magic: The Gathering cards matching "lightning bolt" and display a list of matching results.

### Choose a card

The application will show a numbered list of matching cards. You will be prompted to enter the number corresponding to the card you want to see more details about.

**Example output:**

```
Found the following cards:
1. Solemn Offering - Battlebond
2. Sol Ring - Tarkir: Dragonstorm Commander
Enter the number of the card you want to view: 2
🃏 Name:       Sol Ring
🌄 Mana Cost:  {1}
📜 Type:       Artifact
📖 Oracle Text: {T}: Add {C}{C}.
📦 Set:        Tarkir: Dragonstorm Commander
🖼️ Image:      https://cards.scryfall.io/normal/front/7/b/7b0ccfd4-f118-4497-9c5a-54cf12000e16.jpg?1743206094
🔗 URL:        https://scryfall.com/card/tdc/106/sol-ring?utm_source=api
⚖️ Standard Legal: not_legal
```

### More Examples

```bash
mtg "counterspell"
mtg "black lotus"
mtg "lightning bolt"
```

## Uninstalling

To remove the application:

```bash
sudo rm /usr/local/bin/mtg
```

## How It Works

- The application uses the Scryfall API to search for cards using a query string.
- Once the search is complete, it displays the cards that match your query.
- After selecting a card, it fetches detailed information about the selected card and displays it in the terminal.

## Requirements

- Go 1.19 or later
- Internet connection (for API calls)

## License

This project is licensed under the MIT License - see the LICENSE file for details.
