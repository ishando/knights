package main

import (
	"math/rand"
)

const (
	dirNorth = -1
	dirSouth = 1
	dirWest  = -2
	dirEast  = 2
)

// Coord -
type Coord struct {
	X, Y int
}

// NewCoord -
func NewCoord(x, y int) *Coord {
	return &Coord{x, y}
}

// NewRandomCoord -
func NewRandomCoord() *Coord {
	c := NewCoord(0, 0)
	c.SetRandom()
	return c
}

// Move -
func (c *Coord) Move(dir int) {
	switch dir {
	case dirNorth:
		c.Y--
	case dirSouth:
		c.Y++
	case dirEast:
		c.X++
	case dirWest:
		c.X--
	}
}

// SetRandom -
func (c *Coord) SetRandom() {
	// loop until we find an unoccupied coord
setCoord:
	for {
		c.X = rand.Intn(areaWidth)
		c.Y = rand.Intn(areaHeight)

		for _, o := range gameObjects {
			if o.Contains(*c) {
				continue setCoord
			}
		}

		break
	}
}

func (c *Coord) adjacent(c1 *Coord) bool {
	if c1.X == c.X && c1.Y == c.Y {
		return true
	}
	if c1.X == c.X && (c1.Y-c.Y == 1 || c.Y-c1.Y == 1) {
		return true
	}
	if c1.Y == c.Y && (c1.X-c.X == 1 || c.X-c1.X == 1) {
		return true
	}
	return false
}

func (c *Coord) distance(c1 *Coord) move {
	distX := c1.X - c.X
	distY := c1.Y - c.Y

	m := move{}

	switch distY >= 0 {
	// South
	case true:
		switch distX >= 0 {
		// East
		case true:
			m.distance = distX + distY
			switch distY >= distX {
			case true:
				m.directions = []int{dirSouth, dirEast, dirWest, dirNorth}
			default:
				m.directions = []int{dirEast, dirSouth, dirNorth, dirWest}
			}
		// West
		default:
			m.distance = distY - distX
			switch distY >= distX {
			case true:
				m.directions = []int{dirSouth, dirWest, dirEast, dirNorth}
			default:
				m.directions = []int{dirWest, dirSouth, dirNorth, dirEast}
			}
		}
	// North
	default:
		switch {
		// East
		case distX >= 0:
			m.distance = distX - distY
			switch -distY >= distX {
			case true:
				m.directions = []int{dirNorth, dirEast, dirWest, dirSouth}
			default:
				m.directions = []int{dirEast, dirNorth, dirSouth, dirWest}
			}
		// West
		default:
			m.distance = -(distX + distY)
			switch distY <= distX {
			case true:
				m.directions = []int{dirNorth, dirWest, dirEast, dirSouth}
			default:
				m.directions = []int{dirWest, dirNorth, dirSouth, dirEast}
			}
		}
	}

	return m
}
