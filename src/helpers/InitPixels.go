package helpers

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Initialize pixel grid
func InitPixels() [NumRects]sdl.Rect {
	var rects [NumRects]sdl.Rect
	for i := 0; i < NoPixelsPerRow; i++ {
		for j := 0; j < NoPixelsPerCol; j++ {
			//log.Printf("(%d,%d) -> %d\n", i, j, NoPixelsPerRow*j+i)
			rects[NoPixelsPerRow*j+i] = sdl.Rect{
				X: int32(i * PixelWidth),
				Y: int32(j * PixelHeight),
				W: int32(PixelWidth),
				H: int32(PixelHeight),
			}
		}
	}
	return rects
}
