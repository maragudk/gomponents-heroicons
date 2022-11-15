package mini

import (
	g "github.com/maragudk/gomponents"

	h "github.com/maragudk/gomponents-heroicons"
)

func UserMinus(children ...g.Node) g.Node {
	return h.Mini(g.Group(children),
		g.Raw(`<path d="M11 5a3 3 0 11-6 0 3 3 0 016 0zM2.046 15.253c-.058.468.172.92.57 1.175A9.953 9.953 0 008 18c1.982 0 3.83-.578 5.384-1.573.398-.254.628-.707.57-1.175a6.001 6.001 0 00-11.908 0zM12.75 7.75a.75.75 0 000 1.5h5.5a.75.75 0 000-1.5h-5.5z"/>`))
}
