// main.go
package main

import (
	"fmt"
)

type gb struct {
	level int
}

func (building gb) Level() int {
	return gb.level
}

type cdm struct {
	gb
}

type player struct {
	name string
	cdm  cdm
}

type tradable interface {
	Level() int
	FpToLevel() int
	FpDonated() map[string]int
	Rewards() map[int]int
	Reward(string)
}

func main() {
	fmt.Println("Hello World!")
}
