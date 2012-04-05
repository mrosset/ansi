package color

import (
	"fmt"
	"testing"
)

var attributes = []*Escape{
	Bold("bold"),
    Dark("dark"),
    Italic("italic"),
	Underline("underscore"),
	Blink("blink"),
	Reverse("dark"),
	Concealed("italic"),
}

var foreground = []*Escape{
	Black("black"),
	Red("red"),
	Green("green"),
	Blue("blue"),
	Magenta("magenta"),
	Cyan("cyan"),
	White("white"),
}

var background = []*Escape{
	BgGrey("grey"),
	BgRed("red"),
	BgCyan("green"),
    BgYellow("yellow"),
    BgBlue("blue"),
    BgMagenta("magenta"),
    BgCyan("cyan"),
	BgWhite("white"),
}

var composed = []*Escape{
    Red("red"),
    Bold(Red("bold red")),
    BgGreen(Blue("blue on white background")),
    Underline(Blink(Magenta("magenta blinking underlined text"))),
}

func TestAttributes(t *testing.T) {
	fmt.Println(Underline("Attributes"))
	fmt.Println()
	for _, a := range attributes {
		fmt.Printf("%s\n", a)
	}
	fmt.Println("----------------")
}

func TestForeground(t *testing.T) {
	fmt.Println(Underline("Foreground"))
	fmt.Println()
	for _, a := range foreground {
		fmt.Printf("%s\n", a)
	}
	fmt.Println("----------------")
}

func TestBackground(t *testing.T) {
	fmt.Println(Underline("Background"))
	fmt.Println()
	for _, a := range background {
		fmt.Printf("%s\n", a)
	}
	fmt.Println("----------------")
}

func TestComposed(t *testing.T) {
	fmt.Println(Underline("Composed"))
	fmt.Println()
	for _, a := range composed {
		fmt.Printf("%s\n", a)
	}
	fmt.Println("----------------")
}

