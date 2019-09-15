package main

import (
	"math/rand"
	"sort"

	tl "github.com/JoelOtter/termloop"
)

const (
	templeColour = tl.ColorRed | tl.AttrBold
	templeChar   = '⨦' //✚'
)

var respawn = 50 + rand.Intn(100)

// Temples -
type Temples struct {
	*tl.Entity
	coords map[Coord]int
}

// NewTemples -
func NewTemples(size int) *Temples {
	t := new(Temples)
	t.Entity = tl.NewEntity(1, 1, 1, 1)

	t.coords = make(map[Coord]int)

	for i := 0; i < size; i++ {
		pos := NewRandomCoord()
		t.coords[pos] = 1
	}

	return t
}

// Collide -
func (t *Temples) Collide(collision tl.Physical) {
	// check if its a trap collision
	if _, ok := collision.(*Knights); ok {

		temps := []Coord{}
		for tc := range t.coords {
			if gameObjects[objKnights].Contains(tc) {
				temps = append(temps, tc)
				continue
			}
		}
		for _, tmp := range temps {
			t.close(tmp)
		}
	}
}

// Contains -
func (t *Temples) Contains(c Coord) bool {
	_, exists := t.coords[c]
	return exists
}

func (t *Temples) close(c Coord) {
	delete(t.coords, c)
}

// Closest -
func (t *Temples) Closest(c Coord) *move {
	dists := []move{}
	for tc := range t.coords {
		d := c.getMoves(tc)
		if d.distance <= 4 {
			dists = append(dists, d)
		}
	}
	if len(dists) == 0 {
		return nil
	}
	sort.Slice(dists, func(a, b int) bool { return dists[a].distance < dists[b].distance })
	return &dists[0]
}

// Draw -
func (t *Temples) Draw(screen *tl.Screen) {
	if t == nil {
		return
	}

	for c := range t.coords {
		screen.RenderCell(c.X, c.Y, &tl.Cell{
			Fg: templeColour,
			Ch: templeChar,
		})
	}
}

// Tick -
func (t *Temples) Tick(ev tl.Event) {
	if ticks%respawn == 0 && len(t.coords) == 0 {
		t.coords[NewRandomCoord()] = 1
	}
}
