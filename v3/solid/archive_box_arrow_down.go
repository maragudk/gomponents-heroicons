package solid

import (
	g "maragu.dev/gomponents"

	h "maragu.dev/gomponents-heroicons/v3"
)

func ArchiveBoxArrowDown(children ...g.Node) g.Node {
	return h.Solid(g.Group(children),
		g.Raw(`<path d="M3.375 3C2.339 3 1.5 3.84 1.5 4.875v.75c0 1.036.84 1.875 1.875 1.875h17.25c1.035 0 1.875-.84 1.875-1.875v-.75C22.5 3.839 21.66 3 20.625 3H3.375z"/>
  <path fill-rule="evenodd" d="M3.087 9l.54 9.176A3 3 0 006.62 21h10.757a3 3 0 002.995-2.824L20.913 9H3.087zM12 10.5a.75.75 0 01.75.75v4.94l1.72-1.72a.75.75 0 111.06 1.06l-3 3a.75.75 0 01-1.06 0l-3-3a.75.75 0 111.06-1.06l1.72 1.72v-4.94a.75.75 0 01.75-.75z" clip-rule="evenodd"/>`))
}
