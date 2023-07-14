package dom

import (
	// "bytes"

	"reflect"

	"github.com/icecake-framework/icecake/pkg/console"
	"github.com/icecake-framework/icecake/pkg/ickcore"
	"github.com/icecake-framework/icecake/pkg/js"
)

type UIComposer interface {
	ickcore.TagBuilder
	ickcore.ContentComposer

	Wrap(js.JSValueProvider)

	AddListeners()

	RemoveListeners()
}

// type Composer interface {
// 	ickcore.ContentComposer
// 	UIComposer
// }

/*****************************************************************************/

// UISnippet combines an htmlSnippet allowing html rendering of ick-name tags in different ways, and
// a wrapped dom.Element allowing event listening and other direct DOM interactions.
type UI struct {
	DOM Element
}

// AddListeners does nothing by default. Can be implemented by the component embedding UISnippet.
func (ui *UI) AddListeners() {}

// Wrap implements the JSValueWrapper to enable wrapping of a dom.Element usually
// to wrap embedded component instantiated during unfolding an html string.
// Does not need to be overloaded by the component embedding UISnippet.
func (ui *UI) Wrap(jsvp js.JSValueProvider) {
	if ui.DOM.Value().Truthy() {
		console.Warnf("UI.wrap: UI element %q already wrapped", ui.DOM.Id())
	}
	ui.DOM.JSValue = jsvp.Value()
	if !ui.DOM.IsInDOM() {
		console.Warnf("UI.wrap: fails, %q not in DOM", ui.DOM.Id())
	}
}

// RemoveListeners remove all event handlers attached to this UISnippet Element.
// If RemoveListeners is implemented by the component embedding UISnippet then the UISnippet one should be called.
// Usually RemoveListeners does not need to be overloaded because every listeners added to the Element are automatically removed.
func (ui *UI) RemoveListeners() {
	ui.DOM.RemoveListeners()
}

// mountSnippetTree addlisteners to the snippet and looks recursively for every childs with an id and add listeners to each of them.
// Nothing is done with the parent but its IsMounted RMeta is turned on in case of success.
func mountSnippetTree(parent ickcore.RMetaProvider) (err error) {
	if parent.RMeta().IsMounted {
		console.Warnf("mountSnippetTree: parent:%s id:%q is already mounted", reflect.TypeOf(parent).String(), parent.RMeta().VirtualId)
		return
	}

	// mount children
	embedded := parent.RMeta().Embedded()
	// DEBUG: console.Warnf("mountSnippetTree: %v children", len(embedded))
	console.Warnf("mountSnippetTree: %v children", len(embedded))
	if embedded != nil {
		for _, emb := range embedded {
			// DEBUG: console.Warnf("mountSnippetTree: %+v child %s", emb, reflect.TypeOf(emb).String())
			if child, ok := emb.(UIComposer); ok {
				childid := child.RMeta().TagId
				if childid != "" {
					console.Logf("mountSnippetTree: parent:%v is mounting %v id:%q", reflect.TypeOf(parent).String(), reflect.TypeOf(child).String(), childid)
					errm := TryMountId(child, childid)
					if errm != nil && err == nil {
						err = errm
					}

				}
			}
		}
	}
	if err == nil {
		parent.RMeta().IsMounted = true
	}
	return err
}

// unmountSnippetTree remove listeners recusrively for every embedded child
func unmountSnippetTree(parent ickcore.RMetaProvider) {
	if embedded := parent.RMeta().Embedded(); embedded != nil {
		for _, sub := range embedded {
			if child, ok := sub.(UIComposer); ok {
				child.RemoveListeners()
				unmountSnippetTree(child)
			}
		}
	}
}
