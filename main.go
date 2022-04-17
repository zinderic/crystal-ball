package main

import (
	"fmt"
	"log"
	"math/rand"
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

func main() {
	fmt.Println(orb.Blue + orb.Orb + orb.ResetColor)
	input, err := orb.ReadInput()
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
