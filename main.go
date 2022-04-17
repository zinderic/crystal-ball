package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/zinderic/crystal-ball/orb"
)

const (
	MoonCount  = 8
	TarotCount = 78
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func readInput() (string, error) {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Print("Enter your question: ")
	_, err := fmt.Fscan(reader)
	if err != nil {
		return "", err
	}
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.TrimSuffix(text, "\n")
	return text, nil
}

func main() {
	input, err := readInput()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf(">>> %s <<<\n", input)
	pickedMoonData := rand.Intn(MoonCount)
	pickedTarotData := rand.Intn(TarotCount)
	fmt.Println("Projecting point in time..")
	fmt.Println("Calibrating universal truth.. 42")
	fmt.Println("Moon phase is " + orb.MoonData[pickedMoonData].MoonPhase + " with " + orb.MoonData[pickedMoonData].Visibility + " visibility...")
	fmt.Println("Picked " + orb.TarotData[pickedTarotData].Name + " tarot card of the " + orb.TarotData[pickedTarotData].Arcana + "...")

}
