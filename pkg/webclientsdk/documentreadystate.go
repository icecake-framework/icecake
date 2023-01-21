package browser

import "syscall/js"

/****************************************************************************
* DocumentReadyState
*****************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/Document/readyState
type DocumentReadyState int

const (
	LoadingDocumentReadyState DocumentReadyState = iota
	InteractiveDocumentReadyState
	CompleteDocumentReadyState
)

var documentReadyStateToWasmTable = []string{"loading", "interactive", "complete"}

var documentReadyStateFromWasmTable = map[string]DocumentReadyState{
	"loading":     LoadingDocumentReadyState,
	"interactive": InteractiveDocumentReadyState,
	"complete":    CompleteDocumentReadyState,
}

// JSValue is converting this enum into a javascript object
func (_this *DocumentReadyState) JSValue() js.Value {
	return js.ValueOf(_this.Value())
}

// Value is converting this into javascript defined string value
func (_this DocumentReadyState) Value() string {
	idx := int(_this)
	if idx >= 0 && idx < len(documentReadyStateToWasmTable) {
		return documentReadyStateToWasmTable[idx]
	}
	panic("unknown input value")
}

// DocumentReadyStateFromJS is converting a javascript value into
// a DocumentReadyState enum value.
func DocumentReadyStateFromJS(value js.Value) DocumentReadyState {
	key := value.String()
	conv, ok := documentReadyStateFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}
