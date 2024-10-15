package heroicons_test

import (
	"testing"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	heroicons "maragu.dev/gomponents-heroicons/v3"
	"maragu.dev/gomponents-heroicons/v3/internal/assert"
)

func TestOutline(t *testing.T) {
	t.Run("returns svg with no fill, viewBox 0 0 24 24, stroke width 1.5, currentColor stroke", func(t *testing.T) {
		assert.Equal(t, `<svg fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true" class="hat"><path></path></svg>`, heroicons.Outline(Class("hat"), El("path")))
	})
}

func TestSolid(t *testing.T) {
	t.Run("returns svg with viewbox 0 0 24 24, currentColor fill", func(t *testing.T) {
		assert.Equal(t, `<svg viewBox="0 0 24 24" fill="currentColor" aria-hidden="true" class="hat"><path></path></svg>`, heroicons.Solid(Class("hat"), El("path")))
	})
}

func TestMini(t *testing.T) {
	t.Run("returns svg with viewbox 0 0 20 20, currentColor fill", func(t *testing.T) {
		assert.Equal(t, `<svg viewBox="0 0 20 20" fill="currentColor" aria-hidden="true" class="hat"><path></path></svg>`, heroicons.Mini(Class("hat"), El("path")))
	})
}
