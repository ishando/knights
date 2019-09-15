package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRot(t *testing.T) {
	testcases := []Coord{
		{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1},
	}

	for i, tc := range testcases {
		resC := tc.Rot(RotClock)
		assert.Equal(t, testcases[(i+1)%8], resC, "clockwise rot mismatch for %d", i)
		resA := tc.Rot(RotAntiClock)
		if i == 0 {
			assert.Equalf(t, testcases[7], resA, "anticlockwise rot mismatch for %d", i)
		} else {
			assert.Equalf(t, testcases[i-1], resA, "anticlockwise rot mismatch for %d", i)
		}
	}
}

func TestMove(t *testing.T) {
	c1 := NewCoord(1, 3)
	c2 := NewCoord(1, 3)
	c3 := NewCoord(-4, -1)

	e1 := NewCoord(2, 6)
	e2 := NewCoord(-2, 5)

	c1.Move(c2)
	assert.Equal(t, e1, c1)
	c1.Move(c3)
	assert.Equal(t, e2, c1)
	c1.Unmove(c3)
	assert.Equal(t, e1, c1)
}

func TestInvert(t *testing.T) {
	c1 := NewCoord(1, 3)
	c2 := NewCoord(-1, 3)
	c3 := NewCoord(-4, -1)

	e1 := NewCoord(-1, -3)
	e2 := NewCoord(1, -3)
	e3 := NewCoord(4, 1)

	c1.Invert()
	c2.Invert()
	c3.Invert()

	assert.Equal(t, e1, c1)
	assert.Equal(t, e2, c2)
	assert.Equal(t, e3, c3)
}

func TestSetRandom(t *testing.T) {

	c1 := NewCoord(0, 0)
	areaWidth, areaHeight = 10, 20
	for i := 0; i < 20; i++ {
		c1.SetRandom()

		assert.True(t, c1.X >= 0 && c1.X < areaWidth)
		assert.True(t, c1.Y >= 0 && c1.Y < areaHeight)
	}
}

func TestAdjacent(t *testing.T) {
	c1 := []Coord{{0, 0}, {0, 1}, {1, 2}, {0, 3}, {-1, 4}}

	for i := 0; i < len(c1)-1; i++ {
		for j := i + 1; j < len(c1); j++ {
			if j == i+1 {
				assert.True(t, c1[i].adjacent(c1[j]))
			} else {
				assert.False(t, c1[i].adjacent(c1[j]))
			}
		}
	}
}

func TestGetMoves(t *testing.T) {
	c1 := NewCoord(10, 10)

	testCases := []struct {
		name   string
		input  Coord
		expect []Coord
	}{
		{
			name:   "N then NE",
			input:  NewCoord(12, 3),
			expect: []Coord{dirN, dirNE, dirNW, dirE},
		},
		{
			name:   "N then NW",
			input:  NewCoord(9, 4),
			expect: []Coord{dirN, dirNW, dirNE, dirW},
		},
		{
			name:   "NE then E",
			input:  NewCoord(15, 4),
			expect: []Coord{dirNE, dirE, dirN, dirSE},
		},
		{
			name:   "NE then N",
			input:  NewCoord(14, 5),
			expect: []Coord{dirNE, dirN, dirE, dirNW},
		},
		{
			name:   "E then SE",
			input:  NewCoord(15, 11),
			expect: []Coord{dirE, dirSE, dirNE, dirS},
		},
		{
			name:   "E then NE",
			input:  NewCoord(15, 9),
			expect: []Coord{dirE, dirNE, dirSE, dirN},
		},
		{
			name:   "SE then S",
			input:  NewCoord(16, 18),
			expect: []Coord{dirSE, dirS, dirE, dirSW},
		},
		{
			name:   "SE then E",
			input:  NewCoord(17, 14),
			expect: []Coord{dirSE, dirE, dirS, dirNE},
		},
		{
			name:   "S then SW",
			input:  NewCoord(9, 19),
			expect: []Coord{dirS, dirSW, dirSE, dirW},
		},
		{
			name:   "S then SE",
			input:  NewCoord(12, 19),
			expect: []Coord{dirS, dirSE, dirSW, dirE},
		},
		{
			name:   "SW then W",
			input:  NewCoord(2, 15),
			expect: []Coord{dirSW, dirW, dirS, dirNW},
		},
		{
			name:   "SW then S",
			input:  NewCoord(3, 15),
			expect: []Coord{dirSW, dirS, dirW, dirSE},
		},
		{
			name:   "W then NW",
			input:  NewCoord(4, 8),
			expect: []Coord{dirW, dirNW, dirSW, dirN},
		},
		{
			name:   "W then SW",
			input:  NewCoord(4, 11),
			expect: []Coord{dirW, dirSW, dirNW, dirS},
		},
		{
			name:   "NW then N",
			input:  NewCoord(3, 2),
			expect: []Coord{dirNW, dirN, dirSW, dirNE},
		},
		{
			name:   "NW then W",
			input:  NewCoord(3, 4),
			expect: []Coord{dirNW, dirW, dirN, dirSW},
		},
	}

	for _, tc := range testCases {
		m := c1.getMoves(tc.input)
		for i, dir := range tc.expect {
			assert.Equalf(t, dir, m.directions[i], "%dth direction for %s", i, tc.name)
		}
		// assert.Equalf(t, tc.expect[0], m.directions[0], "direction for %s", tc.name)
	}
}
