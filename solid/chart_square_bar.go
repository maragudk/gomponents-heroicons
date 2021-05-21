package solid

import (
	g "github.com/maragudk/gomponents"

	h "github.com/maragudk/gomponents-heroicons"
)

func ChartSquareBar(children ...g.Node) g.Node {
	return h.Solid(g.Group(children),
		g.Raw(`<path fill-rule="evenodd" d="M5 3a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2V5a2 2 0 00-2-2H5zm9 4a1 1 0 10-2 0v6a1 1 0 102 0V7zm-3 2a1 1 0 10-2 0v4a1 1 0 102 0V9zm-3 3a1 1 0 10-2 0v1a1 1 0 102 0v-1z" clip-rule="evenodd"/>`))
}
