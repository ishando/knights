package main

import (
	"flag"
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

// Container -
type Container interface {
	Contains(c Coord) bool
	// Draw(s *tl.Screen)
}

var gameObjects = []Container{}

func main() {
	rand.Seed(time.Now().UnixNano())

	speed := flag.Int("speed", 3, "speed (1-5)")
	players := flag.Int("knights", 3, "knights (2-6)")
	flag.Parse()

	if *speed < 1 {
		*speed = 1
	} else if *speed > 5 {
		*speed = 5
	}
	if *players < 2 {
		*players = 2
	} else if *players > 6 {
		*players = 6
	}

	game = tl.NewGame()
	game.SetDebugOn(true)

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorGreen,
		// Ch: charGrass,
	})
	level.SetOffset(60, 10)

	areaWidth, areaHeight = 40+*players, 30+*players

	// cnv := tl.NewCanvas(areaWidth, areaHeight)

	gameObjects = append(gameObjects, NewBorder(areaWidth, areaHeight))
	gameObjects = append(gameObjects, NewTrees(150+*players*2))
	gameObjects = append(gameObjects, NewTraps(*players+7))
	gameObjects = append(gameObjects, NewTemples(*players+2))

	gameObjects = append(gameObjects, NewKnights(*players))

	for _, gObj := range gameObjects {
		level.AddEntity(gObj.(tl.Drawable))
	}
	for _, gk := range gameObjects[objKnights].(*Knights).knights {
		level.AddEntity(gk.text)
	}

	game.Screen().SetLevel(level)
	game.Screen().SetFps(float64(*speed))
	// game.SetEndKey(tl.KeyCtrlQ)
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
