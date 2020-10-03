package design_pattern

import "fmt"

// design patterns implemented in golang

// bridge
type pen interface {
	Draw(color) string
}

type color interface {
	GetColor() string
}

type brushPen struct{}

func (brushPen) Draw(c color) string {
	return fmt.Sprintf("draw %s with brush", c.GetColor())
}

type pencilPen struct{}

func (pencilPen) Draw(c color) string {
	return fmt.Sprintf("draw %s with pencil", c.GetColor())
}

type red struct{}

func (red) GetColor() string {
	return "red color"
}

type blue struct{}

func (blue) GetColor() string {
	return "blue color"
}
