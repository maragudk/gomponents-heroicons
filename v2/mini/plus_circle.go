package mini

import (
	g "github.com/maragudk/gomponents"

	h "github.com/maragudk/gomponents-heroicons/v2"
)

func PlusCircle(children ...g.Node) g.Node {
	return h.Mini(g.Group(children),
		g.Raw(`<path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm.75-11.25a.75.75 0 00-1.5 0v2.5h-2.5a.75.75 0 000 1.5h2.5v2.5a.75.75 0 001.5 0v-2.5h2.5a.75.75 0 000-1.5h-2.5v-2.5z" clip-rule="evenodd"/>`))
}
