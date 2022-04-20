package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/zinderic/crystal-ball/orb"
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

	fmt.Println("Projecting point in time..")
	fmt.Println("Calibrating universal truth.. 42")
	err = orb.Moon()
	if err != nil {
		log.Fatalln(err)
	}
	err = orb.Tarot()
	if err != nil {
		log.Fatalln(err)
	}

}
