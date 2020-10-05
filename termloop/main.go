package main

import tl "github.com/JoelOtter/termloop"

type Player struct {
	*tl.Entity
	prevX  int
	prevY  int
	level  *tl.BaseLevel
	screen *tl.Screen
}

type Border struct {
	*tl.Entity
	Ch rune
}

func NewBorder(x, y int, ch rune) *Border {
	border := &Border{
		Entity: tl.NewEntity(x, y, 1, 1),
		Ch:     ch,
	}
	cell := &tl.Cell{Fg: tl.ColorWhite, Ch: border.Ch}
	border.Fill(cell)

	return border
}

func (player *Player) Tick(event tl.Event) {
	player.prevX, player.prevY = player.Position()
	if event.Type == tl.EventKey { // Is it a keyboard event?
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			player.SetPosition(player.prevX+1, player.prevY)
		case tl.KeyArrowLeft:
			player.SetPosition(player.prevX-1, player.prevY)
		case tl.KeyArrowUp:
			player.SetPosition(player.prevX, player.prevY-1)
		case tl.KeyArrowDown:
			player.SetPosition(player.prevX, player.prevY+1)
		}
	} else if event.Type == tl.EventMouse {
		if event.Key == tl.MouseRelease {
			screenWidth, screenHeight := player.screen.Size()
			diffX := event.MouseX - screenWidth/2
			diffY := event.MouseY - screenHeight/2
			player.SetPosition(player.prevX+diffX, player.prevY+diffY)
		}
	}
}

func (player *Player) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*tl.Rectangle); ok {
		player.SetPosition(player.prevX, player.prevY)
	} else if _, ok := collision.(*Border); ok {
		player.SetPosition(player.prevX, player.prevY)
	}
}

func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := player.Position()
	player.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	// We need to make sure and call Draw on the underlying Entity.
	player.Entity.Draw(screen)
}

func InitBorder(level *tl.BaseLevel, width, height int) {
	// top and bottom border
	bottomY := height - 1
	for x := 0; x < width; x++ {
		topBorder := NewBorder(x, 0, '=')
		bottomBorder := NewBorder(x, bottomY, '=')
		level.AddEntity(topBorder)
		level.AddEntity(bottomBorder)
	}
	// left and right border
	rightX := width - 2
	for y := 0; y < height; y++ {
		leftBorder := NewBorder(1, y, '‖')
		rightBorder := NewBorder(rightX, y, '‖')
		level.AddEntity(leftBorder)
		level.AddEntity(rightBorder)
	}
}

func main() {
	game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: 'w',
	})

	InitBorder(level, 50, 50)

	player := Player{
		Entity: tl.NewEntity(3, 5, 1, 1),
		level:  level,
		screen: game.Screen(),
	}
	// Set the character at position (0, 0) on the entity.
	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '♂'})
	level.AddEntity(&player)

	game.Screen().SetLevel(level)
	game.Start()
}
