package solid

import (
	g "github.com/maragudk/gomponents"

	h "github.com/maragudk/gomponents-heroicons"
)

func ArrowUpLeft(children ...g.Node) g.Node {
	return h.Solid(g.Group(children),
		g.Raw(`<path fill-rule="evenodd" d="M5.25 6.31v9.44a.75.75 0 01-1.5 0V4.5a.75.75 0 01.75-.75h11.25a.75.75 0 010 1.5H6.31l13.72 13.72a.75.75 0 11-1.06 1.06L5.25 6.31z" clip-rule="evenodd"/>`))
}
