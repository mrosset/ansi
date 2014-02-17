package color

import "testing"
import "fmt"

// colour parsing for foreground
func TestForegroundRed(t *testing.T) {
	c := colourText("r", "test")
	v := fmt.Sprintf("%s", Red("test"))

	fmt.Printf("%s vs %s\n", c, v)

	if c != v {
		t.Fail()
	}
}

func TestForegroundGreen(t *testing.T) {
	c := colourText("g", "test")
	v := fmt.Sprintf("%s", Green("test"))

	fmt.Printf("%s vs %s\n", c, v)

	if c != v {
		t.Fail()
	}
}

func TestForegroundYellow(t *testing.T) {
	c := colourText("y", "test")
	v := fmt.Sprintf("%s", Yellow("test"))

	fmt.Printf("%s vs %s\n", c, v)

	if c != v {
		t.Fail()
	}
}

func TestForegroundBlue(t *testing.T) {
	c := colourText("b", "test")
	v := fmt.Sprintf("%s", Blue("test"))

	fmt.Printf("%s vs %s\n", c, v)

	if c != v {
		t.Fail()
	}
}

func TestForegroundBlack(t *testing.T) {
	c := colourText("x", "test")
	v := fmt.Sprintf("%s", Black("test"))

	fmt.Printf("%s vs %s\n", c, v)

	if c != v {
		t.Fail()
	}
}

func TestForegroundMagenta(t *testing.T) {
	c := colourText("m", "test")
	v := fmt.Sprintf("%s", Magenta("test"))

	fmt.Printf("%s vs %s\n", c, v)

	if c != v {
		t.Fail()
	}
}

func TestForegroundCyan(t *testing.T) {
	c := colourText("c", "test")
	v := fmt.Sprintf("%s", Cyan("test"))

	fmt.Printf("%s vs %s\n", c, v)

	if c != v {
		t.Fail()
	}
}

func TestForegroundWhite(t *testing.T) {
	c := colourText("w", "test")
	v := fmt.Sprintf("%s", White("test"))

	fmt.Printf("%s vs %s\n", c, v)

	if c != v {
		t.Fail()
	}
}

// colour parsing for background
func TestBackgroundRed(t *testing.T) {
	c := colourText("R", "test")
	v := fmt.Sprintf("%s", BgRed("test"))

	fmt.Printf("%s vs %s\n", c, v)

	if c != v {
		t.Fail()
	}
}

func TestBackgroundGreen(t *testing.T) {
	c := colourText("G", "test")
	v := fmt.Sprintf("%s", BgGreen("test"))

	fmt.Printf("%s vs %s\n", c, v)

	if c != v {
		t.Fail()
	}
}

func TestBackgroundYellow(t *testing.T) {
	c := colourText("Y", "test")
	v := fmt.Sprintf("%s", BgYellow("test"))

	fmt.Printf("%s vs %s\n", c, v)

	if c != v {
		t.Fail()
	}
}

func TestBackgroundBlue(t *testing.T) {
	c := colourText("B", "test")
	v := fmt.Sprintf("%s", BgBlue("test"))

	fmt.Printf("%s vs %s\n", c, v)

	if c != v {
		t.Fail()
	}
}

func TestBackgroundMagenta(t *testing.T) {
	c := colourText("M", "test")
	v := fmt.Sprintf("%s", BgMagenta("test"))

	fmt.Printf("%s vs %s\n", c, v)

	if c != v {
		t.Fail()
	}
}

func TestBackgroundCyan(t *testing.T) {
	c := colourText("C", "test")
	v := fmt.Sprintf("%s", BgCyan("test"))

	fmt.Printf("%s vs %s\n", c, v)

	if c != v {
		t.Fail()
	}
}

func TestBackgroundWhite(t *testing.T) {
	c := colourText("W", "test")
	v := fmt.Sprintf("%s", BgWhite("test"))

	fmt.Printf("%s vs %s\n", c, v)

	if c != v {
		t.Fail()
	}
}

// some multiple parsings
func TestBoldMagenta(t *testing.T) {
	c := colourText("+m", "test")
	v := fmt.Sprintf("%s", Bold(Magenta("test")))

	fmt.Printf("%s vs %s\n", c, v)
}

func TestUnderlinedRed(t *testing.T) {
	c := colourText("_r", "test")
	v := fmt.Sprintf("%s", Underline(Red("test")))

	fmt.Printf("%s vs %s\n", c, v)
}

func TestGreenOnYellow(t *testing.T) {
	c := colourText("gY", "test")
	v := fmt.Sprintf("%s", Green(BgYellow("test")))

	fmt.Printf("%s vs %s\n", c, v)
}

func TestConcealed(t *testing.T) {
	c := colourText("?", "test")
	v := fmt.Sprintf("%s", Concealed("test"))

	fmt.Printf("%s vs %s\n", c, v)
	if c != v {
		t.Fail()
	}
}

func TestReverse(t *testing.T) {
	c := colourText("~", "test")
	v := fmt.Sprintf("%s", Reverse("test"))

	fmt.Printf("%s vs %s\n", c, v)
	if c != v {
		t.Fail()
	}
}

// printing
func TestPrintf(t *testing.T) {
	Printf("*r", "%s\n", "printf italic red")
}

func TestPrint(t *testing.T) {
	Print("+y", "print bold yellow")
}
