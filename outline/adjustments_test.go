package outline_test

import (
	"testing"

	. "github.com/maragudk/gomponents/html"

	"github.com/maragudk/gomponents-heroicons/internal/assert"
	"github.com/maragudk/gomponents-heroicons/outline"
)

func TestAdjustments(t *testing.T) {
	t.Run("returns svg component", func(t *testing.T) {
		assert.Equal(t, `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" aria-hidden="true" id="hat"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4"/></svg>`, outline.Adjustments(ID("hat")))
	})
}
