package main

import (
	"math/rand"
	"fmt"
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

var Price []int = []int{
	70,
	66,
	35,
	79,
	79,
	108,
	75,
	0,
	321,
}

var Size []int = []int{
	500,
	500,
	500,
	2000,
	500,
	1500,
	500,
	0,
	2000,
}

func main() {
	fmt.Println("id,name,price,size")
	for i := 0; i < 100000; i++ {
		rand.Seed(time.Now().UnixNano())
		var rand_int int = rand.Intn(8)
		fmt.Printf("%d,%s,%d,%d\n", i, Water(rand_int), Price[rand_int], Size[rand_int])

	}
}
