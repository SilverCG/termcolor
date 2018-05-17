package termcolor

import (
	"encoding/hex"
	"fmt"
)

var (
	termBefore = []byte("\033[")
	termAfter  = []byte("m")
	termReset  = []byte("\033[0;00m")
)

// Some predefined colors from the Go Brand Book

// GopherBlue #00ADD8
var GopherBlue = GetColor([]byte{1, 173, 216}, true, true)

// InfoColor same as GopherBlue
var InfoColor = GopherBlue

// Aqua #00A29C
var Aqua = GetColor([]byte{0, 162, 156}, true, true)

// DebugColor Shade of Light Blue #D5EFF6
var DebugColor = GetColor([]byte{213, 239, 246}, true, true)

// ErrorColor Fuchsia #CE3262
var ErrorColor = GetColor([]byte{206, 48, 98}, true, true)

// WarnColor Yellow #FDDD00
var WarnColor = GetColor([]byte{253, 221, 0}, true, true)

// convertColor to nearest ANSI color
func convertColor(c uint8) uint16 {
	return (uint16(c) * 5) / 255
}
func GetColorFromHex(hexRaw string, foreground bool, trueColor bool) []byte {
	rgbCodes := []byte{}
	hexStr := ""
	if len(hexRaw) == 7 {
		hexStr = hexRaw[1:]
	} else if len(hexRaw) == 6 {
		hexStr = hexRaw
	} else {
		return []byte{}
	}
	if len(hexStr) == 6 {
		rgbCodes, _ = hex.DecodeString(hexStr)
	}
	return GetColor(rgbCodes, foreground, trueColor)
}

// GetColor returns []byte of ANSI code
func GetColor(rgb []byte, foreground bool, trueColor bool) []byte {
	if !trueColor {
		if len(rgb) != 3 {
			return []byte{}
		}
		// convert RGB to ANSI color code.
		r6 := 36 * convertColor(rgb[0])
		g6 := 6 * convertColor(rgb[1])
		b6 := convertColor(rgb[2])
		ansi := r6 + g6 + b6 + 16

		if foreground {
			return []byte(fmt.Sprintf("38;5;%d", ansi))
		}
		return []byte(fmt.Sprintf("48;5;%d", ansi))
	}
	if foreground {
		return []byte(fmt.Sprintf("38;2;%d;%d;%d", rgb[0], rgb[1], rgb[2]))
	}
	return []byte(fmt.Sprintf("48;2;%d;%d;%d", rgb[0], rgb[1], rgb[2]))
}

// ColorBefore is used to insert the start of the color.
// it can be used on it's own for more custom control
func ColorBefore(color []byte) string {
	return fmt.Sprintf("%s%s%s", termBefore, color, termAfter)
}

// ColorAfter is used to insert the reset at the end of a color output.
// it can be used on it's own with ColorBefore for more custom control
func ColorAfter() string {
	return string(termReset)
}

// Color used to create the start of the color, insert the body, and reset the color at the end.
func Color(ansi []byte, msg ...interface{}) []interface{} {
	output := append([]interface{}{ColorBefore(ansi)}, msg...)
	output = append(output, ColorAfter())
	return output
}

func valIn(key []byte, val [][]byte) bool {
	for _, v := range val {
		if string(v) == string(key) {
			return true
		}
	}
	return false
}

func ColorANSICheck() {
	var outputArray [][]byte
	for r := 0; r < 256; r++ {
		for g := 0; g < 256; g++ {
			for b := 0; b < 256; b++ {
				check := GetColor([]byte{uint8(r), uint8(g), uint8(b)}, true, false)
				fmt.Print(Color(check, "█")...)
				if !valIn(check, outputArray) {
					outputArray = append(outputArray, check)
				}
			}
		}
	}
	for _, c := range outputArray {
		fmt.Print(Color(c, "█")...)
	}
}
