package dom

import (
	"fmt"
	"net/url"
	"syscall/js"
)

/******************************************************************************
* Document
******************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/Document/readyState
type DOCUMENT_READYSTATE string

const (
	DOC_READY       DOCUMENT_READYSTATE = "loading"
	DOC_INTERACTIVE DOCUMENT_READYSTATE = "interactive"
	DOC_COMPLETE    DOCUMENT_READYSTATE = "complete"
)

// The Document.visibilityState returns the visibility of the document,
// that is in which context this element is now visible.
// It is useful to know if the document is in the background or an invisible tab, or only loaded for pre-rendering.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/visibilityState
type DOCUMENT_VISIBILITYSTATE string

const (
	DOC_HIDDEN    DOCUMENT_VISIBILITYSTATE = "hidden"
	DOC_VISIBLE   DOCUMENT_VISIBILITYSTATE = "visible"
	DOC_PRERENDER DOCUMENT_VISIBILITYSTATE = "prerender"
)

// Document represents any web page loaded in the browser and serves as an entry point into the web page's content, which is the DOM tree.
//
// The Document describes the common properties and methods for any kind of document.
// Depending on the document's type (e.g. HTML, XML, SVG, …), a larger API is available:
// HTML documents, served with the "text/html" content type,
// also implement the HTMLDocument interface, whereas XML and SVG documents implement the XMLDocument interface.
type Document struct {
	Node
}

// NewDocumentFromJS is casting a js.Value into Document.
func NewDocumentFromJS(value js.Value) *Document {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := new(Document)
	ret.jsValue = value
	return ret
}

// GetDocument returns the current document within the current window
func GetDocument() (_ret *Document) {
	value := js.Global().Get("document")
	_ret = NewDocumentFromJS(value)
	if _ret == nil {
		fmt.Println("GetDocument() failed")
	}
	return _ret
}

/******************************************************************************
* Document's properties
******************************************************************************/

// CompatMode ndicates whether the document is rendered in Quirks mode or Standards mode.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/compatMode
func (_doc *Document) CompatMode() string {
	return _doc.jsValue.Get("compatMode").String()
}

// CharacterSet returns the character encoding of the document that it's currently rendered with.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/characterSet
func (_doc *Document) CharacterSet() string {
	return _doc.jsValue.Get("characterSet").String()
}

// Doctype Returns the Document Type Declaration (DTD) associated with current document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/doctype
func (_doc *Document) Doctype() *DocumentType {
	value := _doc.jsValue.Get("doctype")
	return MakeDocumentTypeFromJS(value)
}

// ContentType returns the MIME type that the document is being rendered as.
// This may come from HTTP headers or other sources of MIME information,
// and might be affected by automatic type conversions performed by either the browser or extensions.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/contentType
func (_doc *Document) ContentType() string {
	return _doc.jsValue.Get("contentType").String()
}

// DocumentElement returns the Element that is the root element of the document (for example, the <html> element for HTML documents).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/documentElement
func (_doc *Document) RootElement() (_ret *Element) {
	value := _doc.jsValue.Get("documentElement")
	return NewElementFromJS(value)
}

// Location returns a Location object, which contains information about the URL of the document and provides methods for changing that URL and loading another URL.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/location
func (_doc *Document) Location() *Location {
	value := _doc.jsValue.Get("location")
	return NewLocationFromJS(value)
}

// Referrer returns the URI of the page that linked to this page.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/referrer
func (_doc *Document) Referrer() (_ret *url.URL) {
	value := _doc.jsValue.Get("referrer").String()
	_ret, _ = url.Parse(value)
	return _ret
}

// Cookie lets you read and write cookies associated with the document.
// It serves as a getter and setter for the actual values of the cookies.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/cookie
func (_doc *Document) Cookie() string {
	return _doc.jsValue.Get("cookie").String()
}

// Cookie lets you read and write cookies associated with the document.
// It serves as a getter and setter for the actual values of the cookies.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/cookie
func (_doc *Document) SetCookie(value string) *Document {
	_doc.jsValue.Set("cookie", value)
	return _doc
}

// LastModified returns a string containing the date and time on which the current document was last modified.
//
// TODO: handle time.Time
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/lastModified
func (_doc *Document) LastModified() string {
	return _doc.jsValue.Get("lastModified").String()
}

// ReadyState describes the loading state of the document. When the value of this property changes, a readystatechange event fires on the document object.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/readyState
func (_doc *Document) ReadyState() DOCUMENT_READYSTATE {
	value := _doc.jsValue.Get("readyState").String()
	return DOCUMENT_READYSTATE(value)
}

// Title gets or sets the current title of the document. When present, it defaults to the value of the <title>.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/title
func (_doc *Document) Title() string {
	return _doc.jsValue.Get("title").String()
}

// Title gets or sets the current title of the document. When present, it defaults to the value of the <title>.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/title
func (_doc *Document) SetTitle(value string) *Document {
	_doc.jsValue.Set("title", value)
	return _doc
}

// The Document.body property represents the <body> or <frameset> node of the current document, or null if no such element exists.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/body
func (_doc *Document) Body() *HTMLElement {
	value := _doc.jsValue.Get("body")
	return NewHTMLElementFromJS(value)
}

// The Document.body property represents the <body> or <frameset> node of the current document, or null if no such element exists.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/body
func (_doc *Document) SetBody(value *HTMLElement) *Document {
	_doc.jsValue.Set("body", value.JSValue())
	return _doc
}

// Head  returns the <head> element of the current document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/head
func (_doc *Document) Head() *HTMLHeadElement {
	value := _doc.jsValue.Get("head")
	return NewHTMLHeadElementFromJS(value)
}

// Images returns a collection of the images in the current HTML document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/images
func (_doc *Document) Images() Nodes {
	value := _doc.jsValue.Get("images")
	// return newHTMLCollectionFromJS(value)
	return MakeNodesFromJSNodeList(value)
}

// Embeds returns a list of the embedded <embed> elements within the current document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/embeds
func (_doc *Document) Embeds() Nodes {
	value := _doc.jsValue.Get("embeds")
	// return newHTMLCollectionFromJS(value)
	return MakeNodesFromJSNodeList(value)
}

// Plugins returns an HTMLCollection object containing one or more HTMLEmbedElements representing the <embed> elements in the current document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/plugins
func (_doc *Document) Plugins() Nodes {
	value := _doc.jsValue.Get("plugins")
	// return newHTMLCollectionFromJS(value)
	return MakeNodesFromJSNodeList(value)
}

// Links returns a collection of all <area> elements and <a> elements in a document with a value for the href attribute.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/links
func (_doc *Document) Links() Nodes {
	value := _doc.jsValue.Get("links")
	// return newHTMLCollectionFromJS(value)
	return MakeNodesFromJSNodeList(value)
}

// Forms returns an HTMLCollection listing all the <form> elements contained in the document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/forms
func (_doc *Document) Forms() Nodes {
	value := _doc.jsValue.Get("forms")
	return MakeNodesFromJSNodeList(value)
}

// DefaultView returns the window object associated with a document, or null if none is available.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/defaultView
func (_doc *Document) Window() *Window {
	w := _doc.jsValue.Get("defaultView")
	return NewWindowFromJS(w)
}

// ActiveElement returns the Element within the DOM that currently has focus.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/activeElement
func (_doc *Document) FocusedElement() *Element {
	value := _doc.jsValue.Get("activeElement")
	return NewElementFromJS(value)
}

// DesignMode controls whether the entire document is editable.
// Valid values are "on" and "off". According to the specification, this property is meant to default to "off".
// Firefox follows this standard. The earlier versions of Chrome and IE default to "inherit".
// Starting in Chrome 43, the default is "off" and "inherit" is no longer supported. In IE6-10, the value is capitalized.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/designMode
func (_doc *Document) DesignMode() string {
	return _doc.jsValue.Get("designMode").String()
}

// DesignMode controls whether the entire document is editable.
// Valid values are "on" and "off". According to the specification, this property is meant to default to "off".
// Firefox follows this standard. The earlier versions of Chrome and IE default to "inherit".
// Starting in Chrome 43, the default is "off" and "inherit" is no longer supported. In IE6-10, the value is capitalized.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/designMode
func (_doc *Document) SetDesignMode(value string) *Document {
	_doc.jsValue.Set("designMode", value)
	return _doc
}

// Hidden returns a Boolean value indicating if the page is considered hidden or not.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/hidden
func (_doc *Document) Hidden() bool {
	return _doc.jsValue.Get("hidden").Bool()
}

// VisibilityState returns the visibility of the document, that is in which context this element is now visible.
// It is useful to know if the document is in the background or an invisible tab, or only loaded for pre-rendering.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/visibilityState
func (_doc *Document) VisibilityState() DOCUMENT_VISIBILITYSTATE {
	value := _doc.jsValue.Get("visibilityState").String()
	return DOCUMENT_VISIBILITYSTATE(value)
}

// FullscreenElement returns the Element that is currently being presented in fullscreen mode in this document, or null if fullscreen mode is not currently in use.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/fullscreenElement
func (_doc *Document) FullscreenElement() *Element {
	e := _doc.jsValue.Get("fullscreenElement")
	return NewElementFromJS(e)
}

// ChildElementCount returns the number of child elements of the document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/childElementCount
func (_doc *Document) ChildrenCount() uint {
	return (uint)(_doc.jsValue.Get("childElementCount").Int())
}

// Children returns a live HTMLCollection which contains all of the child elements of the document upon which it was called.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/children
func (_doc *Document) Children() Nodes {
	nodes := _doc.jsValue.Get("children")
	return MakeNodesFromJSNodeList(nodes)
}

// FirstElementChild returns the document's first child Element, or null if there are no child elements.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/firstElementChild
func (_doc *Document) ChildFirst() *Element {
	e := _doc.jsValue.Get("firstElementChild")
	return NewElementFromJS(e)
}

// LastElementChild eturns the document's last child Element, or null if there are no child elements.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/lastElementChild
func (_doc *Document) ChildLast() *Element {
	e := _doc.jsValue.Get("lastElementChild")
	return NewElementFromJS(e)
}

// GetElementsByTagName returns an HTMLCollection of elements with the given tag name.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByTagName
func (_doc *Document) ChildrenByTagName(qualifiedName string) Nodes {
	nodes := _doc.jsValue.Call("getElementsByTagName", qualifiedName)
	if typ := nodes.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		ConsoleWarn("ChildrenByTagName failed: %q not found\n", qualifiedName)
		return make(Nodes, 0)
	}
	return MakeNodesFromJSNodeList(nodes)
}

// GetElementsByClassName returns an array-like object of all child elements which have all of the given class name(s).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByClassName
func (_doc *Document) ChildrenByClassName(classNames string) Nodes {
	nodes := _doc.jsValue.Call("getElementsByClassName", classNames)
	if typ := nodes.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		ConsoleWarn("ChildrenByClassName failed: %q not found\n", classNames)
		return make(Nodes, 0)
	}
	return MakeNodesFromJSNodeList(nodes)
}

// GetElementsByName returns a NodeList Collection of elements with a given name attribute in the document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByName
func (_doc *Document) ChildrenByName(elementName string) (_result Nodes) {
	nodes := _doc.jsValue.Call("getElementsByName", elementName)
	if typ := nodes.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		ConsoleWarn("ChildrenByName failed: %q not found\n", elementName)
		return nil
	}
	return MakeNodesFromJSNodeList(nodes)
}

// GetElementById returns an Element object representing the element whose id property matches the specified string.
// Since element IDs are required to be unique if specified, they're a useful way to get access to a specific element quickly.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById
func (_doc *Document) ChildById(elementId string) (_result *Element) {
	e := _doc.jsValue.Call("getElementById", elementId)
	if typ := e.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		ConsoleWarn("ChildById failed: %q not found\n", elementId)
		return nil
	}
	return NewElementFromJS(e)
}

// QuerySelector returns the first Element within the document that matches the specified selector, or group of selectors.
// If no matches are found, null is returned.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector
func (_doc *Document) SelectorQueryFirst(selectors string) *Element {
	e := _doc.jsValue.Call("querySelector", selectors)
	return NewElementFromJS(e)
}

// querySelectorAll returns a static (not live) NodeList representing a list of the document's elements that match the specified group of selectors.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelectorAll
func (_doc *Document) SelectorQueryAll(selectors string) Nodes {
	nodes := _doc.jsValue.Call("querySelectorAll", selectors)
	if typ := nodes.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		ConsoleWarn("SelectorQueryAll failed: %q not found\n", selectors)
		return nil
	}
	return MakeNodesFromJSNodeList(nodes)
}

// CreateElement creates the HTML element specified by tagName, or an HTMLUnknownElement if tagName isn't recognized.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement
func (_doc *Document) CreateElement(localName string) *Element {
	e := _doc.jsValue.Call("createElement", localName)
	return NewElementFromJS(e)
}

// CreateAttribute  creates a new attribute node, and returns it. The object created is a node implementing the Attr interface.
// The DOM does not enforce what sort of attributes can be added to a particular element in this manner.
//
// # TODO test CreateAttribute
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/createAttribute
func (_doc *Document) CreateAttribute(localName string) *Attribute {
	attr := _doc.jsValue.Call("createAttribute", localName)
	return NewAttributeFromJS(attr)
}

// GetElementAtPoint returns the topmost Element at the specified coordinates (relative to the viewport).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/elementFromPoint
func (_doc *Document) ChildAtPoint(x float64, y float64) *Element {
	e := _doc.jsValue.Call("elementFromPoint", x, y)
	return NewElementFromJS(e)
}

// GetElementsAtPoint eturns an array of all elements at the specified coordinates (relative to the viewport).
// The elements are ordered from the topmost to the bottommost box of the viewport.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/elementsFromPoint
func (_doc *Document) ChildrenAtPoint(x float64, y float64) (_result []*Element) {
	elems := _doc.jsValue.Call("elementsFromPoint", x, y)
	len := elems.Length()
	_result = make([]*Element, len)
	for i := 0; i < len; i++ {
		elem := elems.Index(i)
		_result[i] = NewElementFromJS(elem)
	}
	return _result
}

// HasFocus returns a boolean value indicating whether the document or any element inside the document has focus.
// This method can be used to determine whether the active element in a document has focus.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/hasFocus
func (_doc *Document) HasFocus() (_result bool) {
	return _doc.jsValue.Call("hasFocus").Bool()
}

// Prepend inserts a set of Node objects or string objects before the first child of the document.
// String objects are inserted as equivalent Text nodes.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/prepend
func (_doc *Document) Prepend(nodes ...*Union) {
	var _args []interface{} = make([]interface{}, len(nodes))
	var _end int
	for _, n := range nodes {
		jsn := n.JSValue()
		_args[_end] = jsn
		_end++
	}
	_doc.jsValue.Call("prepend", _args[0:_end]...)
}

// Append inserts a set of Node objects or string objects after the last child of the document.
// String objects are inserted as equivalent Text nodes.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/append
func (_doc *Document) Append(nodes ...*Union) {
	var _args []interface{} = make([]interface{}, len(nodes))
	var _end int
	for _, n := range nodes {
		jsn := n.JSValue()
		_args[_end] = jsn
		_end++
	}
	_doc.jsValue.Call("append", _args[0:_end]...)
}

/******************************************************************************
* Document's  GENRIC_EVENT
******************************************************************************/

type ListenerDoc_Generic func(event *Event, target *Document)

func makeDoc_Generic_Event(listener ListenerDoc_Generic) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := NewEventFromJS(value)
		target := NewDocumentFromJS(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_doc *Document) AddGenericEvent(evttype GENERIC_EVENT, listener ListenerDoc_Generic) js.Func {
	callback := makeDoc_Generic_Event(listener)
	_doc.jsValue.Call("addEventListener", evttype, callback)
	return callback
}

/******************************************************************************
* Document's  MOUSE_EVENT
******************************************************************************/

type ListenerDoc_Mouse func(event *MouseEvent, target *Document)

func makeDoc_Mouse_Event(listener ListenerDoc_Mouse) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := NewMouseEventFromJS(value)
		target := NewDocumentFromJS(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_doc *Document) AddMouseEvent(evttype MOUSE_EVENT, listener ListenerDoc_Mouse) js.Func {
	callback := makeDoc_Mouse_Event(listener)
	_doc.jsValue.Call("addEventListener", evttype, callback)
	return callback
}

/******************************************************************************
* Document's  FOCUS_EVENT
******************************************************************************/

type ListenerDoc_Focus func(event *FocusEvent, target *Document)

func makeDoc_Focus_Event(listener ListenerDoc_Focus) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := NewFocusEventFromJS(value)
		target := NewDocumentFromJS(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_doc *Document) AddFocusEvent(evttype FOCUS_EVENT, listener ListenerDoc_Focus) js.Func {
	callback := makeDoc_Focus_Event(listener)
	_doc.jsValue.Call("addEventListener", evttype, callback)
	return callback
}

/******************************************************************************
* Document's  POINTER_EVENT
******************************************************************************/

type ListenerDoc_Pointer func(event *PointerEvent, target *Document)

func makeDoc_Pointer_Event(listener ListenerDoc_Pointer) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := NewPointerEventFromJS(value)
		target := NewDocumentFromJS(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_doc *Document) AddPointerEvent(evttype POINTER_EVENT, listener ListenerDoc_Pointer) js.Func {
	callback := makeDoc_Pointer_Event(listener)
	_doc.jsValue.Call("addEventListener", evttype, callback)
	return callback
}

/******************************************************************************
* Document's  INPUT_EVENT
******************************************************************************/

type ListenerDoc_Input func(event *InputEvent, target *Document)

func makeDoc_Input_Event(listener ListenerDoc_Input) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := NewInputEventFromJS(value)
		target := NewDocumentFromJS(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_doc *Document) AddInputEvent(evttype INPUT_EVENT, listener ListenerDoc_Input) js.Func {
	callback := makeDoc_Input_Event(listener)
	_doc.jsValue.Call("addEventListener", evttype, callback)
	return callback
}

/******************************************************************************
* Document's  KEYBOARD_EVENT
******************************************************************************/

type ListenerDoc_Keyboard func(event *KeyboardEvent, target *Document)

func makeDoc_Keyboard_Event(listener ListenerDoc_Keyboard) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := NewKeyboardEventFromJS(value)
		target := NewDocumentFromJS(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_doc *Document) AddKeyboardEvent(evttype KEYBOARD_EVENT, listener ListenerDoc_Keyboard) js.Func {
	callback := makeDoc_Keyboard_Event(listener)
	_doc.jsValue.Call("addEventListener", evttype, callback)
	return callback
}

/******************************************************************************
* Document's  WHEEL_EVENT
******************************************************************************/

type Listener_DocResize func(event *UIEvent, target *Document)

// event attribute: UIEvent
func makeDoc_UIEvent(listener Listener_DocResize) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := NewUIEventFromJS(value)
		target := NewDocumentFromJS(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddResize is adding doing AddEventListener for 'Resize' on target.
// This method is returning allocated javascript function that need to be released.
func (_doc *Document) AddEventResize(listener Listener_DocResize) js.Func {
	callback := makeDoc_UIEvent(listener)
	_doc.jsValue.Call("addEventListener", "resize", callback)
	return callback
}

/******************************************************************************
* Document's  WHEEL_EVENT
******************************************************************************/

type ListenerDoc_Wheel func(event *WheelEvent, target *Document)

func makeDoc_Wheel_Event(listener ListenerDoc_Wheel) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := NewWheelEventFromJS(value)
		target := NewDocumentFromJS(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddWheel is adding doing AddEventListener for 'Wheel' on target.
// This method is returning allocated javascript function that need to be released.
func (_doc *Document) AddEventWheel(listener func(event *WheelEvent, currentTarget *Document)) js.Func {
	callback := makeDoc_Wheel_Event(listener)
	_doc.jsValue.Call("addEventListener", "wheel", callback)
	return callback
}

/******************************************************************************
* Document's  FULLSCREEN_EVENT
******************************************************************************/

// AddFullscreenChange is adding doing AddEventListener for 'FullscreenChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_elem *Document) AddFullscreenEvent(evttype FULLSCREEN_EVENT, listener ListenerDoc_Generic) js.Func {
	cb := makeDoc_Generic_Event(listener)
	_elem.jsValue.Call("addEventListener", evttype, cb)
	return cb
}
