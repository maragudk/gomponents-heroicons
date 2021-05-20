package heroicons

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	"github.com/maragudk/gomponents/svg"
)

func Outline(children ...g.Node) g.Node {
	return SVG(
		svg.ViewBox("0 0 24 24"),
		svg.Fill("none"),
		svg.Stroke("currentColor"),
		Aria("hidden", "true"),
		g.Group(children),
	)
}

func Solid(children ...g.Node) g.Node {
	return SVG(
		svg.ViewBox("0 0 20 20"),
		svg.Fill("currentColor"),
		Aria("hidden", "true"),
		g.Group(children),
	)
}
