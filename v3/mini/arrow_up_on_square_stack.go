package mini

import (
	g "maragu.dev/gomponents"

	h "maragu.dev/gomponents-heroicons/v3"
)

func ArrowUpOnSquareStack(children ...g.Node) g.Node {
	return h.Mini(g.Group(children),
		g.Raw(`<path fill-rule="evenodd" d="M10.75 6h-2v4.25a.75.75 0 01-1.5 0V6h1.5V3.704l.943 1.048a.75.75 0 001.114-1.004l-2.25-2.5a.75.75 0 00-1.114 0l-2.25 2.5a.75.75 0 001.114 1.004l.943-1.048V6h-2A2.25 2.25 0 003 8.25v4.5A2.25 2.25 0 005.25 15h5.5A2.25 2.25 0 0013 12.75v-4.5A2.25 2.25 0 0010.75 6zM7 16.75v-.25h3.75a3.75 3.75 0 003.75-3.75V10h.25A2.25 2.25 0 0117 12.25v4.5A2.25 2.25 0 0114.75 19h-5.5A2.25 2.25 0 017 16.75z" clip-rule="evenodd"/>`))
}
