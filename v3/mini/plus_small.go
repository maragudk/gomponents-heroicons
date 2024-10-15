package mini

import (
	g "maragu.dev/gomponents"

	h "maragu.dev/gomponents-heroicons/v3"
)

func PlusSmall(children ...g.Node) g.Node {
	return h.Mini(g.Group(children),
		g.Raw(`<path d="M10.75 6.75a.75.75 0 00-1.5 0v3.5h-2.5a.75.75 0 000 1.5h2.5v3.5a.75.75 0 001.5 0v-2.5h2.5a.75.75 0 000-1.5h-2.5v-2.5z"/>`))
}
