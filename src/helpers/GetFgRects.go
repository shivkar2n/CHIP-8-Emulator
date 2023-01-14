package helpers

import (
	"github.com/shivkar2n/Chip8-Emulator/Display"
	"github.com/veandco/go-sdl2/sdl"
)

func GetFgRects(rects [NumRects]sdl.Rect) []sdl.Rect {
	FgRects := make([]sdl.Rect, NumRects, NumRects)
	for i, rect := range rects {
		posX := i % NoPixelsPerRow
		posY := i / NoPixelsPerRow
		if Display.Screen[posX][posY] {
			FgRects = append(FgRects, rect)
		}
	}
	return FgRects
}
