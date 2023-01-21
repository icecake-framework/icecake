package browser

import (
	"log"
	"syscall/js"
)

/******************************************************************************
* Document
******************************************************************************/

// Document represents any web page loaded in the browser and serves as an entry point into the web page's content, which is the DOM tree.
//
// The Document describes the common properties and methods for any kind of document.
// Depending on the document's type (e.g. HTML, XML, SVG, …), a larger API is available:
// HTML documents, served with the "text/html" content type,
// also implement the HTMLDocument interface, whereas XML and SVG documents implement the XMLDocument interface.
type Document struct {
	Node
}

// MakeDocumentFromJS is casting a js.Value into Document.
func MakeDocumentFromJS(value js.Value) *Document {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Document{}
	ret.jsValue = value
	return ret
}

// GetDocument returns the current document within the current window
func GetDocument() *Document {
	var ret *Document
	_klass := js.Global()
	value := _klass.Get("document")
	ret = MakeDocumentFromJS(value)
	if ret == nil {
		log.Println("GetDocument() failed")
	}
	return ret
}

/******************************************************************************
* Document's properties
******************************************************************************/

// URL returns the document location as a string.
//
// Extensions for HTMLDocument.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/URL
func (_this *Document) URL() string {
	if _this == nil {
		log.Println("URL() call on a nil Document")
		return ""
	}
	var ret string
	value := _this.jsValue.Get("URL")
	ret = (value).String()
	return ret
}

// DocumentURI returns the document location as a string.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/documentURI
func (_this *Document) DocumentURI() string {
	if _this == nil {
		log.Println("DocumentURI() call on a nil Document")
		return ""
	}
	var ret string
	value := _this.jsValue.Get("documentURI")
	ret = (value).String()
	return ret
}

// Origin returning attribute 'origin' with
// type string (idl: USVString).
func (_this *Document) Origin() string {
	if _this == nil {
		log.Println("Origin() call on a nil Document")
		return ""
	}
	var ret string
	value := _this.jsValue.Get("origin")
	ret = (value).String()
	return ret
}

// CompatMode ndicates whether the document is rendered in Quirks mode or Standards mode.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/compatMode
func (_this *Document) CompatMode() string {
	if _this == nil {
		log.Println("CompatMode() call on a nil Document")
		return ""
	}
	var ret string
	value := _this.jsValue.Get("compatMode")
	ret = (value).String()
	return ret
}

// CharacterSet returns the character encoding of the document that it's currently rendered with.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/characterSet
func (_this *Document) CharacterSet() string {
	if _this == nil {
		log.Println("CharacterSet() call on a nil Document")
		return ""
	}
	var ret string
	value := _this.jsValue.Get("characterSet")
	ret = (value).String()
	return ret
}

// Doctype Returns the Document Type Declaration (DTD) associated with current document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/doctype
func (_this *Document) Doctype() *DocumentType {
	if _this == nil {
		log.Println("Doctype() call on a nil Document")
		return nil
	}
	value := _this.jsValue.Get("doctype")
	return MakeDocumentTypeFromJS(value)
}

// ContentType returns the MIME type that the document is being rendered as.
// This may come from HTTP headers or other sources of MIME information,
// and might be affected by automatic type conversions performed by either the browser or extensions.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/contentType
func (_this *Document) ContentType() string {
	if _this == nil {
		log.Println("ContentType() call on a nil Document")
		return ""
	}
	var ret string
	value := _this.jsValue.Get("contentType")
	ret = (value).String()
	return ret
}

// DocumentElement returns the Element that is the root element of the document (for example, the <html> element for HTML documents).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/documentElement
func (_this *Document) DocumentElement() *Element {
	if _this == nil {
		log.Println("DocumentElement() call on a nil Document")
		return nil
	}
	var ret *Element
	value := _this.jsValue.Get("documentElement")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = MakeElementFromJS(value)
	}
	return ret
}

// Location returns a Location object, which contains information about the URL of the document and provides methods for changing that URL and loading another URL.
//
// Extensions for HTMLDocument.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/location
func (_this *Document) Location() *Location {
	if _this == nil {
		log.Println("Location() call on a nil Document")
		return nil
	}
	var ret *Location
	value := _this.jsValue.Get("location")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = LocationFromJS(value)
	}
	return ret
}

// Referrer returns the URI of the page that linked to this page.
//
// Extensions for HTMLDocument.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/referrer
func (_this *Document) Referrer() string {
	if _this == nil {
		log.Println("Referrer() call on a nil Document")
		return ""
	}
	var ret string
	value := _this.jsValue.Get("referrer")
	ret = (value).String()
	return ret
}

// Cookie lets you read and write cookies associated with the document.
// It serves as a getter and setter for the actual values of the cookies.
//
// Extensions for HTMLDocument.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/cookie
func (_this *Document) Cookie() string {
	if _this == nil {
		log.Println("Cookie() call on a nil Document")
		return ""
	}
	var ret string
	value := _this.jsValue.Get("cookie")
	ret = (value).String()
	return ret
}

// Cookie lets you read and write cookies associated with the document.
// It serves as a getter and setter for the actual values of the cookies.
//
// Extensions for HTMLDocument.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/cookie
func (_this *Document) SetCookie(value string) *Document {
	if _this == nil {
		log.Println("SetCookie() call on a nil Document")
		return nil
	}
	input := value
	_this.jsValue.Set("cookie", input)
	return _this
}

// LastModified returns a string containing the date and time on which the current document was last modified.
//
// Extensions for HTMLDocument.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/lastModified
func (_this *Document) LastModified() string {
	if _this == nil {
		log.Println("LastModified() call on a nil Document")
		return ""
	}
	var ret string
	value := _this.jsValue.Get("lastModified")
	ret = (value).String()
	return ret
}

// ReadyState describes the loading state of the document. When the value of this property changes, a readystatechange event fires on the document object.
//
// Extensions for HTMLDocument.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/readyState
func (_this *Document) ReadyState() DocumentReadyState {
	if _this == nil {
		log.Println("ReadyState() call on a nil Document")
		return 0 // aka LoadingDocumentReadyState
	}
	var ret DocumentReadyState
	value := _this.jsValue.Get("readyState")
	ret = DocumentReadyStateFromJS(value)
	return ret
}

// Title gets or sets the current title of the document. When present, it defaults to the value of the <title>.
//
// Extensions for HTMLDocument.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/title
func (_this *Document) Title() string {
	if _this == nil {
		log.Println("Title() call on a nil Document")
		return ""
	}
	var ret string
	value := _this.jsValue.Get("title")
	ret = (value).String()
	return ret
}

// Title gets or sets the current title of the document. When present, it defaults to the value of the <title>.
//
// Extensions for HTMLDocument.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/title
func (_this *Document) SetTitle(value string) *Document {
	if _this == nil {
		log.Println("SetTitle() call on a nil Document")
		return nil
	}
	input := value
	_this.jsValue.Set("title", input)
	return _this
}

// Dir is a string representing the directionality of the text of the document, whether left to right (default) or right to left.
// Possible values are 'rtl', right to left, and 'ltr', left to right.
//
// Extensions for HTMLDocument.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/dir
func (_this *Document) Dir() string {
	if _this == nil {
		log.Println("Dir() call on a nil Document")
		return ""
	}
	var ret string
	value := _this.jsValue.Get("dir")
	ret = (value).String()
	return ret
}

// Dir is a string representing the directionality of the text of the document, whether left to right (default) or right to left.
// Possible values are 'rtl', right to left, and 'ltr', left to right.
//
// Extensions for HTMLDocument.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/dir
func (_this *Document) SetDir(value string) *Document {
	if _this == nil {
		log.Println("SetDir() call on a nil Document")
		return nil
	}
	input := value
	_this.jsValue.Set("dir", input)
	return _this
}

// The Document.body property represents the <body> or <frameset> node of the current document, or null if no such element exists.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/body
func (_this *Document) Body() *HTMLElement {
	if _this == nil {
		log.Println("Body() call on a nil Document")
		return nil
	}
	var ret *HTMLElement
	value := _this.jsValue.Get("body")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = HTMLElementFromJS(value)
	}
	return ret
}

// The Document.body property represents the <body> or <frameset> node of the current document, or null if no such element exists.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/body
func (_this *Document) SetBody(value *HTMLElement) *Document {
	if _this == nil {
		log.Println("SetBody() call on a nil Document")
		return nil
	}
	input := value.JSValue()
	_this.jsValue.Set("body", input)
	return _this
}

// Head  returns the <head> element of the current document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/head
func (_this *Document) Head() *HTMLHeadElement {
	if _this == nil {
		log.Println("SetBody() call on a nil Document")
		return nil
	}
	value := _this.jsValue.Get("head")
	return HTMLHeadElementFromJS(value)
}

// Images returns a collection of the images in the current HTML document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/images
func (_this *Document) Images() *HTMLCollection {
	if _this == nil {
		log.Println("SetBody() call on a nil Document")
		return nil
	}
	value := _this.jsValue.Get("images")
	return HTMLCollectionFromJS(value)
}

// Embeds returns a list of the embedded <embed> elements within the current document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/embeds
func (_this *Document) Embeds() *HTMLCollection {
	if _this == nil {
		log.Println("SetBody() call on a nil Document")
		return nil
	}
	value := _this.jsValue.Get("embeds")
	return HTMLCollectionFromJS(value)
}

// Plugins returns an HTMLCollection object containing one or more HTMLEmbedElements representing the <embed> elements in the current document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/plugins
func (_this *Document) Plugins() *HTMLCollection {
	if _this == nil {
		log.Println("SetBody() call on a nil Document")
		return nil
	}
	value := _this.jsValue.Get("plugins")
	return HTMLCollectionFromJS(value)
}

// Links returns a collection of all <area> elements and <a> elements in a document with a value for the href attribute.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/links
func (_this *Document) Links() *HTMLCollection {
	if _this == nil {
		log.Println("SetBody() call on a nil Document")
		return nil
	}
	value := _this.jsValue.Get("links")
	return HTMLCollectionFromJS(value)
}

// Forms returns an HTMLCollection listing all the <form> elements contained in the document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/forms
func (_this *Document) Forms() *HTMLCollection {
	if _this == nil {
		log.Println("SetBody() call on a nil Document")
		return nil
	}
	value := _this.jsValue.Get("forms")
	return HTMLCollectionFromJS(value)
}

// DefaultView returns the window object associated with a document, or null if none is available.
//
// Extensions for HTMLDocument.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/defaultView
func (_this *Document) DefaultView() *Window {
	if _this == nil {
		log.Println("SetBody() call on a nil Document")
		return nil
	}
	value := _this.jsValue.Get("defaultView")
	return WindowFromJS(value)
}

// ActiveElement returns the Element within the DOM that currently has focus.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/activeElement
func (_this *Document) ActiveElement() *Element {
	if _this == nil {
		log.Println("SetBody() call on a nil Document")
		return nil
	}
	value := _this.jsValue.Get("activeElement")
	return MakeElementFromJS(value)
}

// DesignMode controls whether the entire document is editable.
// Valid values are "on" and "off". According to the specification, this property is meant to default to "off".
// Firefox follows this standard. The earlier versions of Chrome and IE default to "inherit".
// Starting in Chrome 43, the default is "off" and "inherit" is no longer supported. In IE6-10, the value is capitalized.
//
// Extensions for HTMLDocument.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/designMode
func (_this *Document) DesignMode() string {
	if _this == nil {
		log.Println("SetBody() call on a nil Document")
		return ""
	}
	value := _this.jsValue.Get("designMode")
	return (value).String()
}

// DesignMode controls whether the entire document is editable.
// Valid values are "on" and "off". According to the specification, this property is meant to default to "off".
// Firefox follows this standard. The earlier versions of Chrome and IE default to "inherit".
// Starting in Chrome 43, the default is "off" and "inherit" is no longer supported. In IE6-10, the value is capitalized.
//
// Extensions for HTMLDocument.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/designMode
func (_this *Document) SetDesignMode(value string) *Document {
	if _this == nil {
		log.Println("SetBody() call on a nil Document")
		return nil
	}
	input := value
	_this.jsValue.Set("designMode", input)
	return _this
}

// Hidden returns a Boolean value indicating if the page is considered hidden or not.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/hidden
func (_this *Document) Hidden() bool {
	if _this == nil {
		log.Println("SetBody() call on a nil Document")
		return false
	}
	value := _this.jsValue.Get("hidden")
	return (value).Bool()
}

// VisibilityState returns the visibility of the document, that is in which context this element is now visible.
// It is useful to know if the document is in the background or an invisible tab, or only loaded for pre-rendering.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/visibilityState
func (_this *Document) VisibilityState() VisibilityState {
	if _this == nil {
		log.Println("SetBody() call on a nil Document")
		return 0 // aka HiddenVisibilityState
	}
	value := _this.jsValue.Get("visibilityState")
	return VisibilityStateFromJS(value)
}

// FullscreenElement returns the Element that is currently being presented in fullscreen mode in this document, or null if fullscreen mode is not currently in use.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/fullscreenElement
func (_this *Document) FullscreenElement() *Element {
	if _this == nil {
		log.Println("SetBody() call on a nil Document")
		return nil
	}
	value := _this.jsValue.Get("fullscreenElement")
	return MakeElementFromJS(value)
}

// Children returns a live HTMLCollection which contains all of the child elements of the document upon which it was called.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/children
func (_this *Document) Children() *HTMLCollection {
	if _this == nil {
		log.Println("SetBody() call on a nil Document")
		return nil
	}
	value := _this.jsValue.Get("children")
	return HTMLCollectionFromJS(value)
}

// FirstElementChild returns the document's first child Element, or null if there are no child elements.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/firstElementChild
func (_this *Document) FirstElementChild() *Element {
	if _this == nil {
		log.Println("SetBody() call on a nil Document")
		return nil
	}
	value := _this.jsValue.Get("firstElementChild")
	return MakeElementFromJS(value)
}

// LastElementChild eturns the document's last child Element, or null if there are no child elements.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/lastElementChild
func (_this *Document) LastElementChild() *Element {
	if _this == nil {
		log.Println("SetBody() call on a nil Document")
		return nil
	}
	value := _this.jsValue.Get("lastElementChild")
	return MakeElementFromJS(value)
}

// ChildElementCount returns the number of child elements of the document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/childElementCount
func (_this *Document) ChildElementCount() uint {
	if _this == nil {
		log.Println("SetBody() call on a nil Document")
		return 0
	}
	var ret uint
	value := _this.jsValue.Get("childElementCount")
	ret = (uint)((value).Int())
	return ret
}

/******************************************************************************
* Document's  events
******************************************************************************/

// event attribute: Event
func eventFuncDocument_Event(listener func(event *Event, target *Document)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *Event
		value := args[0]
		incoming := value.Get("target")
		ret = MakeEventFromJS(value)
		src := MakeDocumentFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddAbort is adding doing AddEventListener for 'Abort' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventAbort(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "abort", cb)
	return cb
}

// SetOnAbort is assigning a function to 'onabort'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnAbort(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onabort", cb)
	return cb
}

// event attribute: htmlevent.MouseEvent
func eventFuncDocument_MouseEvent(listener func(event *MouseEvent, target *Document)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *MouseEvent
		value := args[0]
		incoming := value.Get("target")
		ret = MouseEventFromJS(value)
		src := MakeDocumentFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddAuxclick is adding doing AddEventListener for 'Auxclick' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventAuxclick(listener func(event *MouseEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "auxclick", cb)
	return cb
}

// SetOnAuxclick is assigning a function to 'onauxclick'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnAuxclick(listener func(event *MouseEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_MouseEvent(listener)
	_this.jsValue.Set("onauxclick", cb)
	return cb
}

// event attribute: FocusEvent
func eventFuncDocument_FocusEvent(listener func(event *FocusEvent, target *Document)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *FocusEvent
		value := args[0]
		incoming := value.Get("target")
		ret = FocusEventFromJS(value)
		src := MakeDocumentFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddBlur is adding doing AddEventListener for 'Blur' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventBlur(listener func(event *FocusEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_FocusEvent(listener)
	_this.jsValue.Call("addEventListener", "blur", cb)
	return cb
}

// SetOnBlur is assigning a function to 'onblur'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnBlur(listener func(event *FocusEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_FocusEvent(listener)
	_this.jsValue.Set("onblur", cb)
	return cb
}

// AddCancel is adding doing AddEventListener for 'Cancel' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventCancel(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "cancel", cb)
	return cb
}

// SetOnCancel is assigning a function to 'oncancel'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnCancel(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("oncancel", cb)
	return cb
}

// AddCanPlay is adding doing AddEventListener for 'CanPlay' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventCanPlay(listener func(event *Event, currentTarget *Document)) js.Func {
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "canplay", cb)
	return cb
}

// SetOnCanPlay is assigning a function to 'oncanplay'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnCanPlay(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("oncanplay", cb)
	return cb
}

// AddCanPlayThrough is adding doing AddEventListener for 'CanPlayThrough' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventCanPlayThrough(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "canplaythrough", cb)
	return cb
}

// SetOnCanPlayThrough is assigning a function to 'oncanplaythrough'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnCanPlayThrough(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("oncanplaythrough", cb)
	return cb
}

// AddChange is adding doing AddEventListener for 'Change' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventChange(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "change", cb)
	return cb
}

// SetOnChange is assigning a function to 'onchange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnChange(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onchange", cb)
	return cb
}

// AddClick is adding doing AddEventListener for 'Click' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventClick(listener func(event *MouseEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "click", cb)
	return cb
}

// SetOnClick is assigning a function to 'onclick'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnClick(listener func(event *MouseEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_MouseEvent(listener)
	_this.jsValue.Set("onclick", cb)
	return cb
}

// AddClose is adding doing AddEventListener for 'Close' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventClose(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "close", cb)
	return cb
}

// SetOnClose is assigning a function to 'onclose'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnClose(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onclose", cb)
	return cb
}

// AddContextMenu is adding doing AddEventListener for 'ContextMenu' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventContextMenu(listener func(event *MouseEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "contextmenu", cb)
	return cb
}

// SetOnContextMenu is assigning a function to 'oncontextmenu'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnContextMenu(listener func(event *MouseEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_MouseEvent(listener)
	_this.jsValue.Set("oncontextmenu", cb)
	return cb
}

// AddDblClick is adding doing AddEventListener for 'DblClick' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventDblClick(listener func(event *MouseEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "dblclick", cb)
	return cb
}

// SetOnDblClick is assigning a function to 'ondblclick'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnDblClick(listener func(event *MouseEvent, currentTarget *Document)) js.Func {
	cb := eventFuncDocument_MouseEvent(listener)
	_this.jsValue.Set("ondblclick", cb)
	return cb
}

// AddEmptied is adding doing AddEventListener for 'Emptied' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventEmptied(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "emptied", cb)
	return cb
}

// SetOnEmptied is assigning a function to 'onemptied'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnEmptied(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onemptied", cb)
	return cb
}

// AddEnded is adding doing AddEventListener for 'Ended' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventEnded(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "ended", cb)
	return cb
}

// SetOnEnded is assigning a function to 'onended'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnEnded(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onended", cb)
	return cb
}

// AddError is adding doing AddEventListener for 'Error' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventError(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "error", cb)
	return cb
}

// SetOnError is assigning a function to 'onerror'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnError(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onerror", cb)
	return cb
}

// AddFocus is adding doing AddEventListener for 'Focus' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventFocus(listener func(event *FocusEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_FocusEvent(listener)
	_this.jsValue.Call("addEventListener", "focus", cb)
	return cb
}

// SetOnFocus is assigning a function to 'onfocus'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnFocus(listener func(event *FocusEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_FocusEvent(listener)
	_this.jsValue.Set("onfocus", cb)
	return cb
}

// AddFullscreenChange is adding doing AddEventListener for 'FullscreenChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventFullscreenChange(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "fullscreenchange", cb)
	return cb
}

// SetOnFullscreenChange is assigning a function to 'onfullscreenchange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnFullscreenChange(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onfullscreenchange", cb)
	return cb
}

// AddFullscreenError is adding doing AddEventListener for 'FullscreenError' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventFullscreenError(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "fullscreenerror", cb)
	return cb
}

// SetOnFullscreenError is assigning a function to 'onfullscreenerror'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnFullscreenError(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onfullscreenerror", cb)
	return cb
}

// event attribute: PointerEvent
func eventFuncDocument_PointerEvent(listener func(event *PointerEvent, target *Document)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *PointerEvent
		value := args[0]
		incoming := value.Get("target")
		ret = PointerEventFromJS(value)
		src := MakeDocumentFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddGotPointerCapture is adding doing AddEventListener for 'GotPointerCapture' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventGotPointerCapture(listener func(event *PointerEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "gotpointercapture", cb)
	return cb
}

// SetOnGotPointerCapture is assigning a function to 'ongotpointercapture'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnGotPointerCapture(listener func(event *PointerEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_PointerEvent(listener)
	_this.jsValue.Set("ongotpointercapture", cb)
	return cb
}

// event attribute: InputEvent
func eventFuncDocument_InputEvent(listener func(event *InputEvent, target *Document)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *InputEvent
		value := args[0]
		incoming := value.Get("target")
		ret = InputEventFromJS(value)
		src := MakeDocumentFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddInput is adding doing AddEventListener for 'Input' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventInput(listener func(event *InputEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_InputEvent(listener)
	_this.jsValue.Call("addEventListener", "input", cb)
	return cb
}

// SetOnInput is assigning a function to 'oninput'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnInput(listener func(event *InputEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_InputEvent(listener)
	_this.jsValue.Set("oninput", cb)
	return cb
}

// AddInvalid is adding doing AddEventListener for 'Invalid' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventInvalid(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "invalid", cb)
	return cb
}

// SetOnInvalid is assigning a function to 'oninvalid'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnInvalid(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("oninvalid", cb)
	return cb
}

// event attribute: KeyboardEvent
func eventFuncDocument_KeyboardEvent(listener func(event *KeyboardEvent, target *Document)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *KeyboardEvent
		value := args[0]
		incoming := value.Get("target")
		ret = KeyboardEventFromJS(value)
		src := MakeDocumentFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddKeyDown is adding doing AddEventListener for 'KeyDown' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventKeyDown(listener func(event *KeyboardEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_KeyboardEvent(listener)
	_this.jsValue.Call("addEventListener", "keydown", cb)
	return cb
}

// SetOnKeyDown is assigning a function to 'onkeydown'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnKeyDown(listener func(event *KeyboardEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_KeyboardEvent(listener)
	_this.jsValue.Set("onkeydown", cb)
	return cb
}

// AddKeyPress is adding doing AddEventListener for 'KeyPress' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventKeyPress(listener func(event *KeyboardEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_KeyboardEvent(listener)
	_this.jsValue.Call("addEventListener", "keypress", cb)
	return cb
}

// SetOnKeyPress is assigning a function to 'onkeypress'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnKeyPress(listener func(event *KeyboardEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_KeyboardEvent(listener)
	_this.jsValue.Set("onkeypress", cb)
	return cb
}

// AddKeyUp is adding doing AddEventListener for 'KeyUp' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventKeyUp(listener func(event *KeyboardEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_KeyboardEvent(listener)
	_this.jsValue.Call("addEventListener", "keyup", cb)
	return cb
}

// SetOnKeyUp is assigning a function to 'onkeyup'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnKeyUp(listener func(event *KeyboardEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_KeyboardEvent(listener)
	_this.jsValue.Set("onkeyup", cb)
	return cb
}

// AddLoad is adding doing AddEventListener for 'Load' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventLoad(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "load", cb)
	return cb
}

// SetOnLoad is assigning a function to 'onload'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnLoad(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onload", cb)
	return cb
}

// AddLoadedData is adding doing AddEventListener for 'LoadedData' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventLoadedData(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "loadeddata", cb)
	return cb
}

// SetOnLoadedData is assigning a function to 'onloadeddata'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnLoadedData(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onloadeddata", cb)
	return cb
}

// AddLoadedMetaData is adding doing AddEventListener for 'LoadedMetaData' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventLoadedMetaData(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "loadedmetadata", cb)
	return cb
}

// SetOnLoadedMetaData is assigning a function to 'onloadedmetadata'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnLoadedMetaData(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onloadedmetadata", cb)
	return cb
}

// AddLoadStart is adding doing AddEventListener for 'LoadStart' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventLoadStart(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "loadstart", cb)
	return cb
}

// SetOnLoadStart is assigning a function to 'onloadstart'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnLoadStart(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onloadstart", cb)
	return cb
}

// AddLostPointerCapture is adding doing AddEventListener for 'LostPointerCapture' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventLostPointerCapture(listener func(event *PointerEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "lostpointercapture", cb)
	return cb
}

// SetOnLostPointerCapture is assigning a function to 'onlostpointercapture'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnLostPointerCapture(listener func(event *PointerEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_PointerEvent(listener)
	_this.jsValue.Set("onlostpointercapture", cb)
	return cb
}

// AddMouseDown is adding doing AddEventListener for 'MouseDown' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventMouseDown(listener func(event *MouseEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "mousedown", cb)
	return cb
}

// SetOnMouseDown is assigning a function to 'onmousedown'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnMouseDown(listener func(event *MouseEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_MouseEvent(listener)
	_this.jsValue.Set("onmousedown", cb)
	return cb
}

// AddMouseEnter is adding doing AddEventListener for 'MouseEnter' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventMouseEnter(listener func(event *MouseEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "mouseenter", cb)
	return cb
}

// SetOnMouseEnter is assigning a function to 'onmouseenter'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnMouseEnter(listener func(event *MouseEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_MouseEvent(listener)
	_this.jsValue.Set("onmouseenter", cb)
	return cb
}

// AddMouseLeave is adding doing AddEventListener for 'MouseLeave' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventMouseLeave(listener func(event *MouseEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "mouseleave", cb)
	return cb
}

// SetOnMouseLeave is assigning a function to 'onmouseleave'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnMouseLeave(listener func(event *MouseEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_MouseEvent(listener)
	_this.jsValue.Set("onmouseleave", cb)
	return cb
}

// AddMouseMove is adding doing AddEventListener for 'MouseMove' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventMouseMove(listener func(event *MouseEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "mousemove", cb)
	return cb
}

// SetOnMouseMove is assigning a function to 'onmousemove'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnMouseMove(listener func(event *MouseEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_MouseEvent(listener)
	_this.jsValue.Set("onmousemove", cb)
	return cb
}

// AddMouseOut is adding doing AddEventListener for 'MouseOut' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventMouseOut(listener func(event *MouseEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "mouseout", cb)
	return cb
}

// SetOnMouseOut is assigning a function to 'onmouseout'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnMouseOut(listener func(event *MouseEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_MouseEvent(listener)
	_this.jsValue.Set("onmouseout", cb)
	return cb
}

// AddMouseOver is adding doing AddEventListener for 'MouseOver' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventMouseOver(listener func(event *MouseEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "mouseover", cb)
	return cb
}

// SetOnMouseOver is assigning a function to 'onmouseover'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnMouseOver(listener func(event *MouseEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_MouseEvent(listener)
	_this.jsValue.Set("onmouseover", cb)
	return cb
}

// AddMouseUp is adding doing AddEventListener for 'MouseUp' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventMouseUp(listener func(event *MouseEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "mouseup", cb)
	return cb
}

// SetOnMouseUp is assigning a function to 'onmouseup'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnMouseUp(listener func(event *MouseEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_MouseEvent(listener)
	_this.jsValue.Set("onmouseup", cb)
	return cb
}

// AddPause is adding doing AddEventListener for 'Pause' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventPause(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "pause", cb)
	return cb
}

// SetOnPause is assigning a function to 'onpause'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnPause(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onpause", cb)
	return cb
}

// AddPlay is adding doing AddEventListener for 'Play' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventPlay(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "play", cb)
	return cb
}

// SetOnPlay is assigning a function to 'onplay'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnPlay(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onplay", cb)
	return cb
}

// AddPlaying is adding doing AddEventListener for 'Playing' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventPlaying(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "playing", cb)
	return cb
}

// SetOnPlaying is assigning a function to 'onplaying'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnPlaying(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onplaying", cb)
	return cb
}

// AddPointerCancel is adding doing AddEventListener for 'PointerCancel' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventPointerCancel(listener func(event *PointerEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointercancel", cb)
	return cb
}

// SetOnPointerCancel is assigning a function to 'onpointercancel'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnPointerCancel(listener func(event *PointerEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_PointerEvent(listener)
	_this.jsValue.Set("onpointercancel", cb)
	return cb
}

// AddPointerDown is adding doing AddEventListener for 'PointerDown' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventPointerDown(listener func(event *PointerEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointerdown", cb)
	return cb
}

// SetOnPointerDown is assigning a function to 'onpointerdown'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnPointerDown(listener func(event *PointerEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_PointerEvent(listener)
	_this.jsValue.Set("onpointerdown", cb)
	return cb
}

// AddPointerEnter is adding doing AddEventListener for 'PointerEnter' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventPointerEnter(listener func(event *PointerEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointerenter", cb)
	return cb
}

// SetOnPointerEnter is assigning a function to 'onpointerenter'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnPointerEnter(listener func(event *PointerEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_PointerEvent(listener)
	_this.jsValue.Set("onpointerenter", cb)
	return cb
}

// AddPointerLeave is adding doing AddEventListener for 'PointerLeave' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventPointerLeave(listener func(event *PointerEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointerleave", cb)
	return cb
}

// SetOnPointerLeave is assigning a function to 'onpointerleave'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnPointerLeave(listener func(event *PointerEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_PointerEvent(listener)
	_this.jsValue.Set("onpointerleave", cb)
	return cb
}

// AddPointerLockChange is adding doing AddEventListener for 'PointerLockChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventPointerLockChange(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "pointerlockchange", cb)
	return cb
}

// SetOnPointerLockChange is assigning a function to 'onpointerlockchange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnPointerLockChange(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onpointerlockchange", cb)
	return cb
}

// AddPointerLockError is adding doing AddEventListener for 'PointerLockError' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventPointerLockError(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "pointerlockerror", cb)
	return cb
}

// SetOnPointerLockError is assigning a function to 'onpointerlockerror'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnPointerLockError(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onpointerlockerror", cb)
	return cb
}

// AddPointerMove is adding doing AddEventListener for 'PointerMove' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventPointerMove(listener func(event *PointerEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointermove", cb)
	return cb
}

// SetOnPointerMove is assigning a function to 'onpointermove'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnPointerMove(listener func(event *PointerEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_PointerEvent(listener)
	_this.jsValue.Set("onpointermove", cb)
	return cb
}

// AddPointerOut is adding doing AddEventListener for 'PointerOut' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventPointerOut(listener func(event *PointerEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointerout", cb)
	return cb
}

// SetOnPointerOut is assigning a function to 'onpointerout'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnPointerOut(listener func(event *PointerEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_PointerEvent(listener)
	_this.jsValue.Set("onpointerout", cb)
	return cb
}

// AddPointerOver is adding doing AddEventListener for 'PointerOver' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventPointerOver(listener func(event *PointerEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointerover", cb)
	return cb
}

// SetOnPointerOver is assigning a function to 'onpointerover'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnPointerOver(listener func(event *PointerEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_PointerEvent(listener)
	_this.jsValue.Set("onpointerover", cb)
	return cb
}

// AddPointerUp is adding doing AddEventListener for 'PointerUp' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventPointerUp(listener func(event *PointerEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointerup", cb)
	return cb
}

// SetOnPointerUp is assigning a function to 'onpointerup'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnPointerUp(listener func(event *PointerEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_PointerEvent(listener)
	_this.jsValue.Set("onpointerup", cb)
	return cb
}

// AddRateChange is adding doing AddEventListener for 'RateChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventRateChange(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "ratechange", cb)
	return cb
}

// SetOnRateChange is assigning a function to 'onratechange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnRateChange(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onratechange", cb)
	return cb
}

// AddReadyStateChange is adding doing AddEventListener for 'ReadyStateChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventReadyStateChange(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "readystatechange", cb)
	return cb
}

// SetOnReadyStateChange is assigning a function to 'onreadystatechange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnReadyStateChange(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onreadystatechange", cb)
	return cb
}

// AddReset is adding doing AddEventListener for 'Reset' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventReset(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "reset", cb)
	return cb
}

// SetOnReset is assigning a function to 'onreset'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnReset(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onreset", cb)
	return cb
}

// event attribute: UIEvent
func eventFuncDocument_UIEvent(listener func(event *UIEvent, target *Document)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *UIEvent
		value := args[0]
		incoming := value.Get("target")
		ret = UIEventFromJS(value)
		src := MakeDocumentFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddResize is adding doing AddEventListener for 'Resize' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventResize(listener func(event *UIEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_UIEvent(listener)
	_this.jsValue.Call("addEventListener", "resize", cb)
	return cb
}

// SetOnResize is assigning a function to 'onresize'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnResize(listener func(event *UIEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_UIEvent(listener)
	_this.jsValue.Set("onresize", cb)
	return cb
}

// AddScroll is adding doing AddEventListener for 'Scroll' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventScroll(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "scroll", cb)
	return cb
}

// SetOnScroll is assigning a function to 'onscroll'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnScroll(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onscroll", cb)
	return cb
}

// AddSeeked is adding doing AddEventListener for 'Seeked' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventSeeked(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "seeked", cb)
	return cb
}

// SetOnSeeked is assigning a function to 'onseeked'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnSeeked(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onseeked", cb)
	return cb
}

// AddSeeking is adding doing AddEventListener for 'Seeking' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventSeeking(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "seeking", cb)
	return cb
}

// SetOnSeeking is assigning a function to 'onseeking'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnSeeking(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onseeking", cb)
	return cb
}

// AddSelect is adding doing AddEventListener for 'Select' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventSelect(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "select", cb)
	return cb
}

// SetOnSelect is assigning a function to 'onselect'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnSelect(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onselect", cb)
	return cb
}

// AddSelectionChange is adding doing AddEventListener for 'SelectionChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventSelectionChange(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "selectionchange", cb)
	return cb
}

// SetOnSelectionChange is assigning a function to 'onselectionchange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnSelectionChange(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onselectionchange", cb)
	return cb
}

// AddSelectStart is adding doing AddEventListener for 'SelectStart' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventSelectStart(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "selectstart", cb)
	return cb
}

// SetOnSelectStart is assigning a function to 'onselectstart'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnSelectStart(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onselectstart", cb)
	return cb
}

// AddStalled is adding doing AddEventListener for 'Stalled' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventStalled(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "stalled", cb)
	return cb
}

// SetOnStalled is assigning a function to 'onstalled'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnStalled(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onstalled", cb)
	return cb
}

// AddSubmit is adding doing AddEventListener for 'Submit' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventSubmit(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "submit", cb)
	return cb
}

// SetOnSubmit is assigning a function to 'onsubmit'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnSubmit(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onsubmit", cb)
	return cb
}

// AddSuspend is adding doing AddEventListener for 'Suspend' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventSuspend(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "suspend", cb)
	return cb
}

// SetOnSuspend is assigning a function to 'onsuspend'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnSuspend(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onsuspend", cb)
	return cb
}

// AddTimeUpdate is adding doing AddEventListener for 'TimeUpdate' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventTimeUpdate(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "timeupdate", cb)
	return cb
}

// SetOnTimeUpdate is assigning a function to 'ontimeupdate'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnTimeUpdate(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("ontimeupdate", cb)
	return cb
}

// AddToggle is adding doing AddEventListener for 'Toggle' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventToggle(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "toggle", cb)
	return cb
}

// SetOnToggle is assigning a function to 'ontoggle'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnToggle(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("ontoggle", cb)
	return cb
}

// AddVisibilityChange is adding doing AddEventListener for 'VisibilityChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventVisibilityChange(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "visibilitychange", cb)
	return cb
}

// SetOnVisibilityChange is assigning a function to 'onvisibilitychange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnVisibilityChange(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onvisibilitychange", cb)
	return cb
}

// AddVolumeChange is adding doing AddEventListener for 'VolumeChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventVolumeChange(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "volumechange", cb)
	return cb
}

// SetOnVolumeChange is assigning a function to 'onvolumechange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnVolumeChange(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onvolumechange", cb)
	return cb
}

// AddWaiting is adding doing AddEventListener for 'Waiting' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventWaiting(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Call("addEventListener", "waiting", cb)
	return cb
}

// SetOnWaiting is assigning a function to 'onwaiting'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnWaiting(listener func(event *Event, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_Event(listener)
	_this.jsValue.Set("onwaiting", cb)
	return cb
}

// event attribute: WheelEvent
func eventFuncDocument_WheelEvent(listener func(event *WheelEvent, target *Document)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *WheelEvent
		value := args[0]
		incoming := value.Get("target")
		ret = WheelEventFromJS(value)
		src := MakeDocumentFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddWheel is adding doing AddEventListener for 'Wheel' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Document) AddEventWheel(listener func(event *WheelEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("AddEvent__() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_WheelEvent(listener)
	_this.jsValue.Call("addEventListener", "wheel", cb)
	return cb
}

// SetOnWheel is assigning a function to 'onwheel'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Document) SetOnWheel(listener func(event *WheelEvent, currentTarget *Document)) js.Func {
	if _this == nil {
		log.Println("SetOn{event}() call on a nil Document")
		return js.Func{Value: js.Null()}
	}
	cb := eventFuncDocument_WheelEvent(listener)
	_this.jsValue.Set("onwheel", cb)
	return cb
}

/******************************************************************************
* Document's method
******************************************************************************/

// GetElementsByTagName returns an HTMLCollection of elements with the given tag name.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByTagName
func (_this *Document) GetElementsByTagName(qualifiedName string) (_result *HTMLCollection) {
	if _this == nil {
		log.Println("GetElementsByTagName() call on a nil Document")
		return nil
	}
	var _args [1]interface{}
	_args[0] = qualifiedName
	_returned := _this.jsValue.Call("getElementsByTagName", _args[0:1]...)
	return HTMLCollectionFromJS(_returned)
}

// GetElementsByClassName returns an array-like object of all child elements which have all of the given class name(s).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByClassName
func (_this *Document) GetElementsByClassName(classNames string) (_result *HTMLCollection) {
	if _this == nil {
		log.Println("GetElementsByClassName() call on a nil Document")
		return nil
	}
	var _args [1]interface{}
	_args[0] = classNames
	_returned := _this.jsValue.Call("getElementsByClassName", _args[0:1]...)
	return HTMLCollectionFromJS(_returned)
}

// GetElementsByName returns a NodeList Collection of elements with a given name attribute in the document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByName
func (_this *Document) GetElementsByName(elementName string) (_result *NodeList) {
	if _this == nil {
		log.Println("GetElementsByName() call on a nil Document")
		return nil
	}
	var _args [1]interface{}
	_args[0] = elementName
	_returned := _this.jsValue.Call("getElementsByName", _args[0:1]...)
	return MakeNodeListFromJS(_returned)
}

// GetElementById returns an Element object representing the element whose id property matches the specified string.
// Since element IDs are required to be unique if specified, they're a useful way to get access to a specific element quickly.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById
func (_this *Document) GetElementById(elementId string) (_result *Element) {
	if _this == nil {
		log.Println("GetElementsByName() call on a nil Document")
		return nil
	}
	var _args [1]interface{}
	_args[0] = elementId
	_returned := _this.jsValue.Call("getElementById", _args[0:1]...)
	return MakeElementFromJS(_returned)
}

// QuerySelector returns the first Element within the document that matches the specified selector, or group of selectors.
// If no matches are found, null is returned.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector
func (_this *Document) QuerySelector(selectors string) (_result *Element) {
	if _this == nil {
		log.Println("QuerySelector() call on a nil Document")
		return nil
	}
	var _args [1]interface{}
	_args[0] = selectors
	_returned := _this.jsValue.Call("querySelector", _args[0:1]...)
	return MakeElementFromJS(_returned)
}

// querySelectorAll returns a static (not live) NodeList representing a list of the document's elements that match the specified group of selectors.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelectorAll
func (_this *Document) QuerySelectorAll(selectors string) (_result *NodeList) {
	if _this == nil {
		log.Println("QuerySelectorAll() call on a nil Document")
		return nil
	}
	var _args [1]interface{}
	_args[0] = selectors
	_returned := _this.jsValue.Call("querySelectorAll", _args[0:1]...)
	return MakeNodeListFromJS(_returned)
}

// CreateElement creates the HTML element specified by tagName, or an HTMLUnknownElement if tagName isn't recognized.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement
func (_this *Document) CreateElement(localName string) (_result *Element) {
	if _this == nil {
		log.Println("CreateElement() call on a nil Document")
		return nil
	}
	var _args [1]interface{}
	_args[0] = localName
	_returned := _this.jsValue.Call("createElement", _args[0:1]...)
	_result = MakeElementFromJS(_returned)
	return _result
}

// CreateAttribute  creates a new attribute node, and returns it. The object created is a node implementing the Attr interface.
// The DOM does not enforce what sort of attributes can be added to a particular element in this manner.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/createAttribute
func (_this *Document) CreateAttribute(localName string) (_result *Attribute) {
	if _this == nil {
		log.Println("CreateAttribute() call on a nil Document")
		return nil
	}
	var _args [1]interface{}
	_args[0] = localName
	_returned := _this.jsValue.Call("createAttribute", _args[0:1]...)
	return MakeAttributeFromJS(_returned)
}

// CreateNodeIterator has been removed and replaced with the MakeNodes() function
// func (_this *Document) CreateNodeIterator(root *Node, whatToShow *uint, filter *NodeFilter) *Nodes {
// }

// GetElementAtPoint returns the topmost Element at the specified coordinates (relative to the viewport).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/elementFromPoint
func (_this *Document) GetElementAtPoint(x float64, y float64) (_result *Element) {
	if _this == nil {
		log.Println("GetElementAtPoint() call on a nil Document")
		return nil
	}
	var _args [2]interface{}
	_args[0] = x
	_args[1] = y
	_returned := _this.jsValue.Call("elementFromPoint", _args[0:2]...)
	return MakeElementFromJS(_returned)
}

// GetElementsAtPoint eturns an array of all elements at the specified coordinates (relative to the viewport).
// The elements are ordered from the topmost to the bottommost box of the viewport.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/elementsFromPoint
func (_this *Document) GetElementsAtPoint(x float64, y float64) (_result []*Element) {
	if _this == nil {
		log.Println("GetElementsAtPoint() call on a nil Document")
		return nil
	}
	var _args [2]interface{}
	_args[0] = x
	_args[1] = y
	_returned := _this.jsValue.Call("elementsFromPoint", _args[0:2]...)

	__length0 := _returned.Length()
	__array0 := make([]*Element, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		__seq_in0 := _returned.Index(__idx0)
		__array0[__idx0] = MakeElementFromJS(__seq_in0)
	}
	return __array0
}

// HasFocus returns a boolean value indicating whether the document or any element inside the document has focus.
// This method can be used to determine whether the active element in a document has focus.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/hasFocus
func (_this *Document) HasFocus() (_result bool) {
	if _this == nil {
		log.Println("HasFocus() call on a nil Document")
		return false
	}
	var _args [0]interface{}
	_returned := _this.jsValue.Call("hasFocus", _args[0:0]...)
	return (_returned).Bool()
}

// Prepend inserts a set of Node objects or string objects before the first child of the document.
// String objects are inserted as equivalent Text nodes.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/prepend
func (_this *Document) Prepend(nodes ...*Union) {
	if _this == nil {
		log.Println("Prepend() call on a nil Document")
		return
	}
	var (
		_args []interface{} = make([]interface{}, 0+len(nodes))
		_end  int
	)
	for _, __in := range nodes {
		__out := __in.JSValue()
		_args[_end] = __out
		_end++
	}
	_this.jsValue.Call("prepend", _args[0:_end]...)
}

// Append inserts a set of Node objects or string objects after the last child of the document.
// String objects are inserted as equivalent Text nodes.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/append
func (_this *Document) Append(nodes ...*Union) {
	if _this == nil {
		log.Println("Append() call on a nil Document")
		return
	}
	var (
		_args []interface{} = make([]interface{}, 0+len(nodes))
		_end  int
	)
	for _, __in := range nodes {
		__out := __in.JSValue()
		_args[_end] = __out
		_end++
	}
	_this.jsValue.Call("append", _args[0:_end]...)
}
