package svg_test

import (
	"fmt"
	"testing"

	g "github.com/christophersw/gomponents-htmx"
	"github.com/christophersw/gomponents-htmx/internal/assert"
	. "github.com/christophersw/gomponents-htmx/svg"
)

func TestSimpleAttributes(t *testing.T) {
	cases := map[string]func(string) g.Node{
		"clip-rule":    ClipRule,
		"d":            D,
		"fill":         Fill,
		"fill-rule":    FillRule,
		"stroke":       Stroke,
		"stroke-width": StrokeWidth,
		"viewBox":      ViewBox,
	}

	for name, fn := range cases {
		t.Run(fmt.Sprintf(`should output %v="hat"`, name), func(t *testing.T) {
			n := g.El("element", fn("hat"))
			assert.Equal(t, fmt.Sprintf(`<element %v="hat"></element>`, name), n)
		})
	}
}
