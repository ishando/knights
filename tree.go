package main

import (
	"math/rand"

	tl "github.com/JoelOtter/termloop"
)

const (
	treeColour = tl.ColorGreen | tl.AttrBold
	treeChar   = '⋏' //♠♣
)

var ticks = 0
var growth = 60 + rand.Intn(70)

// Trees -
type Trees struct {
	*tl.Entity
	coords map[Coord]int
}

// NewTrees -
func NewTrees(size int) *Trees {
	t := new(Trees)
	t.Entity = tl.NewEntity(1, 1, 1, 1)

	t.coords = make(map[Coord]int)

	for i := 0; i < size; i++ {
		pos := NewRandomCoord()
		t.coords[pos] = 1
	}

	return t
}

// Contains -
func (t *Trees) Contains(c Coord) bool {
	_, exists := t.coords[c]
	return exists
}

// Tick -
func (t *Trees) Tick(ev tl.Event) {
	ticks++
	if ticks%growth == 0 {
		t.newTree()
	}
}

func (t *Trees) newTree() {
	pos := NewRandomCoord()
	t.coords[pos] = 1
}

// Draw -
func (t *Trees) Draw(screen *tl.Screen) {
	if t == nil {
		return
	}

	for c := range t.coords {
		screen.RenderCell(c.X, c.Y, &tl.Cell{
			Fg: treeColour,
			Ch: treeChar,
		})
	}
}
