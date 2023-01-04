package Display

var Screen [32][64]bool

func XOR(x bool, y bool) bool {
	return (x || y) && !(x && y)
}

func Display(x int, y int, n int, sprite []byte) { // Draw pixels from image buffer
	x = x & 31
	y = y & 63
	for i := 0; i < n; i++ {
		if x+i == 32 {
			break
		}
		for j := 0; j < 8; j++ {
			if y+j == 64 {
				break
			} else if sprite[i]>>(8-j)&0x01 == 0x01 {
				Screen[x+i][y+j] = XOR(Screen[x+i][y+j], true)
			} else {
				Screen[x+i][y+j] = XOR(Screen[x+i][y+j], false)
			}
		}
	}
}

func ClrScrn() {
	for i := 0; i < 32; i++ {
		for j := 0; j < 64; j++ {
			Screen[i][j] = false
		}
	}
}
