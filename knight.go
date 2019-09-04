package main

import (
	// "fmt"
	// "math/rand"
	"sort"
	"strconv"

	tl "github.com/JoelOtter/termloop"
)

const (
	knightChar  = 'K' //옷'
	graveChar   = '±'
	maxHealth   = 5
	startHealth = 3

	soundHeal     = "/usr/share/sounds/fredesktop/stereo/complete.oga"
	soundTrap     = "/usr/share/sounds/fredesktop/stereo/camera-shutter.oga"
	soundKill     = "/usr/share/sounds/fredesktop/stereo/trash-empty.oga"
	soundTeleport = "/usr/share/sounds/fredesktop/stereo/bell.oga"
	soundEnd      = "/usr/share/sounds/fredesktop/stereo/service-login.oga"
)

var colours = []tl.Attr{tl.ColorRed, tl.ColorBlue, tl.ColorYellow}

// Knight -
type Knight struct {
	Health  int
	Colour  tl.Attr
	Glyph   rune
	Pos     *Coord
	Dirs    []int
	lastDir int
}

// NewKnight -
func NewKnight(colour tl.Attr, pos *Coord) *Knight {
	k := &Knight{
		Health: startHealth,
		Colour: colour | tl.AttrReverse | tl.AttrUnderline,
		Pos:    pos,
		Glyph:  knightChar,
		Dirs:   []int{dirNorth, dirSouth, dirEast, dirWest},
	}

	return k
}

// Knights -
type Knights struct {
	*tl.Entity
	coords  map[Coord]int
	knights []*Knight
	Alive   int
}

// NewKnights -
func NewKnights(size int) *Knights {
	ks := new(Knights)
	ks.Entity = tl.NewEntity(1, 1, 1, 1)
	ks.Alive = size

	ks.coords = make(map[Coord]int)
	ks.knights = make([]*Knight, size)

	for len(ks.coords) < size {
		i := len(ks.coords)
		pos := NewRandomCoord()
		ks.coords[*pos] = i
		ks.knights[i] = NewKnight(colours[i], pos)
		// fmt.Printf("%+v|", pos)
	}

	return ks
}

// Contains -
func (ks *Knights) Contains(c Coord) bool {
	_, exists := ks.coords[c]
	return exists
}

// Draw -
func (ks *Knights) Draw(screen *tl.Screen) {
	if ks == nil {
		return
	}

	for _, k := range ks.knights {
		screen.RenderCell(k.Pos.X, k.Pos.Y, &tl.Cell{
			Fg: k.Colour,
			Ch: k.Glyph,
		})
		// fmt.Printf("%+v|%d|", c, k.Colour)
	}
}

// Tick -
func (ks *Knights) Tick(ev tl.Event) {
	if ks == nil {
		return
	}

	ks.setDirs()

	for i, k := range ks.knights {
		if k.Health == 0 {
			continue
		}

		moved := false
		oldPos := *k.Pos
		newPos := *k.Pos

		di := 0
		for !moved {
			dir := k.Dirs[di%4]
			if di < 4 && dir == -k.lastDir {
				di++
				continue
			}

			newPos.Move(dir)

			if gameObjects[objTrees].Contains(newPos) {
				newPos = oldPos
				di++
				continue
			}
			k.lastDir = k.Dirs[di%4]

			if gameObjects[objBorder].Contains(newPos) {
				newPos = *NewRandomCoord()
			}

			if gameObjects[objTraps].Contains(newPos) {
				k.Health--
				k.Glyph = []rune(strconv.Itoa(k.Health))[0]
				if k.Health == 0 {
					k.Glyph = graveChar
					ks.Alive--
				}
			}

			if gameObjects[objTemples].Contains(newPos) {
				if k.Health < maxHealth {
					k.Health++
					k.Glyph = []rune(strconv.Itoa(k.Health))[0]
				}
				if k.Health == maxHealth {
					k.Glyph = knightChar
				}
			}

			k.Pos = &newPos
			ks.coords[newPos] = i
			delete(ks.coords, oldPos)
			moved = true
		}
	}

	for i := 0; i < len(ks.knights)-1; i++ {
		for j := i + 1; j < len(ks.knights); j++ {
			if ks.knights[i].Pos.adjacent(ks.knights[j].Pos) {
				ks.battle(i, j)
			}
		}
	}

	if ks.Alive <= 1 {
		EndGame()
	}
}

type move struct {
	distance   int
	directions []int
	index      int
}

func (ks *Knights) setDirs() {

	for i := 0; i < len(ks.knights); i++ {
		kdists := []move{}
		tdist := &move{}

		if ks.knights[i].Health == 0 {
			continue
		}

		for j := 0; j < len(ks.knights); j++ {
			if j == i || ks.knights[j].Health == 0 {
				continue
			}
			d := ks.knights[i].Pos.distance(ks.knights[j].Pos)
			d.index = j
			kdists = append(kdists, d)
		}
		sort.Slice(kdists, func(a, b int) bool { return kdists[a].distance < kdists[b].distance })

		if ks.knights[i].Health < maxHealth {
			tdist = gameObjects[objTemples].(*Temples).Closest(ks.knights[i].Pos)

			if tdist != nil && tdist.distance < kdists[0].distance {
				ks.knights[i].Dirs = tdist.directions
				continue
			}
		}

		if ks.knights[i].Health > ks.knights[kdists[0].index].Health {
			ks.knights[i].Dirs = kdists[0].directions
		} else {
			ks.knights[i].Dirs = []int{kdists[0].directions[3], kdists[0].directions[2], kdists[0].directions[1], kdists[0].directions[0]}
		}
	}
}

func (ks *Knights) battle(i, j int) {
	if ks.knights[i].Health == 0 || ks.knights[j].Health == 0 {
		return
	}

	if ks.knights[i].Health == ks.knights[j].Health {
		ks.knights[i].Health = 0
		ks.knights[j].Health = 0
		ks.knights[i].Glyph = graveChar
		ks.knights[j].Glyph = graveChar
		ks.Alive--
	} else if ks.knights[i].Health > ks.knights[j].Health {
		ks.knights[j].Health = 0
		ks.knights[j].Glyph = graveChar
	} else {
		ks.knights[i].Health = 0
		ks.knights[i].Glyph = graveChar
	}
	ks.Alive--
}
