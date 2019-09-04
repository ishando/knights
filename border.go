package main

import (
	tl "github.com/JoelOtter/termloop"
)

// Border -
type Border struct {
	*tl.Entity
	width, height int
	coords        map[Coord]int
}

// NewBorder -
func NewBorder(width, height int) *Border {
	b := new(Border)
	b.Entity = tl.NewEntity(1, 1, 1, 1)
	b.width = width - 1
	b.height = height - 1

	b.coords = make(map[Coord]int)

	// top and bottom
	for x := 0; x < b.width; x++ {
		b.coords[Coord{x, 0}] = 1
		b.coords[Coord{x, b.height}] = 1
	}

	// left and right
	for y := 0; y < height; y++ {
		b.coords[Coord{0, y}] = 1
		b.coords[Coord{b.width, y}] = 1
	}

	return b
}

// Contains -
func (b *Border) Contains(c Coord) bool {
	_, exists := b.coords[c]
	return exists
}

// Draw -
func (b *Border) Draw(screen *tl.Screen) {
	if b == nil {
		return
	}

	for c := range b.coords {
		screen.RenderCell(c.X, c.Y, &tl.Cell{
			Bg: tl.ColorMagenta,
		})
	}
}
