package mini

import (
	g "maragu.dev/gomponents"

	h "maragu.dev/gomponents-heroicons/v3"
)

func ReceiptPercent(children ...g.Node) g.Node {
	return h.Mini(g.Group(children),
		g.Raw(`<path fill-rule="evenodd" d="M4.93 1.31a41.401 41.401 0 0110.14 0C16.194 1.45 17 2.414 17 3.517V18.25a.75.75 0 01-1.075.676l-2.8-1.344-2.8 1.344a.75.75 0 01-.65 0l-2.8-1.344-2.8 1.344A.75.75 0 013 18.25V3.517c0-1.103.806-2.068 1.93-2.207zm8.85 5.97a.75.75 0 00-1.06-1.06l-6.5 6.5a.75.75 0 101.06 1.06l6.5-6.5zM9 8a1 1 0 11-2 0 1 1 0 012 0zm3 5a1 1 0 100-2 1 1 0 000 2z" clip-rule="evenodd"/>`))
}
