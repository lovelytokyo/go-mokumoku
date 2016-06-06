package main

import (
	"fmt"
	"math/rand"
	"time"
)

//go:generate stringer -type=Water
type Water int

const (
	Irohasu Water = iota
	Volvic
	CrystalGeyser
	AlcaliIonWater
	Perrier
	Contrex
	Wilkinson
	AmebaWater
	Onsensui
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		var rand_int int = rand.Intn(8)
		fmt.Printf("%s\n", Water(rand_int))
	}
}