// the functions in this file allow to directly print ANSI coloured text to the
// standard output
//
// Colours are passed to the functions by strings that define a colour code.
//
// Colour code syntax:
// 	r ... red
//	g ... green
// 	y ... yellow
//	b ... blue
//	x ... black
//	m ... magenta
// 	c ... cyan
//	w ... white
//	d ... default
//	+ ... bold
//	* ... italic
//	~ ... reverse
// 	_ ... underlined
//	# ... blink
//	? ... concealed
//
//	background coloring is affected by using upper case letters
//
// examples:
//	+r ... bold red
//	rY ... red text on yellow background
package color

import "fmt"

// --- EXPORTED ---

func Printf(color string, format string, a ...interface{}) {
	// TODO implement
}

// --- INTERNAL ONLY ---

// colours a text based on the given colour code
func colourText(colorCode string, text string) string{
	var coloredText string

	if len(colorCode) < 1 {
		return text // nothing to do
	}

	switch colorCode[0] {
	// foreground
	case 'r': coloredText = fmt.Sprintf("%s", Red(text))
	case 'g': coloredText = fmt.Sprintf("%s", Green(text))
	case 'y': coloredText = fmt.Sprintf("%s", Yellow(text))
	case 'b': coloredText = fmt.Sprintf("%s", Blue(text))
	case 'x': coloredText = fmt.Sprintf("%s", Black(text))
	case 'm': coloredText = fmt.Sprintf("%s", Magenta(text))
	case 'c': coloredText = fmt.Sprintf("%s", Cyan(text))
	case 'w': coloredText = fmt.Sprintf("%s", White(text))
	// background
	case 'R': coloredText = fmt.Sprintf("%s", BgRed(text))
	case 'G': coloredText = fmt.Sprintf("%s", BgGreen(text))
	case 'Y': coloredText = fmt.Sprintf("%s", BgYellow(text))
	case 'B': coloredText = fmt.Sprintf("%s", BgBlue(text))
	// case 'X': coloredText = fmt.Sprintf("%s", BgBlack(text)) -> not implemented
	case 'M': coloredText = fmt.Sprintf("%s", BgMagenta(text))
	case 'C': coloredText = fmt.Sprintf("%s", BgCyan(text))
	case 'W': coloredText = fmt.Sprintf("%s", BgWhite(text))
	// specials
	case '+': coloredText = fmt.Sprintf("%s", Bold(text))
	case '*': coloredText = fmt.Sprintf("%s", Italic(text))
	case '~': coloredText = fmt.Sprintf("%s", Reverse(text))
	case '_': coloredText = fmt.Sprintf("%s", Underline(text))
	case '#': coloredText = fmt.Sprintf("%s", Blink(text))
	case '?': coloredText = fmt.Sprintf("%s", Concealed(text))
	// default -> nothing
	default: coloredText = text
	}

	return colourText(colorCode[1:], coloredText)
}
