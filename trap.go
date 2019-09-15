package main

import (
	"math/rand"

	tl "github.com/JoelOtter/termloop"
)

const (
	trapColour = tl.ColorMagenta
	trapChar   = '❖'
)

var hazard = 40 + rand.Intn(80)

// Traps -
type Traps struct {
	*tl.Entity
	coords map[Coord]int
}

// NewTraps -
func NewTraps(size int) *Traps {
	t := new(Traps)
	t.Entity = tl.NewEntity(1, 1, 1, 1)

	t.coords = make(map[Coord]int)

	for len(t.coords) < size {
		pos := NewRandomCoord()
		t.coords[pos] = 1
	}

	return t
}

// Collide -
func (t *Traps) Collide(collision tl.Physical) {
	// check if its a trap collision
	if _, ok := collision.(*Knights); ok {

		traps := []Coord{}
		for tc := range t.coords {
			if gameObjects[objKnights].Contains(tc) {
				traps = append(traps, tc)
				continue
			}
		}
		for _, trp := range traps {
			t.teleport(trp)
		}
	}
}

// Contains -
func (t *Traps) Contains(c Coord) bool {
	_, exists := t.coords[c]
	return exists
}

func (t *Traps) teleport(c Coord) {
	delete(t.coords, c)
	newPos := NewRandomCoord()

	t.coords[newPos] = 1
}

// Draw -
func (t *Traps) Draw(screen *tl.Screen) {
	if t == nil {
		return
	}

	for c := range t.coords {
		screen.RenderCell(c.X, c.Y, &tl.Cell{
			Fg: trapColour,
			Ch: trapChar,
		})
	}
}

// Tick -
func (t *Traps) Tick(ev tl.Event) {
	if ticks%hazard == 0 {
		t.coords[NewRandomCoord()] = 1
	}
}
