package outline

import (
	g "maragu.dev/gomponents"

	h "maragu.dev/gomponents-heroicons/v3"
)

func ChevronDoubleDown(children ...g.Node) g.Node {
	return h.Outline(g.Group(children),
		g.Raw(`<path stroke-linecap="round" stroke-linejoin="round" d="M19.5 5.25l-7.5 7.5-7.5-7.5m15 6l-7.5 7.5-7.5-7.5"/>`))
}
