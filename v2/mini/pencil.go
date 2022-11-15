package mini

import (
	g "github.com/maragudk/gomponents"

	h "github.com/maragudk/gomponents-heroicons/v2"
)

func Pencil(children ...g.Node) g.Node {
	return h.Mini(g.Group(children),
		g.Raw(`<path d="M2.695 14.763l-1.262 3.154a.5.5 0 00.65.65l3.155-1.262a4 4 0 001.343-.885L17.5 5.5a2.121 2.121 0 00-3-3L3.58 13.42a4 4 0 00-.885 1.343z"/>`))
}
