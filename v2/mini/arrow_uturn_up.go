package mini

import (
	g "github.com/maragudk/gomponents"

	h "github.com/maragudk/gomponents-heroicons/v2"
)

func ArrowUturnUp(children ...g.Node) g.Node {
	return h.Mini(g.Group(children),
		g.Raw(`<path fill-rule="evenodd" d="M17.768 7.793a.75.75 0 01-1.06-.025L12.75 3.622v10.003a5.375 5.375 0 01-10.75 0V10.75a.75.75 0 011.5 0v2.875a3.875 3.875 0 007.75 0V3.622L7.293 7.768a.75.75 0 01-1.086-1.036l5.25-5.5a.75.75 0 011.085 0l5.25 5.5a.75.75 0 01-.024 1.06z" clip-rule="evenodd"/>`))
}
