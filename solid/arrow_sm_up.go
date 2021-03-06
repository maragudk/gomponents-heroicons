package solid

import (
	g "github.com/maragudk/gomponents"

	h "github.com/maragudk/gomponents-heroicons"
)

func ArrowSmUp(children ...g.Node) g.Node {
	return h.Solid(g.Group(children),
		g.Raw(`<path fill-rule="evenodd" d="M5.293 9.707a1 1 0 010-1.414l4-4a1 1 0 011.414 0l4 4a1 1 0 01-1.414 1.414L11 7.414V15a1 1 0 11-2 0V7.414L6.707 9.707a1 1 0 01-1.414 0z" clip-rule="evenodd"/>`))
}
