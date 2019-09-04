package main

import (
	tl "github.com/JoelOtter/termloop"
)

const (
	treeColour = tl.ColorGreen | tl.AttrBold
	treeChar   = '⋏' //♠♣
)

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
		t.coords[*pos] = 1
	}

	return t
}

// Contains -
func (t *Trees) Contains(c Coord) bool {
	_, exists := t.coords[c]
	return exists
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
