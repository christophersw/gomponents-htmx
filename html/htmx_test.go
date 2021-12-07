package html_test

import (
	"fmt"
	"testing"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	"github.com/maragudk/gomponents/internal/assert"
)

func TestHtmxAttributes(t *testing.T) {

	cases := map[string]func() g.Node{
		"hx-boost":       HxBoost,
		"hx-confirm":     HxConfirm,
		"hx-delete":      HxDelete,
		"hx-disable":     HxDisable,
		"hx-encoding":    HxEncoding,
		"hx-ext":         HxExt,
		"hx-get":         HxGet,
		"hx-headers":     HxHeaders,
		"hx-history-elt": HxHistoryElt,
		"hx-include":     HxInclude,
		"hx-indicator":   HxIndicator,
		"hx-params":      HxParams,
		"hx-patch":       HxPatch,
		"hx-post":        HxPost,
		"hx-preserve":    HxPreserve,
		"hx-prompt":      HxPrompt,
		"hx-push-url":    HxPushUrl,
		"hx-put":         HxPut,
		"hx-request":     HxRequest,
		"hx-select":      HxSelect,
		"hx-sse":         HxSse,
		"hx-swap-oob":    HxSwapOob,
		"hx-swap":        HxSwap,
		"hx-target":      HxTarget,
		"hx-trigger":     HxTrigger,
		"hx-vals":        HxVals,
		"hx-ws":          HxWs,
	}

	for name, fn := range cases {
		t.Run(fmt.Sprintf("should output %v", name), func(t *testing.T) {
			n := g.El("div", fn())
			assert.Equal(t, fmt.Sprintf(`<div %v></div>`, name), n)
		})
	}
}
