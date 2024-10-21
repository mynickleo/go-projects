package main

import (
	"fmt"
	"game/internal/variables"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/eiannone/keyboard"
)

const (
	wallChar      = '#'
	emptyChar     = ' '
	maxObjects    = 10
	maxObjectSize = 20
)

func getBlock(gameMap []rune, x, y int, engine variables.Engine) rune {
	if x >= 0 && x < engine.MapWidth && y >= 0 && y <= engine.MapHeight {
		return gameMap[y*engine.MapWidth+x]
	}
	return wallChar
}

func display(screenConsole []string, screen variables.Screen) string {
	var displayOutput string
	for i, char := range screenConsole {
		if i%screen.Width == 0 {
			displayOutput += "\n"
		}
		displayOutput += char
	}
	return displayOutput
}

func displayMap(gameMap []rune, engine variables.Engine) {
	fmt.Println("PRESS 'Space' to change map")
	for i, char := range gameMap {
		if i%engine.MapWidth == 0 {
			fmt.Print("\n")
		}
		fmt.Print(string(char))
	}
	fmt.Println("\nPRESS 'Enter' to start the game")
}

func initializeEmptyMap(engine variables.Engine) []rune {
	gameMap := make([]rune, engine.MapWidth*engine.MapHeight)
	for x := 0; x < engine.MapWidth; x++ {
		for y := 0; y < engine.MapHeight; y++ {
			if x == 0 || x == engine.MapWidth-1 || y == 0 || y == engine.MapHeight-1 {
				gameMap[y*engine.MapWidth+x] = wallChar
			} else {
				gameMap[y*engine.MapWidth+x] = emptyChar
			}
		}
	}
	return gameMap
}

func createRandomMapObjects(gameMap []rune, engine variables.Engine) {
	rand.Seed(time.Now().UnixNano())
	objectCount := rand.Intn(maxObjects)

	for i := 0; i < objectCount; i++ {
		mapX := rand.Intn(engine.MapWidth-2) + 1
		mapY := rand.Intn(engine.MapHeight-2) + 1
		gameMap[mapY*engine.MapWidth+mapX] = wallChar

		objectSize := rand.Intn(maxObjectSize)
		for j := 0; j < objectSize; j++ {
			vectorX := rand.Intn(3) - 1
			vectorY := rand.Intn(3) - 1
			mapX += vectorX
			mapY += vectorY
			if mapX > 0 && mapX < engine.MapWidth-1 && mapY > 0 && mapY < engine.MapHeight-1 {
				gameMap[mapY*engine.MapWidth+mapX] = wallChar
			}
		}
	}
}

func drawDisplay(player variables.Player, screen variables.Screen, engine variables.Engine, gameMap []rune, screenConsole []string) {
	for x := 0; x < screen.Width; x++ {
		rayCast := (player.Z - engine.FOV/2) + float32((x/screen.Width)*int(engine.FOV))

		eyeX := math.Sin(float64(rayCast))
		eyeY := math.Cos(float64(rayCast))
		distance := float32(0)

		for distance < engine.Depth {
			distance += engine.StepRay
			pointX := math.Round(float64(player.X) + eyeX*float64(distance))
			pointY := math.Round(float64(player.Y) + eyeY*float64(distance))

			if getBlock(gameMap, int(pointX), int(pointY), engine) == wallChar {
				break
			}
		}

		ceiling := math.Round(float64(screen.Height)/2 - float64(screen.Height)/float64(distance))
		floor := math.Round(float64(screen.Height) - ceiling)

		var shade rune

		for y := 0; y < screen.Height; y++ {
			switch {
			case y <= int(ceiling):
				shade = '`'
			case y > int(ceiling) && y <= int(floor):
				switch {
				case distance <= engine.Depth/4:
					shade = '█'
				case distance <= engine.Depth/3:
					shade = '▓'
				case distance <= engine.Depth/2:
					shade = '▒'
				case distance <= engine.Depth:
					shade = '░'
				default:
					shade = emptyChar
				}
			default:
				bottom := float32(1 - (y-screen.Height/2)/(screen.Height/2))
				switch {
				case bottom < 0.25:
					shade = wallChar
				case bottom < 0.5:
					shade = 'x'
				case bottom < 0.75:
					shade = '.'
				default:
					shade = emptyChar
				}
			}

			screenConsole[y*screen.Width+x] = string(shade)
		}
		fmt.Println(display(screenConsole, screen))
	}
}

func main() {
	log.Println("Starting program...")

	engine := variables.InitializationEngine()
	player := variables.InitializationPlayer(1, 1)
	screen := variables.InitializationScreen()
	gameMap := initializeEmptyMap(engine)
	displayMap(gameMap, engine)

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	START := true

	screenConsole := make([]string, screen.Height*screen.Width)

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		if key == keyboard.KeyEsc {
			fmt.Println("Exiting the game...")
			break
		}

		if START {
			if key == keyboard.KeySpace {
				gameMap = initializeEmptyMap(engine)
				createRandomMapObjects(gameMap, engine)
				displayMap(gameMap, engine)
			}
			if key == keyboard.KeyEnter {
				START = false
				drawDisplay(player, screen, engine, gameMap, screenConsole)
			}
		}
	}
}
