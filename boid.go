package main

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	Position Vector2D
	Velocity Vector2D
	Id       int
}

func CreateBoid(id int) {
	b := Boid{
		Position: Vector2D{rand.Float64() * screenWidth, rand.Float64() * screenHeight},
		Velocity: Vector2D{(rand.Float64() * 2) - 1, (rand.Float64() * 2) - 1},
		Id:       id,
	}
	gameBoids[id] = &b
	boidMap[int(b.Position.X)][int(b.Position.Y)] = b.Id
	go b.Start()
}

func (b *Boid) Start() {
	for {
		b.MoveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func (b *Boid) MoveOne() {
	b.Velocity = b.Velocity.Add(b.calcAccel().Limit(-1, 1))
	boidMap[int(b.Position.X)][int(b.Position.Y)] = -1
	b.Position = b.Position.Add(b.Velocity)
	boidMap[int(b.Position.X)][int(b.Position.Y)] = b.Id
	next := b.Position.Add(b.Velocity)
	if next.X >= screenWidth || next.X < 0 {
		b.Velocity = Vector2D{-b.Velocity.X, b.Velocity.Y}
	}
	if next.Y >= screenHeight || next.Y < 0 {
		b.Velocity = Vector2D{b.Velocity.X, -b.Velocity.Y}
	}
}

func (b Boid) calcAccel() Vector2D {

	upper, lower := b.Position.AddV(viewRadius), b.Position.AddV(-viewRadius)
	avgVelocity := Vector2D{0, 0}
	count := 0.0

	for i := math.Max(lower.X, 0); i <= math.Min(upper.X, screenWidth); i++ {
		for j := math.Max(lower.Y, 0); j <= math.Min(upper.Y, screenHeight); j++ {
			if otherBoidId := boidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.Id {
				if dist := gameBoids[otherBoidId].Position.Distance(b.Position); dist < viewRadius {
					count++
					avgVelocity = avgVelocity.Add(gameBoids[otherBoidId].Velocity)
				}
			}
		}
	}

	accel := Vector2D{0, 0}
	if count > 0 {
		avgVelocity = avgVelocity.DivideV(count)
		accel = avgVelocity.Subtract(b.Velocity).MultiplyV(adjRate)
	}
	return accel
}
