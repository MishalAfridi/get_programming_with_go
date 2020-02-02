package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
How Universe looks, all populated with false to begin with.
[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]
[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]
[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]
[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]
[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]
[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]
[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]
[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]
[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]
[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]
[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]
[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]
[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]
[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]
[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]
*/

const (
	width  = 80
	height = 15
)

//Universe represent alive(true) or dead(false) cells
type Universe [][]bool

//Step etermines next state of every cell in the universe based on previous state of universe
func Step(a, b Universe) {
	//based on prev state of universe
  for i, row := range a {
    for j := range row {
			if a.Next(j, i) {
				b[i][j] = true
			} else {
        b[i][j] = false
			}
		}
		fmt.Println("\x0c")
	}
}

//Next determines next state of a cell
func (u Universe) Next(x, y int) bool {
	var next bool
  if u.Alive(x, y) == true {
		if u.Neighbours(x, y) == 2 || u.Neighbours(x, y) == 3 {
			next = true
		}
	} else {
		if u.Neighbours(x, y) == 3 {
			next = true
		}
	}
	return next
}

//Alive determines state of a cell and wraps around if out of range
func (u Universe) Alive(x, y int) bool {
	x = (x + width) % width
  y = (y + height) % height
	return u[y][x]
}

//Neighbours counts number of live neighbours per cell
func (u Universe) Neighbours(x, y int) int {
	neighbours := make(map[string]bool)
	count := 0
	
	neighbours["UpLeft"] = u.Alive(x-1, y-1)
	neighbours["Left"] = u.Alive(x-1, y)
	neighbours["DownLeft"] = u.Alive(x-1, y+1)
	neighbours["Up"] = u.Alive(x, y-1)
	neighbours["Down"] = u.Alive(x, y+1)
	neighbours["UpRight"] = u.Alive(x+1, y-1)
	neighbours["Right"] = u.Alive(x+1, y)
	neighbours["DownRight"] = u.Alive(x+1, y+1)

  for i := range neighbours {
		if neighbours[i] == true {
			count ++
		}
	}
	return count
}

//Seed populates randomly approx 25% of the Universe with live cells
func (u Universe) Seed() {
	//random numbers - using math/rand lib rand.Intn()
	//25% of cells in Universe
	percentageOfCells := int(float64(len(u)*len(u[0])) * 0.25)
	for i := 0; i <= percentageOfCells; i++ {
		u[rand.Intn(15)][rand.Intn(80)] = true
	}
	//
}

//Show displays a Universe
func (u Universe) Show() {
	for _, row := range u {
		for _, column := range row {
			if column {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

//NewUniverse creates and returns a new Universe
func NewUniverse() Universe {
	u := (make(Universe, height))
	for row := range u {
		u[row] = make([]bool, width)
	}
	return u
}

func main() {
	a := NewUniverse()
	b := NewUniverse()
	a.Seed()
	for i := 0; i < 300; i++ {
		Step(a, b)
		a.Show()
		time.Sleep(time.Second / 30)
		a, b = b, a // Swap universes
  }

}
