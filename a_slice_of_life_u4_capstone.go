package main

import (
	"fmt"
	"math/rand"
)

/*
build a simulation of underpopulation, overpopulation, and reproduction
called Conway's Game of Life
This simulation is played out on a 2D grid of cells
This challenge focuses on SLICES.
Each cell has 8 adjacent cells in the horizontal, vertical and diagonal directions.
In each generation, cells live or die based on the number of living neighbours.
*/
/*
20.1 - A new Universe:
For the first implementation of the Game of Life, limit the universe to a FIXED SIZE.
Decide on the DIMENSIONS of the grid and define some CONSTANTS.
Next, define a Universe type to hold a two-dimentional field of cells. With a Boolean type
each cell will be either dead(false) or alive(true).
Uses slices rather than arrays so that a universe can be shared with, or modified by,
functions or methods.

Write a NewUniverse function that uses 'make' to allocate and return a
Universe with 'height' rows and 'width' columns PER ROW.

Freshly allocated slices will default to the zero value, false in this case,
which means the universe begins empty.
*/

/*
How Universe looks, all populated with false to begin with.
[false][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]
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

/*

 */

const (
	width  = 80
	height = 15
)

//Universe represent alive(true) or dead(false) cells
type Universe [][]bool

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
	// print each row
	// represent live cells (true) with *
	// dead (false) with " "
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
	universe := NewUniverse()
	universe.Seed()
	universe.Show()
}
