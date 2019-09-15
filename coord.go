package main

import (
	"math"
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
	RotAntiClock = mat{{1, -1}, {1, 1}}
	RotClock     = mat{{1, 1}, {-1, 1}}
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

func (c Coord) getMoves(c1 Coord) move {
	baseDir := NewCoord(0, 0)
	rot := []mat{}

	distX := c1.X - c.X
	distY := c1.Y - c.Y

	ratio := 10
	if distX != 0 {
		ratio = int(math.Round(float64(distY) / float64(distX)))
	}

	switch ratio {
	case 0:
		if distX >= 0 && distY > 0 {
			baseDir = dirE
			rot = []mat{RotClock, RotAntiClock}
		} else if distX >= 0 {
			baseDir = dirE
			rot = []mat{RotAntiClock, RotClock}
		} else if distY >= 0 {
			baseDir = dirW
			rot = []mat{RotAntiClock, RotClock}
		} else {
			baseDir = dirW
			rot = []mat{RotClock, RotAntiClock}
		}

	case 1:
		if distX >= 0 && distX > distY {
			baseDir = dirSE
			rot = []mat{RotAntiClock, RotClock}
		} else if distX >= 0 {
			baseDir = dirSE
			rot = []mat{RotClock, RotAntiClock}
		} else if distX >= distY {
			baseDir = dirNW
			rot = []mat{RotClock, RotAntiClock}
		} else {
			baseDir = dirNW
			rot = []mat{RotAntiClock, RotClock}
		}

	case -1:
		if distX >= 0 && distX > -distY {
			baseDir = dirNE
			rot = []mat{RotClock, RotAntiClock}
		} else if distX >= 0 {
			baseDir = dirNE
			rot = []mat{RotAntiClock, RotClock}
		} else if -distX >= distY {
			baseDir = dirSW
			rot = []mat{RotAntiClock, RotClock}
		} else {
			baseDir = dirSW
			rot = []mat{RotClock, RotAntiClock}
		}

	default:
		if distY >= 0 && distX > 0 {
			baseDir = dirS
			rot = []mat{RotAntiClock, RotClock}
		} else if distY >= 0 {
			baseDir = dirS
			rot = []mat{RotClock, RotAntiClock}
		} else if distX >= 0 {
			baseDir = dirN
			rot = []mat{RotClock, RotAntiClock}
		} else {
			baseDir = dirN
			rot = []mat{RotAntiClock, RotClock}
		}
	}

	m := move{
		distance:   abs(distX) + abs(distY),
		directions: make([]Coord, 8),
	}

	m.directions[0] = baseDir
	for i := 1; i < len(m.directions); i += 2 {
		c1 := m.directions[i-1]
		m.directions[i] = c1.Rot(rot[0])
		if i+1 < len(m.directions) {
			m.directions[i+1] = c1.Rot(rot[1])
		}
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
