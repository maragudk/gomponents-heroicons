package mini

import (
	g "maragu.dev/gomponents"

	h "maragu.dev/gomponents-heroicons/v3"
)

func DocumentMinus(children ...g.Node) g.Node {
	return h.Mini(g.Group(children),
		g.Raw(`<path fill-rule="evenodd" d="M4.5 2A1.5 1.5 0 003 3.5v13A1.5 1.5 0 004.5 18h11a1.5 1.5 0 001.5-1.5V7.621a1.5 1.5 0 00-.44-1.06l-4.12-4.122A1.5 1.5 0 0011.378 2H4.5zm7.75 9.75a.75.75 0 000-1.5h-4.5a.75.75 0 000 1.5h4.5z" clip-rule="evenodd"/>`))
}
