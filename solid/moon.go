package solid

import (
	g "github.com/maragudk/gomponents"

	h "github.com/maragudk/gomponents-heroicons"
)

func Moon(children ...g.Node) g.Node {
	return h.Solid(g.Group(children),
		g.Raw(`<path d="M17.293 13.293A8 8 0 016.707 2.707a8.001 8.001 0 1010.586 10.586z"/>`))
}
