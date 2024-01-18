package color

import (
	"fmt"
	"strings"
	"testing"
)

var pt = fmt.Printf

func TestGnome(t *testing.T) {
	buf := new(strings.Builder)
	buf.WriteString("[")
	for i, c := range Ansi16 {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString("'")
		r, g, b, _ := c.RGBA()
		buf.WriteString(fmt.Sprintf("#%x%x%x",
			r*0xff/0xffff,
			g*0xff/0xffff,
			b*0xff/0xffff,
		))
		buf.WriteString("'")
	}
	buf.WriteString("]")
	pt("%s\n", buf.String())
}