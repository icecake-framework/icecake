package dom

import (
	"fmt"
	"net/url"
	"time"

	syscalljs "syscall/js"

	"github.com/icecake-framework/icecake/internal/helper"
	"github.com/icecake-framework/icecake/pkg/console"
	"github.com/icecake-framework/icecake/pkg/event"
	"github.com/icecake-framework/icecake/pkg/js"
)

/******************************************************************************
* Document
******************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/Document/readyState
type DOC_READYSTATE string

const (
	DOC_READY       DOC_READYSTATE = "loading"
	DOC_INTERACTIVE DOC_READYSTATE = "interactive"
	DOC_COMPLETE    DOC_READYSTATE = "complete"
)

// The Document.visibilityState returns the visibility of the document,
// that is in which context this element is now visible.
// It is useful to know if the document is in the background or an invisible tab, or only loaded for pre-rendering.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/visibilityState
type DOC_VISIBILITYSTATE string

const (
	DOC_HIDDEN    DOC_VISIBILITYSTATE = "hidden"
	DOC_VISIBLE   DOC_VISIBILITYSTATE = "visible"
	DOC_PRERENDER DOC_VISIBILITYSTATE = "prerender"
)

var document Document

// GetDocument returns the current document within the current window
func Doc() Document {
	if document.IsDefined() {
		return document
	}
	jsdoc := js.Global().Get("document")
	if !jsdoc.IsObject() {
		console.Stackf()
		console.Panicf("Unable to get document")
	}
	document.JSValue = jsdoc
	return document
}

// Id returns the Element found in the doc with the _elementId attribute.
// returns an undefined Element if the id is not found.
func Id(_elementId string) (_result *Element) {
	return Doc().ChildById(_elementId)
}

// CreateElement creates the HTML element specified by tagName, or an HTMLUnknownElement if tagName isn't recognized.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement
func CreateElement(_tagName string) *Element {
	elem := Doc().Call("createElement", _tagName)
	return CastElement(elem)
}

// CreateElement creates the HTML element specified by tagName, or an HTMLUnknownElement if tagName isn't recognized.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement
func CreateTextNode(_text string) *Node {
	node := Doc().Call("createTextNode", _text)
	return CastNode(node)
}

// Document represents any web page loaded in the browser and serves as an entry point into the web page's content, which is the DOM tree.
//
// The Document describes the common properties and methods for any kind of document.
// Depending on the document's type (e.g. HTML, XML, SVG, …), a larger API is available:
// HTML documents, served with the "text/html" content type,
// also implement the HTMLDocument interface, whereas XML and SVG documents implement the XMLDocument interface.
type Document struct {
	Node
}

// CastDocument is casting a js.Value into Document.
func CastDocument(_jsv js.JSValueProvider) *Document {
	if _jsv.Value().Type() != js.TYPE_OBJECT {
		console.Errorf("casting Document failed")
		return &Document{}
	}
	doc := new(Document)
	doc.JSValue = _jsv.Value()
	return doc
}

/******************************************************************************
* Document's properties
******************************************************************************/

// CompatMode ndicates whether the document is rendered in Quirks mode or Standards mode.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/compatMode
func (_doc Document) CompatMode() string {
	return _doc.Get("compatMode").String()
}

// CharacterSet returns the character encoding of the document that it's currently rendered with.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/characterSet
func (_doc Document) CharacterSet() string {
	return _doc.Get("characterSet").String()
}

// Doctype Name, The type of the document. It is always "html" for HTML documents, but will vary for XML documents.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/doctype
func (_doc Document) DocTypeName() string {
	return _doc.Get("doctype").GetString("name")
}

// A string with an identifier of the type of document. Always empty ("") for HTML, it will be, for example, "-//W3C//DTD SVG 1.1//EN" for SVG documents.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/doctype
func (_doc Document) DocTypePublicId() string {
	return _doc.Get("doctype").GetString("name")
}

// A string containing the URL to the associated DTD. Always empty ("") for HTML, it will be, for example, "http://www.w3.org/2000/svg" for SVG documents.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/doctype
func (_doc Document) DocTypeSystemId() string {
	return _doc.Get("doctype").GetString("name")
}

// ContentType returns the MIME type that the document is being rendered as.
// This may come from HTTP headers or other sources of MIME information,
// and might be affected by automatic type conversions performed by either the browser or extensions.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/contentType
func (_doc Document) ContentType() string {
	return _doc.Get("contentType").String()
}

// Referrer returns the URI of the page that linked to this page.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/referrer
func (_doc Document) Referrer() *url.URL {
	ref := _doc.Get("referrer").String()
	u, _ := url.Parse(ref)
	return u
}

// Cookie lets you read and write cookies associated with the document.
// It serves as a getter and setter for the actual values of the cookies.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/cookie
func (_doc Document) Cookie() string {
	return _doc.Get("cookie").String()
}

// Cookie lets you read and write cookies associated with the document.
// It serves as a getter and setter for the actual values of the cookies.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/cookie
func (_doc *Document) SetCookie(value string) *Document {
	_doc.Set("cookie", value)
	return _doc
}

// LastModified returns a string containing the date and time on which the current document was last modified.
//
// TODO: test LastModified
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/lastModified
func (_doc Document) LastModified() time.Time {
	strtime := _doc.Get("lastModified").String()
	t, err := time.Parse("01/02/2006 15:04:05", strtime)
	if err != nil {
		console.Warnf("Document:LastModified error: %s", err.Error())
	}
	return t
}

// ReadyState describes the loading state of the document. When the value of this property changes, a readystatechange event fires on the document object.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/readyState
func (_doc Document) ReadyState() DOC_READYSTATE {
	value := _doc.Get("readyState").String()
	return DOC_READYSTATE(value)
}

// Title gets or sets the current title of the document. When present, it defaults to the value of the <title>.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/title
func (_doc Document) Title() string {
	return _doc.Get("title").String()
}

// Title gets or sets the current title of the document. When present, it defaults to the value of the <title>.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/title
func (_doc *Document) SetTitle(value string) *Document {
	_doc.Set("title", value)
	return _doc
}

// The Document.body property represents the <body> or <frameset> node of the current document, or null if no such element exists.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/body
func (_doc Document) Body() *Element {
	value := _doc.Get("body")
	return CastElement(value)
}

// The Document.body property represents the <body> or <frameset> node of the current document, or null if no such element exists.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/body
func (_doc *Document) SetBody(value *Element) *Document {
	_doc.Set("body", value)
	return _doc
}

// Head  returns the <head> element of the current document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/head
func (_doc Document) Head() *Element {
	elem := _doc.Get("head")
	return CastElement(elem)
}

// DocumentElement returns the Element that is the root element of the document (for example, the <html> element for HTML documents).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/documentElement
func (_doc Document) RootElement() *Element {
	elem := _doc.Get("documentElement")
	return CastElement(elem)
}

// ActiveElement returns the Element within the DOM that currently has focus.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/activeElement
func (_doc Document) FocusedElement() *Element {
	elem := _doc.Get("activeElement")
	return CastElement(elem)
}

// FullscreenElement returns the Element that is currently being presented in fullscreen mode in this document, or null if fullscreen mode is not currently in use.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/fullscreenElement
func (_doc Document) FullscreenElement() *Element {
	elem := _doc.Get("fullscreenElement")
	return CastElement(elem)
}

// DesignMode controls whether the entire document is editable.
// Valid values are "on" and "off". According to the specification, this property is meant to default to "off".
// Firefox follows this standard. The earlier versions of Chrome and IE default to "inherit".
// Starting in Chrome 43, the default is "off" and "inherit" is no longer supported. In IE6-10, the value is capitalized.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/designMode
func (_doc Document) DesignMode() string {
	return _doc.Get("designMode").String()
}

// DesignMode controls whether the entire document is editable.
// Valid values are "on" and "off". According to the specification, this property is meant to default to "off".
// Firefox follows this standard. The earlier versions of Chrome and IE default to "inherit".
// Starting in Chrome 43, the default is "off" and "inherit" is no longer supported. In IE6-10, the value is capitalized.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/designMode
func (_doc *Document) SetDesignMode(value string) *Document {
	_doc.Set("designMode", value)
	return _doc
}

// Hidden returns a Boolean value indicating if the page is considered hidden or not.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/hidden
func (_doc Document) Hidden() bool {
	return _doc.Get("hidden").Bool()
}

// VisibilityState returns the visibility of the document, that is in which context this element is now visible.
// It is useful to know if the document is in the background or an invisible tab, or only loaded for pre-rendering.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/visibilityState
func (_doc Document) VisibilityState() DOC_VISIBILITYSTATE {
	value := _doc.Get("visibilityState").String()
	return DOC_VISIBILITYSTATE(value)
}

// HasFocus returns a boolean value indicating whether the document or any element inside the document has focus.
// This method can be used to determine whether the active element in a document has focus.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/hasFocus
func (_doc Document) HasFocus() bool {
	return _doc.Call("hasFocus").Bool()
}

// ChildElementCount returns the number of child elements of the document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/childElementCount
func (_doc Document) ChildrenCount() int {
	return _doc.GetInt("childElementCount")
}

// GetElementsByTagName returns an HTMLCollection of elements with the given tag name.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByTagName
func (_doc Document) ChildrenByTagName(qualifiedName string) []*Element {
	elems := _doc.Call("getElementsByTagName", qualifiedName)
	if !elems.IsDefined() {
		console.Warnf("ChildrenByTagName failed: %q not found\n", qualifiedName)
		return make([]*Element, 0)
	}
	return CastElements(elems)
}

// GetElementsByClassName returns an array-like object of all child elements which have all of the given class name(s).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByClassName
func (_doc Document) ChildrenByClassName(classNames string) []*Element {
	elems := _doc.Call("getElementsByClassName", classNames)
	if !elems.IsDefined() {
		console.Warnf("ChildrenByClassName failed: %q not found\n", classNames)
		return make([]*Element, 0)
	}
	return CastElements(elems)
}

// GetElementsByName returns a NodeList Collection of elements with a given name attribute in the document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByName
func (_doc Document) ChildrenByName(elementName string) []*Element {
	elems := _doc.Call("getElementsByName", elementName)
	if !elems.IsDefined() {
		console.Warnf("ChildrenByName failed: %q not found\n", elementName)
		return make([]*Element, 0)
	}
	return CastElements(elems)
}

// GetElementById returns an Element object representing the element whose id property matches the specified string.
// Since element IDs are required to be unique if specified, they're a useful way to get access to a specific element quickly.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById
func (_doc Document) ChildById(_elementId string) (_result *Element) {
	_elementId = helper.Normalize(_elementId)
	if _elementId == "" {
		return new(Element)
	}
	elem := _doc.Call("getElementById", _elementId)
	if elem.Truthy() && CastNode(elem).NodeType() == NT_ELEMENT {
		//DEBUG: console.Warnf("ChildById success: %q \n", _elementId)
		return CastElement(elem)
	}
	console.Warnf("ChildById failed: %q not found, or not an <Element>\n", _elementId)
	return new(Element)
}

func (_doc Document) CheckId(_elementId string) (_result *Element, _err error) {
	_elementId = helper.Normalize(_elementId)
	if _elementId == "" {
		return new(Element), fmt.Errorf("CheckId called with an empty id")
	}
	elem := _doc.Call("getElementById", _elementId)
	if elem.Truthy() && CastNode(elem).NodeType() == NT_ELEMENT {
		//DEBUG: console.Warnf("ChildById success: %q \n", _elementId)
		return CastElement(elem), nil
	}
	return new(Element), fmt.Errorf("CheckId %q not found, or not an <Element>", _elementId)
}

// QuerySelector returns the first Element within the document that matches the specified selector, or group of selectors.
// If no matches are found, null is returned.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector
func (_doc Document) SelectorQueryFirst(selectors string) *Element {
	elem := _doc.Call("querySelector", selectors)
	if elem.Truthy() && CastNode(elem).NodeType() == NT_ELEMENT {
		return CastElement(elem)
	}
	console.Warnf("querySelector failed: %q not found, or not a <Element>\n", selectors)
	return new(Element)
}

// querySelectorAll returns a static (not live) NodeList representing a list of the document's elements that match the specified group of selectors.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelectorAll
func (_doc Document) SelectorQueryAll(selectors string) []*Element {
	elems := _doc.Call("querySelectorAll", selectors)
	if !elems.IsDefined() {
		console.Warnf("SelectorQueryAll failed: %q not found\n", selectors)
		return nil
	}
	return CastElements(elems)
}

// GetElementAtPoint returns the topmost Element at the specified coordinates (relative to the viewport).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/elementFromPoint
func (_doc Document) ChildAtPoint(x float64, y float64) *Element {
	elem := _doc.Call("elementFromPoint", x, y)
	return CastElement(elem)
}

// GetElementsAtPoint eturns an array of all elements at the specified coordinates (relative to the viewport).
// The elements are ordered from the topmost to the bottommost box of the viewport.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/elementsFromPoint
func (_doc Document) ChildrenAtPoint(x float64, y float64) (_result []*Element) {
	elems := _doc.Call("elementsFromPoint", x, y)
	len := elems.Value().Length()
	_result = make([]*Element, len)
	for i := 0; i < len; i++ {
		elem := elems.Index(i)
		_result[i] = CastElement(elem)
	}
	return _result
}

// Prepend inserts a set of Node objects or string objects before the first child of the document.
// String objects are inserted as equivalent Text nodes.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/prepend
func (_doc *Document) PrependNodes(nodes []*Node) {
	var _args []any = make([]any, len(nodes))
	var _end int
	for _, n := range nodes {
		if n != nil {
			jsn := n.JSValue
			_args[_end] = jsn
			_end++
		}
	}
	_doc.Call("prepend", _args[0:_end]...)
}

// Prepend inserts a set of Node objects or string objects before the first child of the document.
// String objects are inserted as equivalent Text nodes.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/prepend
func (_doc *Document) PrepenStrings(strs []string) {
	var _args []any = make([]any, len(strs))
	var _end int
	for _, n := range strs {
		_args[_end] = n
		_end++
	}
	_doc.Call("prepend", _args[0:_end]...)
}

// Append inserts a set of Node objects or string objects after the last child of the document.
// String objects are inserted as equivalent Text nodes.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/append
func (_doc *Document) AppendNodes(nodes []*Node) {
	var _args []any = make([]any, len(nodes))
	var _end int
	for _, n := range nodes {
		if n != nil {
			jsn := n.JSValue
			_args[_end] = jsn
			_end++
		}
	}
	_doc.Call("append", _args[0:_end]...)
}

// Append inserts a set of Node objects or string objects after the last child of the document.
// String objects are inserted as equivalent Text nodes.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/append
func (_doc *Document) AppendStrings(strs []string) {
	var _args []any = make([]any, len(strs))
	var _end int
	for _, n := range strs {
		_args[_end] = n
		_end++
	}
	_doc.Call("append", _args[0:_end]...)
}

/******************************************************************************
* Document's  GENRIC_EVENT
******************************************************************************/

func makeDoc_Generic_Event(_listener func(*event.Event, *Document)) syscalljs.Func {
	fn := func(this js.JSValue, args []js.JSValue) any {
		value := args[0]
		evt := event.CastEvent(value)
		target := CastDocument(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				console.Errorf("Error processing event %q on document: %s", evt.Type(), r)
				console.Stackf()
			}
		}()
		_listener(evt, target)
		return nil
	}
	return js.FuncOf(fn)
}

func (_doc *Document) AddGenericEvent(_evttyp event.GENERIC_EVENT, _listener func(*event.Event, *Document)) {
	evh := makeDoc_Generic_Event(_listener)
	_doc.addListener(string(_evttyp), evh)
}

/******************************************************************************
* Document's  MOUSE_EVENT
******************************************************************************/

func makeDoc_Mouse_Event(_listener func(*event.MouseEvent, *Document)) syscalljs.Func {

	fn := func(this js.JSValue, args []js.JSValue) any {
		value := args[0]
		evt := event.CastMouseEvent(value)
		target := CastDocument(value.Get("target"))
		_listener(evt, target)
		return nil
	}

	return js.FuncOf(fn)
}

func (_doc *Document) AddMouseEvent(_evttyp event.MOUSE_EVENT, _listener func(*event.MouseEvent, *Document)) {
	evh := makeDoc_Mouse_Event(_listener)
	_doc.addListener(string(_evttyp), evh)
}

/******************************************************************************
* Document's  FOCUS_EVENT
******************************************************************************/

func makeDoc_Focus_Event(_listener func(*event.FocusEvent, *Document)) syscalljs.Func {
	fn := func(this js.JSValue, args []js.JSValue) any {
		value := args[0]
		evt := event.CastFocusEvent(value)
		target := CastDocument(value.Get("target"))
		_listener(evt, target)
		return nil
	}
	return js.FuncOf(fn)
}

func (_doc *Document) AddFocusEvent(_evttyp event.FOCUS_EVENT, _listener func(*event.FocusEvent, *Document)) {
	evh := makeDoc_Focus_Event(_listener)
	_doc.addListener(string(_evttyp), evh)
}

/******************************************************************************
* Document's  POINTER_EVENT
******************************************************************************/

func makeDoc_Pointer_Event(_listener func(*event.PointerEvent, *Document)) syscalljs.Func {
	fn := func(this js.JSValue, args []js.JSValue) any {
		value := args[0]
		evt := event.CastPointerEvent(value)
		target := CastDocument(value.Get("target"))
		_listener(evt, target)
		return nil
	}
	return js.FuncOf(fn)
}

func (_doc *Document) AddPointerEvent(_evttyp event.POINTER_EVENT, _listener func(*event.PointerEvent, *Document)) {
	evh := makeDoc_Pointer_Event(_listener)
	_doc.addListener(string(_evttyp), evh)
}

/******************************************************************************
* Document's  INPUT_EVENT
******************************************************************************/

func makeDoc_Input_Event(_listener func(*event.InputEvent, *Document)) syscalljs.Func {
	fn := func(this js.JSValue, args []js.JSValue) any {
		value := args[0]
		evt := event.CastInputEvent(value)
		target := CastDocument(value.Get("target"))
		_listener(evt, target)
		return nil
	}
	return js.FuncOf(fn)
}

func (_doc *Document) AddInputEvent(_evttyp event.INPUT_EVENT, _listener func(*event.InputEvent, *Document)) {
	evh := makeDoc_Input_Event(_listener)
	_doc.addListener(string(_evttyp), evh)
}

/******************************************************************************
* Document's  KEYBOARD_EVENT
******************************************************************************/

func makeDoc_Keyboard_Event(_listener func(*event.KeyboardEvent, *Document)) syscalljs.Func {
	fn := func(this js.JSValue, args []js.JSValue) any {
		value := args[0]
		evt := event.CastKeyboardEvent(value)
		target := CastDocument(value.Get("target"))
		_listener(evt, target)
		return nil
	}
	return js.FuncOf(fn)
}

func (_doc *Document) AddKeyboardEvent(_evttyp event.KEYBOARD_EVENT, _listener func(*event.KeyboardEvent, *Document)) {
	evh := makeDoc_Keyboard_Event(_listener)
	_doc.addListener(string(_evttyp), evh)
}

/******************************************************************************
* Document's  WHEEL_EVENT
******************************************************************************/

// event attribute: UIEvent
func makeDoc_UIEvent(_listener func(*event.UIEvent, *Document)) syscalljs.Func {
	fn := func(this js.JSValue, args []js.JSValue) any {
		value := args[0]
		evt := event.CastUIEvent(value)
		target := CastDocument(value.Get("target"))
		_listener(evt, target)
		return nil
	}
	return js.FuncOf(fn)
}

// AddResize is adding doing AddEventListener for 'Resize' on target.
// This method is returning allocated javascript function that need to be released.
func (_doc *Document) AddEventResize(_listener func(*event.UIEvent, *Document)) {
	evh := makeDoc_UIEvent(_listener)
	_doc.addListener("resize", evh)
}

/******************************************************************************
* Document's  WHEEL_EVENT
******************************************************************************/

func makeDoc_Wheel_Event(_listener func(*event.WheelEvent, *Document)) syscalljs.Func {
	fn := func(this js.JSValue, args []js.JSValue) any {
		value := args[0]
		evt := event.CastWheelEvent(value)
		target := CastDocument(value.Get("target"))
		_listener(evt, target)
		return nil
	}
	return js.FuncOf(fn)
}

// AddWheel is adding doing AddEventListener for 'Wheel' on target.
// This method is returning allocated javascript function that need to be released.
func (_doc *Document) AddEventWheel(_listener func(*event.WheelEvent, *Document)) {
	evh := makeDoc_Wheel_Event(_listener)
	_doc.addListener("wheel", evh)
}

/******************************************************************************
* Document's  FULLSCREEN_EVENT
******************************************************************************/

// AddFullscreenChange is adding doing AddEventListener for 'FullscreenChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_doc *Document) AddFullscreenEvent(_evttyp event.FULLSCREEN_EVENT, _listener func(*event.Event, *Document)) {
	evh := makeDoc_Generic_Event(_listener)
	_doc.addListener(string(_evttyp), evh)
}
