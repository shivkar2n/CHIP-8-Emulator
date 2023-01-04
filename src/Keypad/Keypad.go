package main
//package Keyboard

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

// Map keyboard to COSMAC VIP
var keypadMap = map[string]string{
	"1": "1",
	"2": "2",
	"3": "3",
	"4": "c",
	"q": "4",
	"w": "5",
	"e": "6",
	"r": "d",
	"a": "7",
	"s": "8",
	"d": "9",
	"f": "e",
	"z": "a",
	"x": "o",
	"c": "b",
	"v": "f",
}

func run() (err error) {
	var window *sdl.Window

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return
	}
	defer sdl.Quit()

	window, err = sdl.CreateWindow("Input", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return
	}
	defer window.Destroy()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {

			case *sdl.QuitEvent:
				running = false

			case *sdl.KeyboardEvent:
				keyCode := t.Keysym.Sym
				keys := ""

				switch t.Keysym.Mod {
				case sdl.KMOD_RALT: // Close window on right-alt press
					keys += "Right Alt"
					running = false
				}
				if keyCode < 10000 {
					if keys != "" {
						keys += " + "
					}
					if t.Repeat > 0 {
						keys += keypadMap[string(keyCode)] + " repeating"
					} else {
						if t.State == sdl.RELEASED {
							keys += keypadMap[string(keyCode)] + " released"
						} else if t.State == sdl.PRESSED {
							keys += keypadMap[string(keyCode)] + " pressed"
						}
					}
				}

				if keys != "" {
					fmt.Println(keys)
				}
			}
		}

		sdl.Delay(16)
	}

	return
}


//func Keyboard() {
func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}
