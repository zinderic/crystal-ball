package orb

import (
	"log"
	"math/rand"
	"os"
	"text/template"
)

const TarotCount = 78

func Tarot() error {
	count := rand.Intn(TarotCount)
	tmpl, err := template.ParseFiles("orb/Tarot.txt")
	// Capture any error
	if err != nil {
		log.Fatalln(err)
	}
	// Print out the template to std
	err = tmpl.Execute(os.Stdout, TarotData[count])
	if err != nil {
		return err
	}
	return nil
}

type TarotStr struct {
	Name   string
	Number string
	Arcana string
	Suit   string
}

var TarotData = []TarotStr{
	{
		Name:   "The Fool",
		Number: "0",
		Arcana: "Major Arcana",
		Suit:   "",
	},
	{
		Name:   "The Magician",
		Number: "1",
		Arcana: "Major Arcana",
		Suit:   "",
	},
	{
		Name:   "The High Priestess",
		Number: "2",
		Arcana: "Major Arcana",
		Suit:   "",
	},
	{
		Name:   "The Empress",
		Number: "3",
		Arcana: "Major Arcana",
		Suit:   "",
	},
	{
		Name:   "The Emperor",
		Number: "4",
		Arcana: "Major Arcana",
		Suit:   "",
	},
	{
		Name:   "The Hierophant",
		Number: "5",
		Arcana: "Major Arcana",
		Suit:   "",
	},
	{
		Name:   "The Lovers",
		Number: "6",
		Arcana: "Major Arcana",
		Suit:   "",
	},
	{
		Name:   "The Chariot",
		Number: "7",
		Arcana: "Major Arcana",
		Suit:   "",
	},
	{
		Name:   "Strength",
		Number: "8",
		Arcana: "Major Arcana",
		Suit:   "",
	},
	{
		Name:   "The Hermit",
		Number: "9",
		Arcana: "Major Arcana",
		Suit:   "",
	},
	{
		Name:   "Wheel of Fortune",
		Number: "10",
		Arcana: "Major Arcana",
		Suit:   "",
	},
	{
		Name:   "Justice",
		Number: "11",
		Arcana: "Major Arcana",
		Suit:   "",
	},
	{
		Name:   "The Hanged Man",
		Number: "12",
		Arcana: "Major Arcana",
		Suit:   "",
	},
	{
		Name:   "Death",
		Number: "13",
		Arcana: "Major Arcana",
		Suit:   "",
	},
	{
		Name:   "Temperance",
		Number: "14",
		Arcana: "Major Arcana",
		Suit:   "",
	},
	{
		Name:   "The Devil",
		Number: "15",
		Arcana: "Major Arcana",
		Suit:   "",
	},
	{
		Name:   "The Tower",
		Number: "16",
		Arcana: "Major Arcana",
		Suit:   "",
	},
	{
		Name:   "The Star",
		Number: "17",
		Arcana: "Major Arcana",
		Suit:   "",
	},
	{
		Name:   "The MoonStr",
		Number: "18",
		Arcana: "Major Arcana",
		Suit:   "",
	},
	{
		Name:   "The Sun",
		Number: "19",
		Arcana: "Major Arcana",
		Suit:   "",
	},
	{
		Name:   "Judgement",
		Number: "20",
		Arcana: "Major Arcana",
		Suit:   "",
	},
	{
		Name:   "The World",
		Number: "21",
		Arcana: "Major Arcana",
		Suit:   "",
	},
	{
		Name:   "Ace of Cups",
		Number: "1",
		Arcana: "Minor Arcana",
		Suit:   "Cups",
	},
	{
		Name:   "Two of Cups",
		Number: "2",
		Arcana: "Minor Arcana",
		Suit:   "Cups",
	},
	{
		Name:   "Three of Cups",
		Number: "3",
		Arcana: "Minor Arcana",
		Suit:   "Cups",
	},
	{
		Name:   "Four of Cups",
		Number: "4",
		Arcana: "Minor Arcana",
		Suit:   "Cups",
	},
	{
		Name:   "Five of Cups",
		Number: "5",
		Arcana: "Minor Arcana",
		Suit:   "Cups",
	},
	{
		Name:   "Six of Cups",
		Number: "6",
		Arcana: "Minor Arcana",
		Suit:   "Cups",
	},
	{
		Name:   "Seven of Cups",
		Number: "7",
		Arcana: "Minor Arcana",
		Suit:   "Cups",
	},
	{
		Name:   "Eight of Cups",
		Number: "8",
		Arcana: "Minor Arcana",
		Suit:   "Cups",
	},
	{
		Name:   "Nine of Cups",
		Number: "9",
		Arcana: "Minor Arcana",
		Suit:   "Cups",
	},
	{
		Name:   "Ten of Cups",
		Number: "10",
		Arcana: "Minor Arcana",
		Suit:   "Cups",
	},
	{
		Name:   "Page of Cups",
		Number: "11",
		Arcana: "Minor Arcana",
		Suit:   "Cups",
	},
	{
		Name:   "Knight of Cups",
		Number: "12",
		Arcana: "Minor Arcana",
		Suit:   "Cups",
	},
	{
		Name:   "Queen of Cups",
		Number: "13",
		Arcana: "Minor Arcana",
		Suit:   "Cups",
	},
	{
		Name:   "King of Cups",
		Number: "14",
		Arcana: "Minor Arcana",
		Suit:   "Cups",
	},
	{
		Name:   "Ace of Swords",
		Number: "1",
		Arcana: "Minor Arcana",
		Suit:   "Swords",
	},
	{
		Name:   "Two of Swords",
		Number: "2",
		Arcana: "Minor Arcana",
		Suit:   "Swords",
	},
	{
		Name:   "Three of Swords",
		Number: "3",
		Arcana: "Minor Arcana",
		Suit:   "Swords",
	},
	{
		Name:   "Four of Swords",
		Number: "4",
		Arcana: "Minor Arcana",
		Suit:   "Swords",
	},
	{
		Name:   "Five of Swords",
		Number: "5",
		Arcana: "Minor Arcana",
		Suit:   "Swords",
	},
	{
		Name:   "Six of Swords",
		Number: "6",
		Arcana: "Minor Arcana",
		Suit:   "Swords",
	},
	{
		Name:   "Seven of Swords",
		Number: "7",
		Arcana: "Minor Arcana",
		Suit:   "Swords",
	},
	{
		Name:   "Eight of Swords",
		Number: "8",
		Arcana: "Minor Arcana",
		Suit:   "Swords",
	},
	{
		Name:   "Nine of Swords",
		Number: "9",
		Arcana: "Minor Arcana",
		Suit:   "Swords",
	},
	{
		Name:   "Ten of Swords",
		Number: "10",
		Arcana: "Minor Arcana",
		Suit:   "Swords",
	},
	{
		Name:   "Page of Swords",
		Number: "11",
		Arcana: "Minor Arcana",
		Suit:   "Swords",
	},
	{
		Name:   "Knight of Swords",
		Number: "12",
		Arcana: "Minor Arcana",
		Suit:   "Swords",
	},
	{
		Name:   "Queen of Swords",
		Number: "13",
		Arcana: "Minor Arcana",
		Suit:   "Swords",
	},
	{
		Name:   "King of Swords",
		Number: "14",
		Arcana: "Minor Arcana",
		Suit:   "Swords",
	},
	{
		Name:   "Ace of Wands",
		Number: "1",
		Arcana: "Minor Arcana",
		Suit:   "Wands",
	},
	{
		Name:   "Two of Wands",
		Number: "2",
		Arcana: "Minor Arcana",
		Suit:   "Wands",
	},
	{
		Name:   "Three of Wands",
		Number: "3",
		Arcana: "Minor Arcana",
		Suit:   "Wands",
	},
	{
		Name:   "Four of Wands",
		Number: "4",
		Arcana: "Minor Arcana",
		Suit:   "Wands",
	},
	{
		Name:   "Five of Wands",
		Number: "5",
		Arcana: "Minor Arcana",
		Suit:   "Wands",
	},
	{
		Name:   "Six of Wands",
		Number: "6",
		Arcana: "Minor Arcana",
		Suit:   "Wands",
	},
	{
		Name:   "Seven of Wands",
		Number: "7",
		Arcana: "Minor Arcana",
		Suit:   "Wands",
	},
	{
		Name:   "Eight of Wands",
		Number: "8",
		Arcana: "Minor Arcana",
		Suit:   "Wands",
	},
	{
		Name:   "Nine of Wands",
		Number: "9",
		Arcana: "Minor Arcana",
		Suit:   "Wands",
	},
	{
		Name:   "Ten of Wands",
		Number: "10",
		Arcana: "Minor Arcana",
		Suit:   "Wands",
	},
	{
		Name:   "Page of Wands",
		Number: "11",
		Arcana: "Minor Arcana",
		Suit:   "Wands",
	},
	{
		Name:   "Knight of Wands",
		Number: "12",
		Arcana: "Minor Arcana",
		Suit:   "Wands",
	},
	{
		Name:   "Queen of Wands",
		Number: "13",
		Arcana: "Minor Arcana",
		Suit:   "Wands",
	},
	{
		Name:   "King of Wands",
		Number: "14",
		Arcana: "Minor Arcana",
		Suit:   "Wands",
	},
	{
		Name:   "Ace of Pentacles",
		Number: "1",
		Arcana: "Minor Arcana",
		Suit:   "Pentacles",
	},
	{
		Name:   "Two of Pentacles",
		Number: "2",
		Arcana: "Minor Arcana",
		Suit:   "Pentacles",
	},
	{
		Name:   "Three of Pentacles",
		Number: "3",
		Arcana: "Minor Arcana",
		Suit:   "Pentacles",
	},
	{
		Name:   "Four of Pentacles",
		Number: "4",
		Arcana: "Minor Arcana",
		Suit:   "Pentacles",
	},
	{
		Name:   "Five of Pentacles",
		Number: "5",
		Arcana: "Minor Arcana",
		Suit:   "Pentacles",
	},
	{
		Name:   "Six of Pentacles",
		Number: "6",
		Arcana: "Minor Arcana",
		Suit:   "Pentacles",
	},
	{
		Name:   "Seven of Pentacles",
		Number: "7",
		Arcana: "Minor Arcana",
		Suit:   "Pentacles",
	},
	{
		Name:   "Eight of Pentacles",
		Number: "8",
		Arcana: "Minor Arcana",
		Suit:   "Pentacles",
	},
	{
		Name:   "Nine of Pentacles",
		Number: "9",
		Arcana: "Minor Arcana",
		Suit:   "Pentacles",
	},
	{
		Name:   "Ten of Pentacles",
		Number: "10",
		Arcana: "Minor Arcana",
		Suit:   "Pentacles",
	},
	{
		Name:   "Page of Pentacles",
		Number: "11",
		Arcana: "Minor Arcana",
		Suit:   "Pentacles",
	},
	{
		Name:   "Knight of Pentacles",
		Number: "12",
		Arcana: "Minor Arcana",
		Suit:   "Pentacles",
	},
	{
		Name:   "Queen of Pentacles",
		Number: "13",
		Arcana: "Minor Arcana",
		Suit:   "Pentacles",
	},
	{
		Name:   "King of Pentacles",
		Number: "14",
		Arcana: "Minor Arcana",
		Suit:   "Pentacles",
	},
}
