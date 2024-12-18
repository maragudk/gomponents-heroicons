package solid

import (
	g "maragu.dev/gomponents"

	h "maragu.dev/gomponents-heroicons/v3"
)

func ArrowUturnRight(children ...g.Node) g.Node {
	return h.Solid(g.Group(children),
		g.Raw(`<path fill-rule="evenodd" d="M14.47 2.47a.75.75 0 011.06 0l6 6a.75.75 0 010 1.06l-6 6a.75.75 0 11-1.06-1.06l4.72-4.72H9a5.25 5.25 0 100 10.5h3a.75.75 0 010 1.5H9a6.75 6.75 0 010-13.5h10.19l-4.72-4.72a.75.75 0 010-1.06z" clip-rule="evenodd"/>`))
}
