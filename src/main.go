package main

import (
	_ "fmt"
	"log"
	"os"
	"sync"

	"github.com/shivkar2n/Chip8-Emulator/CPU"
	"github.com/shivkar2n/Chip8-Emulator/Display"
	"github.com/veandco/go-sdl2/mix"
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
// Initialize pixel grid
func InitPixels() {
	for i := 0; i < NoPixelsPerRow; i++ {
		for j := 0; j < NoPixelsPerCol; j++ {
			//log.Printf("(%d,%d) -> %d\n", i, j, NoPixelsPerRow*j+i)

			rects[NoPixelsPerRow*j+i] = sdl.Rect{
				X: int32(i * PixelWidth),
				Y: int32(j * PixelHeight),
				W: PixelWidth,
				H: PixelHeight,
			}
		}
	}
}

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

// }}} Functions //

// Main Event Loop Function {{{ //
func run() int {
	// Initialize SDL {{{ //
	var window *sdl.Window
	var renderer *sdl.Renderer
	var err error
	var running bool

	// Initialize audio {{{ //
	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
		log.Println(err)
	}
	defer sdl.Quit()

	if err := mix.Init(mix.INIT_MP3); err != nil {
		log.Println(err)
	}
	defer mix.Quit()

	if err := mix.OpenAudio(22050, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		log.Println(err)
	}
	defer mix.CloseAudio()
	// }}} Initialize audio //

	window, err = sdl.CreateWindow(WindowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WindowWidth, WindowHeight, sdl.WINDOW_OPENGL)
	if err != nil {
		log.Fatalln("Failed to create window: ", err)
		return 1
	}
	defer func() {
		window.Destroy()
	}()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatalln("Failed to create renderer: ", err)
		return 2
	}
	defer func() {
		renderer.Destroy()
	}()

	renderer.Clear()

	InitPixels()
	// }}} Initialize SDL //

	s.PC = [2]byte{0x02, 0x00} // Initialize PC to beginning of rom
	s.LoadFonts()              // load fonts
	s.LoadROM()                // Load rom into memory

	// SDL Loop {{{ //
	running = true
	for running {

		// Fetch and decode instructions {{{ //
		//s.Print()
		Instruction = s.InstructionFetch()
		s.IncrementPC()
		s.InstructionDecode(Instruction, sdl.GetKeyboardState())
		// }}} Fetch and decode instructions //

		// Rendering on screen {{{ //
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				runningMutex.Lock()
				running = false
				runningMutex.Unlock()
			}
		}
		renderer.Clear()

		// Rendering window
		renderer.SetDrawColor(BgColor[0], BgColor[1], BgColor[2], BgColor[3])
		renderer.FillRect(&sdl.Rect{0, 0, WindowWidth, WindowHeight})

		// Rendering pixels on background
		FgRects := GetFgRects(rects)
		renderer.SetDrawColor(FgColor[0], FgColor[1], FgColor[2], FgColor[3])
		renderer.DrawRects(FgRects)
		renderer.FillRects(FgRects)

		renderer.Present()
		renderer.RenderSetVSync(true)
		sdl.Delay(0)
		// }}} Rendering on screen //

		// Decrement sound, delay timer {{{ //
		s.DecrementDelayTimer()
		s.DecrementSoundTimer()

		// Play sound
		music, _ := mix.LoadMUS("../assets/sounds/beep.mp3")
		music.Play(1)
		if s.PlaySound() {
			sdl.Delay(0)
			music.Free()
		}
		// }}} Decrement sound, display timer //
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
