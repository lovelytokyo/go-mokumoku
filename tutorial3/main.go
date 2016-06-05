//go:generate gen
package main

import "fmt"

// +gen slice:"Count"
type Water struct {
	id    int
	name  string
	price int
	size  int
}

var irohasu = Water{
	1,
	"Irohasu",
	70,
	500,
}
var volvic = Water{
	2,
	"Volvic",
	66,
	500,
}
var crystalGeyser = Water{
	3,
	"CrystalGeyser",
	35,
	500,
}

func main() {
	var waters WaterSlice = WaterSlice{
		irohasu, volvic, crystalGeyser,
		irohasu, volvic, crystalGeyser,
	}
}