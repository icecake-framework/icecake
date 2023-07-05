package dom

import (
	// "bytes"

	"reflect"

	"github.com/icecake-framework/icecake/pkg/console"
	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/js"
)

// type UIListener interface {
// 	Wrap(js.JSValueProvider)
// 	AddListeners()
// 	RemoveListeners()
// }

type UIComposer interface {

	// Meta returns a reference to render meta data
	html.RMetaProvider

	Wrap(js.JSValueProvider)

	AddListeners()

	RemoveListeners()
	// Mount()
	// UnMount()
}

type Composer interface {
	html.HTMLContentComposer
	UIComposer
}

/*****************************************************************************/

// UISnippet combines an htmlSnippet allowing html rendering of ick-name tags in different ways, and
// a wrapped dom.Element allowing event listening and other direct DOM interactions.
type UI struct {
	DOM Element
}

// Mount does nothing by default. Can be implemented by the component embedding UISnippet.
// func (ui *UI) Mount() {}

// UnMount does nothing by default. Can be implemented by the component embedding UISnippet.
// func (ui *UI) UnMount() {}

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

// RenderHTML builds and unfolds the UIcomposer and returns its html string.
// RenderHTML does not mount the component into the DOM.
// func (_parent *UISnippet) RenderHTML(_snippet UIComposer) (_html html.String) {
// 	out := new(bytes.Buffer)
// 	id, err := html.WriteSnippet(out, _snippet, nil)
// 	if err == nil {
// 		_parent.Embed(id, _snippet) // need to embed the snippet itself
// 		_html = html.String(out.String())
// 	}
// 	return _html
// }

// InsertSnippet insrets a _snippet within the _parent (according to the _where location) and add _parents lisneters
// func (_parent *UISnippet) InsertSnippet(_where INSERT_WHERE, _snippet any, _data *html.DataState) (_id string, _err error) {
// 	if _parent == nil || !_parent.DOM.IsDefined() {
// 		return "", console.Errorf("Snippet:InsertSnippetfailed on undefined _parent")
// 	}
// 	_id, _err = _parent.DOM.InsertSnippet(_where, _snippet, _data)
// 	_parent.AddListeners()
// 	return _id, nil
// }

// MountCSSLinks inserts links elements to the Head section of the Document for every csslinkref found in TheRegistry of components.
// If a link already exists for a csslinkref nothing is done.
// MountCSSLinks call is optional if your html head already contains stylesheet links for your css or if you import it in your own js code.
// MountCSSLinks must be called at the early begining of the wasm code.
// func MountCSSLinks() {
// 	reg := registry.Map()
// 	for ickname, e := range reg {
// 		fmt.Println("Mounting CSSLinks for", ickname)
// 		if !e.IsCSSLinkMounted() {
// 			links := e.CSSLinkRefs()
// 			if links != nil {
// 				for _, l := range links {
// 					if l == "" {
// 						continue
// 					}
// 					head := Doc().Head()
// 					children := head.ChildrenMatching(func(e *Element) bool {
// 						if e.TagName() == "LINK" {
// 							href, _ := e.Attribute("href")
// 							if href == l {
// 								return true
// 							}
// 						}
// 						return false
// 					})
// 					if len(children) == 0 {
// 						e := CreateElement("LINK").SetAttribute("href", l).SetAttribute("rel", "stylesheet")
// 						head.InsertElement(INSERT_LAST_CHILD, e)

// 					}
// 				}
// 			}
// 			e.SetCSSLinkMounted()
// 		}
// 	}
// }

/******************************************************************************
* Private Area
*******************************************************************************/

// mountSnippetTree addlisteners to the snippet and looks recursively for every childs with an id and add listeners to each of them.
// The snippet must have been wrapped with a DOM element before
func mountSnippetTree(parent html.RMetaProvider) (err error) {
	if parent.RMeta().IsMounted {
		console.Warnf("mountSnippetTree: parent:%q is already mounted", parent.RMeta().VirtualId)
		return
	}

	// mount children
	if embedded := parent.RMeta().Embedded(); embedded != nil {
		for _, emb := range embedded {
			if child, ok := emb.(UIComposer); ok {
				childid := child.RMeta().Id
				if childid != "" {
					console.Logf("mountSnippetTree: parent:%q is mounting %v id:%s", parent.RMeta().VirtualId, reflect.TypeOf(child).String(), childid)
					errm := TryWrapId(child, childid)
					if errm != nil {
						console.Errorf(errm.Error())
					} else {
						child.AddListeners()
						errm = mountSnippetTree(child)
					}
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

// TODO: must call unmount somewhere
// unmountSnippetTree remove listeners anc all Unmount recusrively for every embedded components
func unmountSnippetTree(_snippet UIComposer) {
	_snippet.RemoveListeners()
	// _snippet.UnMount()

	if embedded := _snippet.RMeta().Embedded(); embedded != nil {
		// DEBUG: console.Warnf("scanning %+v", embedded)
		for _, sub := range embedded {
			if cmp, ok := sub.(UIComposer); ok {
				// DEBUG: console.Warnf("wrapping %+v", w)
				unmountSnippetTree(cmp)
			}
		}
	}
}
