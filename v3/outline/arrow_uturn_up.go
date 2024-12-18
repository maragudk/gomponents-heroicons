package outline

import (
	g "maragu.dev/gomponents"

	h "maragu.dev/gomponents-heroicons/v3"
)

func ArrowUturnUp(children ...g.Node) g.Node {
	return h.Outline(g.Group(children),
		g.Raw(`<path stroke-linecap="round" stroke-linejoin="round" d="M9 9l6-6m0 0l6 6m-6-6v12a6 6 0 01-12 0v-3"/>`))
}
