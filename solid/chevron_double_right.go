package solid

import (
	g "github.com/maragudk/gomponents"

	h "github.com/maragudk/gomponents-heroicons"
)

func ChevronDoubleRight(children ...g.Node) g.Node {
	return h.Solid(g.Group(children),
		g.Raw(`<path fill-rule="evenodd" d="M10.293 15.707a1 1 0 010-1.414L14.586 10l-4.293-4.293a1 1 0 111.414-1.414l5 5a1 1 0 010 1.414l-5 5a1 1 0 01-1.414 0z" clip-rule="evenodd"/>
  <path fill-rule="evenodd" d="M4.293 15.707a1 1 0 010-1.414L8.586 10 4.293 5.707a1 1 0 011.414-1.414l5 5a1 1 0 010 1.414l-5 5a1 1 0 01-1.414 0z" clip-rule="evenodd"/>`))
}
