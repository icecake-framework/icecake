package dom

import (
	// "bytes"

	"bytes"
	"reflect"

	"github.com/icecake-framework/icecake/pkg/console"
	"github.com/icecake-framework/icecake/pkg/ickcore"
	"github.com/icecake-framework/icecake/pkg/js"
	"github.com/lolorenzo777/verbose"
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

// RefreshContent renders the cmp's content and re-insert it into the DOM.
// RefreshContent removes and re-add all listeners
func (ui *UI) RefreshContent(cmp ickcore.ContentComposer) (errx error) {
	unmountSnippet(cmp)
	out := new(bytes.Buffer)
	errx = cmp.RenderContent(out)
	if errx == nil {
		ui.DOM.InsertRawHTML(INSERT_BODY, out.String())
	}
	_, errx = mountSnippet(cmp, &ui.DOM)
	return errx
}

// mountSnippet addlisteners to the cmp snippet and looks recursively for every childs with an id and add listeners to each of them.
func mountSnippet(cmp ickcore.RMetaProvider, elem *Element) (mounted int, err error) {
	cmptype := reflect.TypeOf(cmp).String()
	if cmp.RMeta().IsMounted {
		verbose.Printf(verbose.ALERT, "mountSnippet: %s(vid:%q) is already mounted", cmptype, cmp.RMeta().VirtualId)
		return
	}

	// mount the composer
	mounted = 0
	if ui, is := cmp.(UIComposer); is && elem.IsDefined() {
		verbose.Debug("mountSnippet: mounting %s", cmptype)
		mounted = 1
		ui.Wrap(elem)
		ui.AddListeners()
		ui.RMeta().IsMounted = true
	}

	// mount children
	var errm error
	embedded := cmp.RMeta().Embedded()
	if embedded != nil && len(embedded) > 0 {
		// DEBUG: verbose.Debug("mountSnippet: %v children in %s", len(embedded), cmptype)

		for _, emb := range embedded {

			// DEBUG: verbose.Debug("mountSnippet: %s --> %+v", reflect.TypeOf(emb).String(), emb)

			var e *Element
			if childid := emb.RMeta().TagId; childid != "" {
				if child, ok := emb.(UIComposer); ok {
					e, errm = TryCastId(child, childid)
					if errm != nil && err == nil {
						err = errm
						continue
					}
				}
			}

			var m int
			m, errm = mountSnippet(emb, e)
			if errm != nil && err == nil {
				err = errm
			}
			mounted += m
		}
	}
	return mounted, err
}

// unmountSnippet remove listeners recusrively for every embedded child
func unmountSnippet(cmp ickcore.RMetaProvider) {
	if ui, is := cmp.(UIComposer); is {
		ui.RemoveListeners()
	}
	if embedded := cmp.RMeta().Embedded(); embedded != nil {
		for _, sub := range embedded {
			if child, ok := sub.(UIComposer); ok {
				child.RemoveListeners()
				unmountSnippet(child)
			}
		}
	}
	cmp.RMeta().IsMounted = false
}
