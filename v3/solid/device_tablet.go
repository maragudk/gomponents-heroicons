package solid

import (
	g "maragu.dev/gomponents"

	h "maragu.dev/gomponents-heroicons/v3"
)

func DeviceTablet(children ...g.Node) g.Node {
	return h.Solid(g.Group(children),
		g.Raw(`<path d="M10.5 18a.75.75 0 000 1.5h3a.75.75 0 000-1.5h-3z"/>
  <path fill-rule="evenodd" d="M7.125 1.5A3.375 3.375 0 003.75 4.875v14.25A3.375 3.375 0 007.125 22.5h9.75a3.375 3.375 0 003.375-3.375V4.875A3.375 3.375 0 0016.875 1.5h-9.75zM6 4.875c0-.621.504-1.125 1.125-1.125h9.75c.621 0 1.125.504 1.125 1.125v14.25c0 .621-.504 1.125-1.125 1.125h-9.75A1.125 1.125 0 016 19.125V4.875z" clip-rule="evenodd"/>`))
}
