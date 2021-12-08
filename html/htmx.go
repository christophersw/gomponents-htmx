package html

import g "github.com/christophersw/gomponents-htmx"

// progressively enhances anchors and forms to use AJAX requests
func HxBoost(v string) g.Node {
	return g.Attr("hx-boost", v)
}

// shows a confim() dialog before issuing a request
func HxConfirm(v string) g.Node {
	return g.Attr("hx-confirm", v)
}

// issues a DELETE to the specified URL
func HxDelete(v string) g.Node {
	return g.Attr("hx-delete", v)
}

// disables htmx processing for the given node and any children nodes
func HxDisable(v string) g.Node {
	return g.Attr("hx-disable", v)
}

// changes the request encoding type
func HxEncoding(v string) g.Node {
	return g.Attr("hx-encoding", v)
}

// extensions to use for this element
func HxExt(v string) g.Node {
	return g.Attr("hx-ext", v)
}

// issues a GET to the specified URL
func HxGet(v string) g.Node {
	return g.Attr("hx-get", v)
}

// adds to the headers that will be submitted with the request
func HxHeaders(v string) g.Node {
	return g.Attr("hx-headers", v)
}

// includes additional data in AJAX requests
func HxHistoryElt(v string) g.Node {
	return g.Attr("hx-history-elt", v)
}

// the element to snapshot and restore during history navigation
func HxInclude(v string) g.Node {
	return g.Attr("hx-include", v)
}

// the element to put the htmx-request class on during the AJAX request
func HxIndicator(v string) g.Node {
	return g.Attr("hx-indicator", v)
}

// filters the parameters that will be submitted with a request
func HxParams(v string) g.Node {
	return g.Attr("hx-params", v)
}

// issues a PATCH to the specified URL
func HxPatch(v string) g.Node {
	return g.Attr("hx-patch", v)
}

// issues a POST to the specified URL
func HxPost(v string) g.Node {
	return g.Attr("hx-post", v)
}

// preserves an element between requests
func HxPreserve(v string) g.Node {
	return g.Attr("hx-preserve", v)
}

// shows a prompt before submitting a request
func HxPrompt(v string) g.Node {
	return g.Attr("hx-prompt", v)
}

// pushes the URL into the location bar, creating a new history entry
func HxPushUrl(v string) g.Node {
	return g.Attr("hx-push-url", v)
}

// issues a PUT to the specified URL
func HxPut(v string) g.Node {
	return g.Attr("hx-put", v)
}

// configures various aspects of the request
func HxRequest(v string) g.Node {
	return g.Attr("hx-request", v)
}

// selects a subset of the server response to process
func HxSelect(v string) g.Node {
	return g.Attr("hx-select", v)
}

// establishes and listens to SSE sources for events
func HxSse(v string) g.Node {
	return g.Attr("hx-sse", v)
}

// marks content in a response as being "Out of Band", i.e. swapped somewhere other than the target
func HxSwapOob(v string) g.Node {
	return g.Attr("hx-swap-oob", v)
}

// controls how the response content is swapped into the DOM (e.g. 'outerHTML' or 'beforeEnd')
func HxSwap(v string) g.Node {
	return g.Attr("hx-swap", v)
}

// specifies the target element to be swapped
func HxTarget(v string) g.Node {
	return g.Attr("hx-target", v)
}

// specifies the event that triggers the request
func HxTrigger(v string) g.Node {
	return g.Attr("hx-trigger", v)
}

// adds to the parameters that will be submitted with the request
func HxVals(v string) g.Node {
	return g.Attr("hx-vals", v)
}

// establishes a WebSocket or sends information to one
func HxWs(v string) g.Node {
	return g.Attr("hx-ws", v)
}
