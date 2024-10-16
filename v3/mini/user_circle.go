package mini

import (
	g "maragu.dev/gomponents"

	h "maragu.dev/gomponents-heroicons/v3"
)

func UserCircle(children ...g.Node) g.Node {
	return h.Mini(g.Group(children),
		g.Raw(`<path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-5.5-2.5a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0zM10 12a5.99 5.99 0 00-4.793 2.39A6.483 6.483 0 0010 16.5a6.483 6.483 0 004.793-2.11A5.99 5.99 0 0010 12z" clip-rule="evenodd"/>`))
}
