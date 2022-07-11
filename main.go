package main

import (
	"image/color"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth, screenHeight = 740, 450
	boidCount                 = 500
	viewRadius                = 13
	adjRate                   = float64(0.001)
)

var (
	green     = color.RGBA{10, 255, 50, 255}
	gameBoids [boidCount]*Boid
	boidMap   [screenWidth + 1][screenHeight + 1]int
	lock      = sync.Mutex{}
)

type Game struct {
}

func (g *Game) Update(screen *ebiten.Image) error {
	for _, boid := range gameBoids {
		screen.Set(int(boid.Position.X+1), int(boid.Position.Y), green)
		screen.Set(int(boid.Position.X-1), int(boid.Position.Y), green)
		screen.Set(int(boid.Position.X), int(boid.Position.Y+1), green)
		screen.Set(int(boid.Position.X), int(boid.Position.Y-1), green)
	}
	return nil
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

func main() {

	for i, row := range boidMap {
		for j := range row {
			boidMap[i][j] = -1
		}
	}

	for i := 0; i < boidCount; i++ {
		CreateBoid(i)
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Birdy boids in a box")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
