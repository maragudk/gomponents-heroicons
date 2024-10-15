package heroicons_test

import (
	"testing"

	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	heroicons "github.com/maragudk/gomponents-heroicons"
	"github.com/maragudk/gomponents-heroicons/internal/assert"
)

func TestOutline(t *testing.T) {
	t.Run("returns svg with viewbox 0 0 24 24, no fill, currentColor stroke", func(t *testing.T) {
		assert.Equal(t, `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" aria-hidden="true" class="hat"><path></path></svg>`, heroicons.Outline(Class("hat"), El("path")))
	})
}

func TestSolid(t *testing.T) {
	t.Run("returns svg with viewbox 0 0 20 20, currentColor fill", func(t *testing.T) {
		assert.Equal(t, `<svg viewBox="0 0 20 20" fill="currentColor" aria-hidden="true" class="hat"><path></path></svg>`, heroicons.Solid(Class("hat"), El("path")))
	})
}
