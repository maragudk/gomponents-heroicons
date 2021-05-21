package solid

import (
	g "github.com/maragudk/gomponents"

	h "github.com/maragudk/gomponents-heroicons"
)

func ChartPie(children ...g.Node) g.Node {
	return h.Solid(g.Group(children),
		g.Raw(`<path d="M2 10a8 8 0 018-8v8h8a8 8 0 11-16 0z"/>
  <path d="M12 2.252A8.014 8.014 0 0117.748 8H12V2.252z"/>`))
}
