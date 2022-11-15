package mini

import (
	g "github.com/maragudk/gomponents"

	h "github.com/maragudk/gomponents-heroicons/v2"
)

func VideoCameraSlash(children ...g.Node) g.Node {
	return h.Mini(g.Group(children),
		g.Raw(`<path d="M1 13.75V7.182L9.818 16H3.25A2.25 2.25 0 011 13.75zM13 6.25v6.568L4.182 4h6.568A2.25 2.25 0 0113 6.25zM19 4.75a.75.75 0 00-1.28-.53l-3 3a.75.75 0 00-.22.53v4.5c0 .199.079.39.22.53l3 3a.75.75 0 001.28-.53V4.75zM2.28 4.22a.75.75 0 00-1.06 1.06l10.5 10.5a.75.75 0 101.06-1.06L2.28 4.22z"/>`))
}
