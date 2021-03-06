package solid

import (
	g "github.com/maragudk/gomponents"

	h "github.com/maragudk/gomponents-heroicons"
)

func Folder(children ...g.Node) g.Node {
	return h.Solid(g.Group(children),
		g.Raw(`<path d="M2 6a2 2 0 012-2h5l2 2h5a2 2 0 012 2v6a2 2 0 01-2 2H4a2 2 0 01-2-2V6z"/>`))
}
