package browser

import "syscall/js"

// The Document.visibilityState returns the visibility of the document,
// that is in which context this element is now visible.
// It is useful to know if the document is in the background or an invisible tab, or only loaded for pre-rendering.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/visibilityState
type VisibilityState int

const (
	HiddenVisibilityState VisibilityState = iota
	VisibleVisibilityState
	PrerenderVisibilityState
)

var visibilityStateToWasmTable = []string{"hidden", "visible", "prerender"}

var visibilityStateFromWasmTable = map[string]VisibilityState{
	"hidden":    HiddenVisibilityState,
	"visible":   VisibleVisibilityState,
	"prerender": PrerenderVisibilityState,
}

// JSValue is converting this enum into a javascript object
func (_this *VisibilityState) JSValue() js.Value {
	return js.ValueOf(_this.Value())
}

// Value is converting _this into javascript defined
// string value
func (_this VisibilityState) Value() string {
	idx := int(_this)
	if idx >= 0 && idx < len(visibilityStateToWasmTable) {
		return visibilityStateToWasmTable[idx]
	}
	panic("unknown input value")
}

// VisibilityStateFromJS is converting a javascript value into
// a VisibilityState enum value.
func VisibilityStateFromJS(value js.Value) VisibilityState {
	key := value.String()
	conv, ok := visibilityStateFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}
