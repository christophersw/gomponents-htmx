package html

import g "github.com/maragudk/gomponents"

// progressively enhances anchors and forms to use AJAX requests
func HxBoost() g.Node {
	return g.Attr("hx-boost")
}

// shows a confim() dialog before issuing a request
func HxConfirm() g.Node {
	return g.Attr("hx-confirm")
}

// issues a DELETE to the specified URL
func HxDelete() g.Node {
	return g.Attr("hx-delete")
}

// disables htmx processing for the given node and any children nodes
func HxDisable() g.Node {
	return g.Attr("hx-disable")
}

// changes the request encoding type
func HxEncoding() g.Node {
	return g.Attr("hx-encoding")
}

// extensions to use for this element
func HxExt() g.Node {
	return g.Attr("hx-ext")
}

// issues a GET to the specified URL
func HxGet() g.Node {
	return g.Attr("hx-get")
}

// adds to the headers that will be submitted with the request
func HxHeaders() g.Node {
	return g.Attr("hx-headers")
}

// includes additional data in AJAX requests
func HxHistoryElt() g.Node {
	return g.Attr("hx-history-elt")
}

// the element to snapshot and restore during history navigation
func HxInclude() g.Node {
	return g.Attr("hx-include")
}

// the element to put the htmx-request class on during the AJAX request
func HxIndicator() g.Node {
	return g.Attr("hx-indicator")
}

// filters the parameters that will be submitted with a request
func HxParams() g.Node {
	return g.Attr("hx-params")
}

// issues a PATCH to the specified URL
func HxPatch() g.Node {
	return g.Attr("hx-patch")
}

// issues a POST to the specified URL
func HxPost() g.Node {
	return g.Attr("hx-post")
}

// preserves an element between requests
func HxPreserve() g.Node {
	return g.Attr("hx-preserve")
}

// shows a prompt before submitting a request
func HxPrompt() g.Node {
	return g.Attr("hx-prompt")
}

// pushes the URL into the location bar, creating a new history entry
func HxPushUrl() g.Node {
	return g.Attr("hx-push-url")
}

// issues a PUT to the specified URL
func HxPut() g.Node {
	return g.Attr("hx-put")
}

// configures various aspects of the request
func HxRequest() g.Node {
	return g.Attr("hx-request")
}

// selects a subset of the server response to process
func HxSelect() g.Node {
	return g.Attr("hx-select")
}

// establishes and listens to SSE sources for events
func HxSse() g.Node {
	return g.Attr("hx-sse")
}

// marks content in a response as being "Out of Band", i.e. swapped somewhere other than the target
func HxSwapOob() g.Node {
	return g.Attr("hx-swap-oob")
}

// controls how the response content is swapped into the DOM (e.g. 'outerHTML' or 'beforeEnd')
func HxSwap() g.Node {
	return g.Attr("hx-swap")
}

// specifies the target element to be swapped
func HxTarget() g.Node {
	return g.Attr("hx-target")
}

// specifies the event that triggers the request
func HxTrigger() g.Node {
	return g.Attr("hx-trigger")
}

// adds to the parameters that will be submitted with the request
func HxVals() g.Node {
	return g.Attr("hx-vals")
}

// establishes a WebSocket or sends information to one
func HxWs() g.Node {
	return g.Attr("hx-ws")
}
