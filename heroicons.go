package heroicons

import (
	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
	. "github.com/maragudk/gomponents/svg"
)

func Outline(children ...g.Node) g.Node {
	// <svg fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
	return html.SVG(
		Fill("none"),
		ViewBox("0 0 24 24"),
		g.Attr("stroke-width", "1.5"),
		Stroke("currentColor"),
		html.Aria("hidden", "true"),
		g.Group(children),
	)
}

func Solid(children ...g.Node) g.Node {
	// <svg viewBox="0 0 24 24" fill="currentColor" aria-hidden="true">
	return html.SVG(
		ViewBox("0 0 24 24"),
		Fill("currentColor"),
		html.Aria("hidden", "true"),
		g.Group(children),
	)
}

func Mini(children ...g.Node) g.Node {
	// <svg viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
	return html.SVG(
		ViewBox("0 0 20 20"),
		Fill("currentColor"),
		html.Aria("hidden", "true"),
		g.Group(children),
	)
}
