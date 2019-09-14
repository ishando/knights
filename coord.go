package main

import (
	"math/rand"
)

var (
	dirN  = NewCoord(0, -1)
	dirNW = NewCoord(-1, -1)
	dirW  = NewCoord(-1, 0)
	dirSW = NewCoord(-1, 1)
	dirS  = NewCoord(0, 1)
	dirSE = NewCoord(1, 1)
	dirE  = NewCoord(1, 0)
	dirNE = NewCoord(1, -1)
)

// Coord -
type Coord struct {
	X, Y int
}

// RotMat - rotation matrix
type mat []Coord

// rotation matrices
var (
	RotClock     = mat{{1, -1}, {1, 1}}
	RotAntiClock = mat{{1, 1}, {-1, 1}}
)

// NewCoord -
func NewCoord(x, y int) Coord {
	return Coord{x, y}
}

// NewRandomCoord -
func NewRandomCoord() Coord {
	c := NewCoord(0, 0)
	c.SetRandom()
	return c
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// Move -
func (c *Coord) Move(dir Coord) {
	c.X += dir.X
	c.Y += dir.Y
}

// Unmove -
func (c *Coord) Unmove(dir Coord) {
	c.X -= dir.X
	c.Y -= dir.Y
}

// Invert -
func (c *Coord) Invert() {
	c.X = -c.X
	c.Y = -c.Y
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

func (c Coord) adjacent(c1 Coord) bool {
	if abs(c.X-c1.X) <= 1 && abs(c.Y-c1.Y) <= 1 {
		return true
	}
	return false

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

func (c Coord) compare(c1 Coord) Coord {
	cDir := NewCoord(0, 0)
	switch {
	case c.X < c1.X:
		cDir.X = 1
	case c.X > c1.X:
		cDir.X = -1
	}

	switch {
	case c.Y < c1.Y:
		cDir.Y = 1
	case c.Y > c1.Y:
		cDir.Y = -1
	}

	return cDir
}

func (c Coord) getMoves(c1 Coord) move {
	baseDir := NewCoord(0, 0)
	rot := mat{}

	distX := c1.X - c.X
	distY := c1.Y - c.Y

	ratio := 10
	if distX != 0 {
		ratio = distY / distX
	}

	switch ratio {
	case 0:
		if distX >= 0 && distY > 0 {
			baseDir = dirE
			rot = RotClock
		} else if distX >= 0 {
			baseDir = dirE
			rot = RotAntiClock
		} else if distY >= 0 {
			baseDir = dirW
			rot = RotAntiClock
		} else {
			baseDir = dirW
			rot = RotClock
		}

	case 1:
		if distX >= 0 && distX > distY {
			baseDir = dirSE
			rot = RotAntiClock
		} else if distX >= 0 {
			baseDir = dirSE
			rot = RotClock
		} else if distX >= distY {
			baseDir = dirNW
			rot = RotClock
		} else {
			baseDir = dirNW
			rot = RotAntiClock
		}

	case -1:
		if distX >= 0 && distX > -distY {
			baseDir = dirNE
			rot = RotClock
		} else if distX >= 0 {
			baseDir = dirNE
			rot = RotAntiClock
		} else if -distX >= distY {
			baseDir = dirSW
			rot = RotAntiClock
		} else {
			baseDir = dirSW
			rot = RotClock
		}

	default:
		if distY >= 0 && distX > 0 {
			baseDir = dirS
			rot = RotAntiClock
		} else if distY >= 0 {
			baseDir = dirS
			rot = RotClock
		} else if distX >= 0 {
			baseDir = dirN
			rot = RotClock
		} else {
			baseDir = dirN
			rot = RotAntiClock
		}
	}

	m := move{
		distance:   abs(distX) + abs(distY),
		directions: make([]Coord, 8),
	}

	m.directions[0] = baseDir
	for i := 1; i < len(m.directions); i++ {
		c1 := m.directions[i-1]
		m.directions[i] = c1.Rot(rot)
	}

	return m
}

// Rot - rotate coordinate 45 clockwise
func (c Coord) Rot(r mat) Coord {
	rc := Coord{
		X: (c.X * r[0].X) + (c.Y * r[0].Y),
		Y: (c.X * r[1].X) + (c.Y * r[1].Y),
	}

	if ax := abs(rc.X); ax > 1 {
		rc.X = rc.X / ax
	}
	if ay := abs(rc.Y); ay > 1 {
		rc.Y = rc.Y / ay
	}
	return rc
}
