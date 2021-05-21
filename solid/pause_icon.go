package solid

import (
	g "github.com/maragudk/gomponents"

	h "github.com/maragudk/gomponents-heroicons"
)

func PauseIcon(children ...g.Node) g.Node {
	return h.Solid(g.Group(children),
		g.Raw(`<path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zM7 8a1 1 0 012 0v4a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v4a1 1 0 102 0V8a1 1 0 00-1-1z" clip-rule="evenodd"/>`))
}
