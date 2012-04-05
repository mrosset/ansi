package color

import (
	"fmt"
)

// The old escape sequence was FUBAR
const escapeSeq = "\x1b[%vm%s\x1b[0m"

// Text Attributes
const (
	BOLD = iota + 1
	DARK
	ITALIC
	UNDERLINE
	BLINK
    _
    REVERSE
    CONCEALED
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
	BG_GREY = iota + 40
	BG_RED
    BG_GREEN
    BG_YELLOW
    BG_BLUE
    BG_MAGENTA
    BG_CYAN
    BG_WHITE
)

type escaper interface {
    fmt.Formatter
    fmt.Stringer
}

type Escape struct {
    code int
    escapeStr escaper
}

func newColor(code int, escapeOrStr interface{}) *Escape {
    switch s := escapeOrStr.(type) {
    case *Escape: return &Escape{code, s}
    case string: return &Escape{code, rawString(s)}
    }

    panic(fmt.Sprintf("bad value: %v", escapeOrStr))
}

func (clr *Escape) Format(s fmt.State, c rune) {
    fmt.Fprintf(s, clr.String())
}

func (clr *Escape) String() string {
    return fmt.Sprintf(escapeSeq, clr.code, fmt.Sprintf("%s", clr.escapeStr))
}

type rawString string

func (rs rawString) Format(s fmt.State, c rune) {
    fmt.Fprintf(s, rs.String())
}

func (rs rawString) String() string {
    return string(rs)
}

// Terminal color escape sequences
func Black(cs interface{}) *Escape {
    return newColor(BLACK, cs)
}

func Red(cs interface{}) *Escape {
    return newColor(RED, cs)
}

func Green(cs interface{}) *Escape {
    return newColor(GREEN, cs)
}

func Yellow(cs interface{}) *Escape {
    return newColor(YELLOW, cs)
}

func Blue(cs interface{}) *Escape {
    return newColor(BLUE, cs)
}

func Magenta(cs interface{}) *Escape {
    return newColor(MAGENTA, cs)
}

func Cyan(cs interface{}) *Escape {
    return newColor(CYAN, cs)
}

func White(cs interface{}) *Escape {
    return newColor(WHITE, cs)
}

// Terminal background color escape sequences
func BgGrey(cs interface{}) *Escape {
    return newColor(BG_GREY, cs)
}

func BgRed(cs interface{}) *Escape {
    return newColor(BG_RED, cs)
}

func BgGreen(cs interface{}) *Escape {
    return newColor(BG_GREEN, cs)
}

func BgYellow(cs interface{}) *Escape {
    return newColor(BG_YELLOW, cs)
}

func BgBlue(cs interface{}) *Escape {
    return newColor(BG_BLUE, cs)
}

func BgMagenta(cs interface{}) *Escape {
    return newColor(BG_MAGENTA, cs)
}

func BgCyan(cs interface{}) *Escape {
    return newColor(BG_CYAN, cs)
}

func BgWhite(cs interface{}) *Escape {
    return newColor(BG_WHITE, cs)
}

// Terminal attribute escape sequences
func Bold(cs interface{}) *Escape {
    return newColor(BOLD, cs)
}

func Dark(cs interface{}) *Escape {
    return newColor(DARK, cs)
}

func Italic(cs interface{}) *Escape {
    return newColor(ITALIC, cs)
}

func Underline(cs interface{}) *Escape {
    return newColor(UNDERLINE, cs)
}

func Blink(cs interface{}) *Escape {
    return newColor(BLINK, cs)
}

func Reverse(cs interface{}) *Escape {
    return newColor(REVERSE, cs)
}

func Concealed(cs interface{}) *Escape {
    return newColor(CONCEALED, cs)
}

