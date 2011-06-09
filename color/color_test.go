package color

import (
	"fmt"
	"testing"
)

var attributes = []Color{
	Bold("bold"),
	Underscore("underscore"),
	Dark("dark"),
	Italic("italic"),
	Blink("blink"),
}

var foreground = []Color{
	Black("black"),
	Red("red"),
	Green("green"),
	Blue("blue"),
	Magenta("magenta"),
	Cyan("cyan"),
	White("white"),
}

var background = []Color{
	BgBlack("black"),
	BgRed("red"),
	BgCyan("cyan"),
	BgWhite("white"),
}

func TestAttributes(t *testing.T) {
	fmt.Println(Underscore("Attributes"))
	fmt.Println()
	for _, a := range attributes {
		fmt.Printf("%s\n", a)
	}
	fmt.Println("")
}

func TestForeground(t *testing.T) {
	fmt.Println(Underscore("Foreground"))
	fmt.Println()
	for _, a := range foreground {
		fmt.Printf("%s\n", a)
	}
	fmt.Println("")
}

func TestBackground(t *testing.T) {
	fmt.Println(Underscore("Background"))
	fmt.Println()
	for _, a := range background {
		fmt.Printf("%s\n", a)
	}
	fmt.Println("")
}
