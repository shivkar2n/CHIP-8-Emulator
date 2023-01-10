package main

import (
	"fmt"
	_ "fmt"
	"os"
	"sync"

	"github.com/shivkar2n/Chip8-Emulator/CPU"
	"github.com/shivkar2n/Chip8-Emulator/Display"
	"github.com/veandco/go-sdl2/sdl"
)

// Basic Configuration {{{ //
const (
	WindowTitle    = "Chip8 Emulator"
	WindowWidth    = 960
	WindowHeight   = 480
	NoPixelsPerCol = 32
	NoPixelsPerRow = 64

	FrameRate = 60

	PixelWidth  = 15
	PixelHeight = 15
	NumRects    = 2048
)

var FgColor = [4]uint8{0xd8, 0xde, 0xe9, 0x00}
var BgColor = [4]uint8{0x2e, 0x34, 0x40, 0x00}
var rects [NumRects]sdl.Rect
var runningMutex sync.Mutex
var Instruction [2]byte
var s = new(CPU.State)

// }}} Basic Configuration //

// Functions {{{ //
func InitPixels() { // Initialize pixel grid
	for i := 0; i < NoPixelsPerRow; i++ {
		for j := 0; j < NoPixelsPerCol; j++ {
			//fmt.Printf("(%d,%d) -> %d\n", i, j, NoPixelsPerRow*j+i)

			rects[NoPixelsPerRow*j+i] = sdl.Rect{
				X: int32(i * PixelWidth),
				Y: int32(j * PixelHeight),
				W: PixelWidth,
				H: PixelHeight,
			}
		}
	}
}

// }}} Functions //

// Main Event Loop Function {{{ //
func run() int {
	// Initialize SDL {{{ //
	var window *sdl.Window
	var renderer *sdl.Renderer
	var err error

	sdl.Do(func() {
		window, err = sdl.CreateWindow(WindowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WindowWidth, WindowHeight, sdl.WINDOW_OPENGL)
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer func() {
		sdl.Do(func() {
			window.Destroy()
		})
	}()

	sdl.Do(func() {
		renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	})
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer func() {
		sdl.Do(func() {
			renderer.Destroy()
		})
	}()

	sdl.Do(func() {
		renderer.Clear()
	})

	InitPixels()
	// }}} Initialize SDL //

	s.PC = [2]byte{0x02, 0x00} // Initialize PC to beginning of rom
	s.LoadFonts()              // load fonts
	s.LoadROM()                // Load rom into memory

	// SDL Loop {{{ //
	running := true
	for running {

		//s.Print()
		Instruction = s.InstructionFetch()
		s.IncrementPC()
		s.InstructionDecode(Instruction)

		// Rendering on screen {{{ //
		sdl.Do(func() { // Initialize window
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch event.(type) {
				case *sdl.QuitEvent:
					runningMutex.Lock()
					running = false
					runningMutex.Unlock()
				}
			}
			renderer.Clear()
			renderer.SetDrawColor(0, 0, 0, 0x20)
			renderer.FillRect(&sdl.Rect{0, 0, WindowWidth, WindowHeight})
		})

		for i, rect := range rects { // Render pixels on window
			func(i int) {
				sdl.Do(func() {
					posX := i % NoPixelsPerRow
					posY := i / NoPixelsPerRow
					if Display.Screen[posX][posY] {
						renderer.SetDrawColor(FgColor[0], FgColor[1], FgColor[2], FgColor[3])
					} else {
						renderer.SetDrawColor(BgColor[0], BgColor[1], BgColor[2], BgColor[3])
					}
					renderer.DrawRect(&rect)
					renderer.FillRect(&rect)
				})
			}(i)
		}

		sdl.Do(func() {
			renderer.Present()
			sdl.Delay(100 / FrameRate)
		})
		// }}} Rendering on screen //

	}
	// }}} SDL Loop //

	return 0
}

// }}} Main Event Loop //

// Main function {{{ //
func main() {

	// os.Exit(..) must run AFTER sdl.Main(..) below; so keep track of exit
	// status manually outside the closure passed into sdl.Main(..) below
	var exitcode int
	sdl.Main(func() {
		exitcode = run()
	})
	// os.Exit(..) must run here! If run in sdl.Main(..) above, it will cause
	// premature quitting of sdl.Main(..) function; resource cleaning deferred
	// calls/closing of channels may never run
	os.Exit(exitcode)
}

// }}} Main function //
