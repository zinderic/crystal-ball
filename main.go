package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MoonCount  = 8
	TarotCount = 78
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	pickedMoonData := rand.Intn(MoonCount)
	pickedTarotData := rand.Intn(TarotCount)
	fmt.Println("Projecting point in time..")
	println("Moon phase is " + MoonData[pickedMoonData].MoonPhase + " with " + MoonData[pickedMoonData].Visibility + " visibility...")
	println("Picked " + TarotData[pickedTarotData].Name + " tarot card of the " + TarotData[pickedTarotData].Arcana + "...")

}
