package helpers

// Map keyboard to COSMAC VIP
var KeypadMap = map[byte]byte{
	'1': 0x01,
	'2': 0x02,
	'3': 0x03,
	'4': 0x0c,
	'q': 0x04,
	'w': 0x05,
	'e': 0x06,
	'r': 0x0d,
	'a': 0x07,
	's': 0x08,
	'd': 0x09,
	'f': 0x0e,
	'z': 0x0a,
	'x': 0x00,
	'c': 0x0b,
	'v': 0x0f,
}

//Helper function that returns pressed key
// If no key pressed returns 0xff
func KeyPressed(keyboardState []uint8) byte {
	var i uint8
	for mapkey, key := range KeypadMap {
		if mapkey == '1' || mapkey == '2' || mapkey == '3' || mapkey == '4' {
			i = uint8(mapkey) - uint8('1') + 30
		} else {
			i = uint8(mapkey) - uint8('a') + 4
		}
		state := keyboardState[i]
		if state == 1 {
			//fmt.Printf("Key pressed: %c\n", key)
			return byte(key)
		}
	}
	return byte(0xff)
}
