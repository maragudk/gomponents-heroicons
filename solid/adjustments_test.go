package solid_test

import (
	"testing"

	"github.com/maragudk/gomponents-heroicons/solid"
	"github.com/maragudk/gomponents/assert"
	. "github.com/maragudk/gomponents/html"
)

func TestAdjustments(t *testing.T) {
	t.Run("returns svg component", func(t *testing.T) {
		assert.Equal(t, `<svg viewBox="0 0 20 20" fill="currentColor" aria-hidden="true" id="hat"><path d="M5 4a1 1 0 00-2 0v7.268a2 2 0 000 3.464V16a1 1 0 102 0v-1.268a2 2 0 000-3.464V4zM11 4a1 1 0 10-2 0v1.268a2 2 0 000 3.464V16a1 1 0 102 0V8.732a2 2 0 000-3.464V4zM16 3a1 1 0 011 1v7.268a2 2 0 010 3.464V16a1 1 0 11-2 0v-1.268a2 2 0 010-3.464V4a1 1 0 011-1z"/></svg>`, solid.Adjustments(ID("hat")))
	})
}
