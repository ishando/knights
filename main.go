package main

import (
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

const (
	objBorder  = 0
	objTrees   = 1
	objTraps   = 2
	objTemples = 3
	objKnights = 4

	charGrass = '„'

	background    = tl.ColorBlack
	altBackground = tl.ColorBlack
)

var numKnights = 3
var combattants = []Knight{}
var areaWidth, areaHeight int
var game *tl.Game

// Containers -
type Containers interface {
	Contains(c Coord) bool
	// Draw(s *tl.Screen)
}

var gameObjects = []Containers{}

func main() {
	rand.Seed(time.Now().UnixNano())

	game = tl.NewGame()
	game.SetDebugOn(true)

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorGreen,
		// Ch: charGrass,
	})
	level.SetOffset(60, 10)

	areaWidth, areaHeight = 40, 30

	// cnv := tl.NewCanvas(areaWidth, areaHeight)

	gameObjects = append(gameObjects, NewBorder(areaWidth, areaHeight))
	gameObjects = append(gameObjects, NewTrees(150))
	gameObjects = append(gameObjects, NewTraps(10))
	gameObjects = append(gameObjects, NewTemples(5))

	gameObjects = append(gameObjects, NewKnights(3))

	// for _, k := range combattants {
	// 	k.SetCell(k.Position.X, k.Position.Y, &tl.Cell{Bg: k.BgColour, Fg: k.Colour, Ch: charKnight})
	// }

	level.AddEntity(gameObjects[0].(*Border))
	level.AddEntity(gameObjects[1].(*Trees))
	level.AddEntity(gameObjects[2].(*Traps))
	level.AddEntity(gameObjects[3].(*Temples))
	level.AddEntity(gameObjects[4].(*Knights))

	// for _, gO := range gameObjects {
	// 	level.AddEntity(gO)
	// }

	game.Screen().SetLevel(level)
	game.Screen().SetFps(2)
	game.Start()
}

// EndGame -
func EndGame() {
	var colour tl.Attr

	for i, k := range gameObjects[objKnights].(*Knights).knights {
		if k.Health > 0 {
			colour = colours[i]
			break
		}
	}
	endLevel := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: colour,
	})
	endLevel.AddEntity(tl.NewRectangle(60, 10, 30, 30, colour))

	game.Screen().SetLevel(endLevel)
}
