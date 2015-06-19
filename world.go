package main
import (
	"time"
	"log"
)

type World struct {
	Width int32
	Height int32
	Ball Ball
	Connections map[*connection]bool
	stop chan bool
}

type Ball struct {
	X int32
	Y int32
	XDir int32
	YDir int32
}

func (this *World) updateWorld() {
	ticker := time.NewTicker(1 * time.Second / 30);
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case <- this.stop:
		return
		case <- ticker.C:
		this.calculateBallPosition()
		}
	}
}

func (this *World) stopUpdates() {
	this.stop <- true
}

func (this *World) calculateBallPosition() {
	this.Ball.X += this.Ball.XDir
	this.Ball.Y += this.Ball.YDir
log.Printf("World %s", this)
log.Printf("Ball %s", this.Ball)
	if this.Ball.X >= this.Width {
		this.Ball.XDir *= -1
		this.Ball.X = this.Width
	}
	if this.Ball.X <= 0 {
		this.Ball.XDir *= -1
		this.Ball.X = 0
	}

	if this.Ball.Y >= this.Height {
		this.Ball.YDir *= -1
		this.Ball.Y = this.Height
	}
	if this.Ball.Y <= 0 {
		this.Ball.YDir *= -1
		this.Ball.Y = 0
	}

	response := makeCreateWorldRequest(this)
	for c := range this.Connections {
		c.send <- response
	}
}