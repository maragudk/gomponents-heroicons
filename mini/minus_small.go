package mini

import (
	g "github.com/maragudk/gomponents"

	h "github.com/maragudk/gomponents-heroicons"
)

func MinusSmall(children ...g.Node) g.Node {
	return h.Mini(g.Group(children),
		g.Raw(`<path d="M6.75 9.25a.75.75 0 000 1.5h6.5a.75.75 0 000-1.5h-6.5z"/>`))
}
