package variables

import "math"

type Player struct {
	X             int
	Y             int
	Z             float32
	Speed         float32
	RotationSpeed float32
}

type Engine struct {
	FOV       float32
	StepRay   float32
	Depth     float32
	MapHeight int
	MapWidth  int
}

type Screen struct {
	Width  int
	Height int
}

var player Player = Player{}
var engine Engine = Engine{}
var screen Screen = Screen{}

func InitializationPlayer(x int, y int) Player {
	if x == 0 {
		x = 1
	}
	if y == 0 {
		y = 1
	}
	player = Player{x, y, 0, 0.1, 0.01}
	return player
}

func InitializationEngine() Engine {
	engine = Engine{math.Pi / 4, 0.1, 16, 16, 16}
	return engine
}

func InitializationScreen() Screen {
	screen = Screen{120, 40}
	return screen
}
