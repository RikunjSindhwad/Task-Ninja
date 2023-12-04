package visuals

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	colorRed     = 31
	colorGreen   = 32
	colorYellow  = 33
	colorBlue    = 34
	colorMagenta = 35
	colorCyan    = 36
	colorWhite   = 37
)

var colors = []int{colorRed, colorGreen, colorYellow, colorBlue, colorMagenta, colorCyan, colorWhite}

func PrintRandomColor(s string, specifiedColor ...int) string {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	var colorCode int
	if len(specifiedColor) > 0 {
		colorCode = specifiedColor[0]
	} else {
		colorCode = colors[rand.Intn(len(colors))]
	}

	return fmt.Sprintf("\033[1;%dm%s\033[0m", colorCode, s)
}
