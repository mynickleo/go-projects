package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"

	"github.com/eiannone/keyboard"
)

type Player struct {
	x      int
	y      int
	sprite string
}

const (
	height       = 16
	width        = 32
	wall         = "#"
	freeSpace    = " "
	playerSprite = "@"
	pointSprite  = "^"
	objectSprite = "*"
	maxObjects   = 7
)

var points int = 0
var isUpdatingMap bool = false

func getRandomCoordinates() (int, int) {
	x := rand.Intn(width-2) + 1
	y := rand.Intn(height-2) + 1
	return x, y
}

func fillGameMap(gameMap [height][width]string, player Player) [height][width]string {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if y == 0 || y == height-1 || x == 0 || x == width-1 {
				gameMap[y][x] = wall
			} else {
				gameMap[y][x] = freeSpace
			}
		}
	}

	pointX, pointY := getRandomCoordinates()
	gameMap[pointY][pointX] = pointSprite

	gameMap[player.y][player.x] = player.sprite

	for i := 0; i < maxObjects; i++ {
		x, y := getRandomCoordinates()
		groupSize := rand.Intn(3) + 3

		for j := 0; j < groupSize; j++ {
			offsetX := rand.Intn(3) - 1
			offsetY := rand.Intn(3) - 1

			newX := x + offsetX
			newY := y + offsetY

			if newX > 0 && newX < width-1 && newY > 0 && newY < height-1 {
				gameMap[newY][newX] = objectSprite
			}
		}
	}

	return gameMap
}

func displayScreen(gameMap [height][width]string) {
	clearConsole()

	fmt.Println("Points: ", points)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Print(gameMap[y][x])
		}
		fmt.Println()
	}
}

func clearConsole() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func movePlayerInDisplay(player Player, vectorX int, vectorY int, gameMap [height][width]string) ([height][width]string, Player) {
	if vectorX != 0 && gameMap[player.y][player.x+vectorX] != wall {
		gameMap[player.y][player.x] = freeSpace
		player.x += vectorX

		if gameMap[player.y][player.x] == pointSprite {
			points += 1
			isUpdatingMap = true
		}

		if gameMap[player.y][player.x] == objectSprite {
			points -= 1
			isUpdatingMap = true
		}

		gameMap[player.y][player.x] = playerSprite
	}

	if vectorY != 0 && gameMap[player.y+vectorY][player.x] != wall {
		gameMap[player.y][player.x] = freeSpace
		player.y += vectorY

		if gameMap[player.y][player.x] == pointSprite {
			points += 1
			isUpdatingMap = true
		}

		if gameMap[player.y][player.x] == objectSprite {
			points -= 1
			isUpdatingMap = true
		}

		gameMap[player.y][player.x] = playerSprite
	}

	return gameMap, player
}

func main() {
	var gameMap [height][width]string
	player := Player{8, 8, playerSprite}

	gameMap = fillGameMap(gameMap, player)

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	for {
		displayScreen(gameMap)

		char, _, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'w':
			gameMap, player = movePlayerInDisplay(player, 0, -1, gameMap)
		case 's':
			gameMap, player = movePlayerInDisplay(player, 0, 1, gameMap)
		case 'a':
			gameMap, player = movePlayerInDisplay(player, -1, 0, gameMap)
		case 'd':
			gameMap, player = movePlayerInDisplay(player, 1, 0, gameMap)
		case 'q':
			fmt.Println("Exiting game...")
			return
		}

		if isUpdatingMap {
			player.x, player.y = getRandomCoordinates()
			gameMap = fillGameMap(gameMap, player)
			isUpdatingMap = false
		}
	}
}
