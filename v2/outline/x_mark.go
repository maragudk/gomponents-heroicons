package outline

import (
	g "github.com/maragudk/gomponents"

	h "github.com/maragudk/gomponents-heroicons/v2"
)

func XMark(children ...g.Node) g.Node {
	return h.Outline(g.Group(children),
		g.Raw(`<path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/>`))
}
