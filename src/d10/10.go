package d10

import (
	"aoc2023/src/util"
	"strings"

	"github.com/charmbracelet/log"
)

var logger log.Logger

type Map [][]byte

func (m Map) String() string {
	b := strings.Builder{}
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			b.WriteByte(m[y][x])
		}
		b.WriteRune('\n')
	}
	return b.String()
}

type Route struct {
	StepsFromStart int
	X              int
	Y              int
	Next           *Route
	Prev           *Route
	ConnectsEast   bool
	ConnectsWest   bool
	ConnectsSouth  bool
	ConnectsNorth  bool
}

var directionsMap map[byte][]bool = map[byte][]bool{
	'|': {true, false, true, false},
	'-': {false, true, false, true},
	'L': {true, true, false, false},
	'J': {true, false, false, true},
	'7': {false, false, true, true},
	'F': {false, true, true, false},
}

func InitRoute(x, y int, char byte) *Route {
	connections := directionsMap[char]
	return &Route{
		X:             x,
		Y:             y,
		ConnectsNorth: connections[0],
		ConnectsEast:  connections[1],
		ConnectsSouth: connections[2],
		ConnectsWest:  connections[3],
	}
}

func Main() {
	logger = *log.Default()

	data, err := util.ReadExampleInput(10)
	if err != nil {
		panic(err)
	}

	m := parseInput(data)
	log.Info(m)
	one(m)

	m = Map{
		{'0', '1', '2'},
		{'3', '4', '5'},
		{'6', '7', '8'},
	}
	logger.Info(m)
}

func parseInput(data []byte) Map {
	lines := strings.Split(string(data), "\n")
	lenX := len(lines[0])
	lenY := len(lines)
	m := makeMap(lenX, lenY)

	for x := 0; x < lenX; x++ {
		for y := 0; y < lenY; y++ {
			m[y][x] = lines[y][x]
		}
	}

	return m
}

func makeMap(lenX, lenY int) Map {
	m := make([][]byte, lenY)
	for i := range m {
		m[i] = make([]byte, lenX)
	}
	return m
}

func one(input Map) {
	// 1. Make route until looped back to S
	// 2. Walk in reverse until steps is less than StepsFromStart
	startingX, startingY := getStartingPoint(input)
	if startingX == -1 && startingY == -1 {
		logger.Error("could not determine starting point")
		return
	}

	current := &Route{
		StepsFromStart: 0,
		X:              startingX,
		Y:              startingY,
		Prev:           nil,
		Next:           nil,
	}

	north := InitRoute(startingX, startingY-1, input[startingY-1][startingX])
	east := InitRoute(startingX+1, startingY, input[startingY][startingX+1])
	south := InitRoute(startingX, startingY+1, input[startingY+1][startingX])
	west := InitRoute(startingX-1, startingY, input[startingY][startingX-1])

	nextX := -1
	nextY := -1

	if north.ConnectsSouth {
		current.ConnectsNorth = true
		nextX = north.X
		nextY = north.Y
	}
	if east.ConnectsWest {
		current.ConnectsWest = true
		nextX = east.X
		nextY = east.Y
	}
	if south.ConnectsNorth {
		current.ConnectsSouth = true
		nextX = south.X
		nextY = south.Y
	}
	if west.ConnectsEast {
		current.ConnectsWest = true
		nextX = west.X
		nextY = west.Y
	}

	switch true {
	case north.ConnectsSouth:
		nextX = north.X
		nextY = north.Y
		break
	case east.ConnectsWest:
		nextX = east.X
		nextY = east.Y
		break
	case south.ConnectsNorth:
		nextX = south.X
		nextY = south.Y
		break
	case west.ConnectsEast:
		nextX = west.X
		nextY = west.Y
		break
	}

	current.Next = walk(input, current, nextX, nextY)

	logger.Info("part 1: %d", 0)
}

func walk(input Map, previous *Route, currX, currY int) *Route {
	if input[currY][currX] == 'S' {
		return nil
	}

	nextX := currX
	nextY := currY

	route := InitRoute(currX, currY, input[currY][currX])
	route.StepsFromStart = previous.StepsFromStart + 1
	route.Prev = previous

	fromNorth := previous.Y == currY-1
	fromEast := previous.X == currX-1
	fromSouth := previous.Y == currY+1
	fromWest := previous.X == currX+1

	switch {
	case !fromNorth && route.ConnectsNorth:
		nextY++
	case !fromEast && route.ConnectsEast:
		nextX++
	case !fromSouth && route.ConnectsSouth:
		nextY--
	case !fromWest && route.ConnectsWest:
		nextX--
	}

	route.Next = walk(input, previous, nextX, nextY)
	return route
}

func getStartingPoint(input Map) (int, int) {
	for y, row := range input {
		for x, char := range row {
			if char == 'S' {
				return x, y
			}
		}
	}
	return -1, -1
}

func two() {
	logger.Info("part 2: %d", 0)
}
