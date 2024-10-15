package heroicons

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Outline(children ...Node) Node {
	return SVG(
		Attr("viewBox", "0 0 24 24"),
		Attr("fill", "none"),
		Attr("stroke", "currentColor"),
		Aria("hidden", "true"),
		Group(children),
	)
}

func Solid(children ...Node) Node {
	return SVG(
		Attr("viewBox", "0 0 20 20"),
		Attr("fill", "currentColor"),
		Aria("hidden", "true"),
		Group(children),
	)
}
