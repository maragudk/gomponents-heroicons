package mini

import (
	g "maragu.dev/gomponents"

	h "maragu.dev/gomponents-heroicons/v3"
)

func ArrowLeftCircle(children ...g.Node) g.Node {
	return h.Mini(g.Group(children),
		g.Raw(`<g clip-path="url(#clip0_9_2121)">
    <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.25-7.25a.75.75 0 000-1.5H8.66l2.1-1.95a.75.75 0 10-1.02-1.1l-3.5 3.25a.75.75 0 000 1.1l3.5 3.25a.75.75 0 001.02-1.1l-2.1-1.95h4.59z" clip-rule="evenodd"/>
  </g>
  <defs>
    <clipPath id="clip0_9_2121">
      <path d="M0 0h20v30H0z"/>
    </clipPath>
  </defs>`))
}
