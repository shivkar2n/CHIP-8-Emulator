package Display

var Screen [64][32]bool

func XOR(x bool, y bool) bool {
	return (x || y) && !(x && y)
}

func Display(x int, y int, n int, sprite []byte) { // Draw pixels from image buffer
	x = x & 63
	y = y & 31
	for i := 0; i < n; i++ {
		if y+i == 32 {
			break
		}
		for j := 0; j < 8; j++ {
			if x+j == 64 {
				break
			} else if sprite[i]>>(7-j)&0x01 == 0x01 {
				Screen[x+j][y+i] = XOR(Screen[x+j][y+i], true)
			} else {
				Screen[x+j][y+i] = XOR(Screen[x+j][y+i], false)
			}
		}
	}
}

func ClrScrn() { // Clear all pixels on screen
	for i := 0; i < 64; i++ {
		for j := 0; j < 32; j++ {
			Screen[i][j] = false
		}
	}
}
