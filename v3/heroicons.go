package heroicons

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Outline(children ...Node) Node {
	// <svg fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
	return SVG(
		Attr("fill", "none"),
		Attr("viewBox", "0 0 24 24"),
		Attr("stroke-width", "1.5"),
		Attr("stroke", "currentColor"),
		Aria("hidden", "true"),
		Group(children),
	)
}

func Solid(children ...Node) Node {
	// <svg viewBox="0 0 24 24" fill="currentColor" aria-hidden="true">
	return SVG(
		Attr("viewBox", "0 0 24 24"),
		Attr("fill", "currentColor"),
		Aria("hidden", "true"),
		Group(children),
	)
}

func Mini(children ...Node) Node {
	// <svg viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
	return SVG(
		Attr("viewBox", "0 0 20 20"),
		Attr("fill", "currentColor"),
		Aria("hidden", "true"),
		Group(children),
	)
}
