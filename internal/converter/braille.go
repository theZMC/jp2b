package converter

const (
	// BrailleBaseRune is the Unicode code point for the first braille character (empty)
	BrailleBaseRune = '\u2800'
)

// pixelToBrailleBit maps pixel positions (x, y) in a 2x4 grid to the
// corresponding bit position in the braille pattern
// The mapping follows this pattern:
// 0 3
// 1 4
// 2 5
// 6 7
var pixelToBrailleBit = [2][4]uint{
	{0, 1, 2, 6}, // Column 0
	{3, 4, 5, 7}, // Column 1
}

// CreateBrailleChar creates a braille character from a 2x4 grid of pixels
// Pixels marked as true will be represented as dots in the braille character
func CreateBrailleChar(pixels [2][4]bool) rune {
	var value uint16 = 0

	// Calculate the value of the braille character
	for x := range 2 {
		for y := range 4 {
			if pixels[x][y] {
				bit := pixelToBrailleBit[x][y]
				value |= (1 << bit)
			}
		}
	}

	// Convert the value to a rune
	return BrailleBaseRune + rune(value)
}

// InvertBraillePattern inverts all pixels in a braille pattern
func InvertBraillePattern(pattern [2][4]bool) [2][4]bool {
	var invertedPattern [2][4]bool
	for x := range 2 {
		for y := range 4 {
			invertedPattern[x][y] = !pattern[x][y]
		}
	}
	return invertedPattern
}
