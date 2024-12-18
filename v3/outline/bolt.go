package outline

import (
	g "maragu.dev/gomponents"

	h "maragu.dev/gomponents-heroicons/v3"
)

func Bolt(children ...g.Node) g.Node {
	return h.Outline(g.Group(children),
		g.Raw(`<path stroke-linecap="round" stroke-linejoin="round" d="M3.75 13.5l10.5-11.25L12 10.5h8.25L9.75 21.75 12 13.5H3.75z"/>`))
}
