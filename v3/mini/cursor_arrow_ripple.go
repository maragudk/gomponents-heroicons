package mini

import (
	g "maragu.dev/gomponents"

	h "maragu.dev/gomponents-heroicons/v3"
)

func CursorArrowRipple(children ...g.Node) g.Node {
	return h.Mini(g.Group(children),
		g.Raw(`<path fill-rule="evenodd" d="M6.111 11.89A5.5 5.5 0 1115.501 8 .75.75 0 1017 8a7 7 0 10-11.95 4.95.75.75 0 001.06-1.06zm2.121-5.658a2.5 2.5 0 000 3.536.75.75 0 11-1.06 1.06A4 4 0 1114 8a.75.75 0 01-1.5 0 2.5 2.5 0 00-4.268-1.768zm2.534 1.279a.75.75 0 00-1.37.364l-.492 6.861a.75.75 0 001.204.65l1.043-.799.985 3.678a.75.75 0 001.45-.388l-.978-3.646 1.292.204a.75.75 0 00.74-1.16l-3.874-5.764z" clip-rule="evenodd"/>`))
}
