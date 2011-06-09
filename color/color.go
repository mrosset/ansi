package color

import (
	"fmt"
)

const escape = "\x1b[00;0;%vm%s\x1b[m"

// Text Attributes
const (
	RESET = iota
	BOLD
	DARK
	ITALIC
	UNDERSCORE
	BLINK
)

// Foreground colors
const (
	BLACK = iota + 30
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)

// Background Colors 
const (
	BG_BLACK = iota + 40
	BG_RED   = 41
	BG_CYAN  = 46
	BG_WHITE = 47
)


func ansi(code int, s string) string {
	return fmt.Sprintf(escape, code, s)
}

type Color interface {
	Format(fmt.State, int)
}

// Test Attribute Types
type Reset string

func (r Reset) Format(s fmt.State, c int) {
	fmt.Fprintf(s, escape, RESET, string(r))
}

type Bold string

func (r Bold) Format(s fmt.State, c int) {
	fmt.Fprintf(s, escape, BOLD, string(r))
}

type Dark string

func (r Dark) Format(s fmt.State, c int) {
	fmt.Fprintf(s, escape, DARK, string(r))
}

type Italic string

func (r Italic) Format(s fmt.State, c int) {
	fmt.Fprintf(s, escape, ITALIC, string(r))
}

type Underscore string

func (r Underscore) Format(s fmt.State, c int) {
	fmt.Fprintf(s, escape, UNDERSCORE, string(r))
}

type Blink string

func (r Blink) Format(s fmt.State, c int) {
	fmt.Fprintf(s, escape, BLINK, string(r))
}

type Black string

func (r Black) Format(s fmt.State, c int) {
	fmt.Fprintf(s, escape, BLACK, string(r))
}


// Foreground Types
type Red string

func (r Red) Format(s fmt.State, c int) {
	fmt.Fprintf(s, escape, RED, string(r))
}

type Green string

func (r Green) Format(s fmt.State, c int) {
	fmt.Fprintf(s, escape, GREEN, string(r))
}

type Yellow string

func (r Yellow) Format(s fmt.State, c int) {
	fmt.Fprintf(s, escape, YELLOW, string(r))
}

type Blue string

func (r Blue) Format(s fmt.State, c int) {
	fmt.Fprintf(s, escape, BLUE, string(r))
}

type Magenta string

func (r Magenta) Format(s fmt.State, c int) {
	fmt.Fprintf(s, escape, MAGENTA, string(r))
}

type Cyan string

func (r Cyan) Format(s fmt.State, c int) {
	fmt.Fprintf(s, escape, CYAN, string(r))
}

type White string

func (r White) Format(s fmt.State, c int) {
	fmt.Fprintf(s, escape, WHITE, string(r))
}


// Background Types
type BgBlack string

func (r BgBlack) Format(s fmt.State, c int) {
	fmt.Fprintf(s, escape, BG_BLACK, string(r))
}

type BgRed string

func (r BgRed) Format(s fmt.State, c int) {
	fmt.Fprintf(s, escape, BG_RED, string(r))
}

type BgWhite string

func (r BgWhite) Format(s fmt.State, c int) {
	fmt.Fprintf(s, escape, BG_WHITE, string(r))
}

type BgCyan string

func (r BgCyan) Format(s fmt.State, c int) {
	fmt.Fprintf(s, escape, BG_CYAN, string(r))
}
