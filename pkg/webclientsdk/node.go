package webclientsdk

import (
	"syscall/js"
)

/****************************************************************************
* NodeFilter
*****************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/Node
type Node struct {
	EventTarget
}

// NodeFromJS is casting a js.Value into Node.
func NodeFromJS(value js.Value) *Node {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Node{}
	ret.jsValue = value
	return ret
}

const (
	ELEMENT_NODE                              int = 1
	ATTRIBUTE_NODE                            int = 2
	TEXT_NODE                                 int = 3
	CDATA_SECTION_NODE                        int = 4
	ENTITY_REFERENCE_NODE                     int = 5
	ENTITY_NODE                               int = 6
	PROCESSING_INSTRUCTION_NODE               int = 7
	COMMENT_NODE                              int = 8
	DOCUMENT_NODE                             int = 9
	DOCUMENT_TYPE_NODE                        int = 10
	DOCUMENT_FRAGMENT_NODE                    int = 11
	NOTATION_NODE                             int = 12
	DOCUMENT_POSITION_DISCONNECTED            int = 0x01
	DOCUMENT_POSITION_PRECEDING               int = 0x02
	DOCUMENT_POSITION_FOLLOWING               int = 0x04
	DOCUMENT_POSITION_CONTAINS                int = 0x08
	DOCUMENT_POSITION_CONTAINED_BY            int = 0x10
	DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC int = 0x20
)

// NodeType returning attribute 'nodeType' with
// type int (idl: unsigned short).
func (_this *Node) NodeType() int {
	var ret int
	value := _this.jsValue.Get("nodeType")
	ret = (value).Int()
	return ret
}

// NodeName returning attribute 'nodeName' with
// type string (idl: DOMString).
func (_this *Node) NodeName() string {
	var ret string
	value := _this.jsValue.Get("nodeName")
	ret = (value).String()
	return ret
}

// BaseURI returning attribute 'baseURI' with
// type string (idl: USVString).
func (_this *Node) BaseURI() string {
	var ret string
	value := _this.jsValue.Get("baseURI")
	ret = (value).String()
	return ret
}

// IsConnected returning attribute 'isConnected' with
// type bool (idl: boolean).
func (_this *Node) IsConnected() bool {
	var ret bool
	value := _this.jsValue.Get("isConnected")
	ret = (value).Bool()
	return ret
}

// OwnerDocument returning attribute 'ownerDocument' with
// type js.Value (idl: Document).
func (_this *Node) OwnerDocument() js.Value {
	var ret js.Value
	value := _this.jsValue.Get("ownerDocument")
	ret = value
	return ret
}

// ParentNode returning attribute 'parentNode' with
// type Node (idl: Node).
func (_this *Node) ParentNode() *Node {
	var ret *Node
	value := _this.jsValue.Get("parentNode")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = NodeFromJS(value)
	}
	return ret
}

// ParentElement returning attribute 'parentElement' with
// type Element (idl: Element).
func (_this *Node) ParentElement() *Element {
	var ret *Element
	value := _this.jsValue.Get("parentElement")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = ElementFromJS(value)
	}
	return ret
}

// ChildNodes returning attribute 'childNodes' with
// type NodeList (idl: NodeList).
func (_this *Node) ChildNodes() *NodeList {
	var ret *NodeList
	value := _this.jsValue.Get("childNodes")
	ret = NodeListFromJS(value)
	return ret
}

// FirstChild returning attribute 'firstChild' with
// type Node (idl: Node).
func (_this *Node) FirstChild() *Node {
	var ret *Node
	value := _this.jsValue.Get("firstChild")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = NodeFromJS(value)
	}
	return ret
}

// LastChild returning attribute 'lastChild' with
// type Node (idl: Node).
func (_this *Node) LastChild() *Node {
	var ret *Node
	value := _this.jsValue.Get("lastChild")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = NodeFromJS(value)
	}
	return ret
}

// PreviousSibling returning attribute 'previousSibling' with
// type Node (idl: Node).
func (_this *Node) PreviousSibling() *Node {
	var ret *Node
	value := _this.jsValue.Get("previousSibling")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = NodeFromJS(value)
	}
	return ret
}

// NextSibling returning attribute 'nextSibling' with
// type Node (idl: Node).
func (_this *Node) NextSibling() *Node {
	var ret *Node
	value := _this.jsValue.Get("nextSibling")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = NodeFromJS(value)
	}
	return ret
}

// NodeValue returning attribute 'nodeValue' with
// type string (idl: DOMString).
func (_this *Node) NodeValue() *string {
	var ret *string
	value := _this.jsValue.Get("nodeValue")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetNodeValue setting attribute 'nodeValue' with
// type string (idl: DOMString).
func (_this *Node) SetNodeValue(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("nodeValue", input)
}

// TextContent returning attribute 'textContent' with
// type string (idl: DOMString).
func (_this *Node) TextContent() *string {
	var ret *string
	value := _this.jsValue.Get("textContent")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetTextContent setting attribute 'textContent' with
// type string (idl: DOMString).
func (_this *Node) SetTextContent(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("textContent", input)
}

func (_this *Node) GetRootNode() (_result *Node) {
	var (
		_args [1]interface{}
		_end  int
	)
	_returned := _this.jsValue.Call("getRootNode", _args[0:_end]...)
	var (
		_converted *Node // javascript: Node _what_return_name
	)
	_converted = NodeFromJS(_returned)
	_result = _converted
	return
}

func (_this *Node) HasChildNodes() (_result bool) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.jsValue.Call("hasChildNodes", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *Node) Normalize() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.jsValue.Call("normalize", _args[0:_end]...)
	return
}

func (_this *Node) CloneNode(deep *bool) (_result *Node) {
	var (
		_args [1]interface{}
		_end  int
	)
	if deep != nil {

		var _p0 interface{}
		if deep != nil {
			_p0 = *(deep)
		} else {
			_p0 = nil
		}
		_args[0] = _p0
		_end++
	}
	_returned := _this.jsValue.Call("cloneNode", _args[0:_end]...)
	var (
		_converted *Node // javascript: Node _what_return_name
	)
	_converted = NodeFromJS(_returned)
	_result = _converted
	return
}

func (_this *Node) IsEqualNode(otherNode *Node) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := otherNode.jsValue
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("isEqualNode", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *Node) IsSameNode(otherNode *Node) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := otherNode.jsValue
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("isSameNode", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *Node) CompareDocumentPosition(other *Node) (_result int) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := other.jsValue
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("compareDocumentPosition", _args[0:_end]...)
	var (
		_converted int // javascript: unsigned short _what_return_name
	)
	_converted = (_returned).Int()
	_result = _converted
	return
}

func (_this *Node) Contains(other *Node) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := other.jsValue
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("contains", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *Node) InsertBefore(node *Node, child *Node) (_result *Node) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := node.jsValue
	_args[0] = _p0
	_end++
	_p1 := child.jsValue
	_args[1] = _p1
	_end++
	_returned := _this.jsValue.Call("insertBefore", _args[0:_end]...)
	var (
		_converted *Node // javascript: Node _what_return_name
	)
	_converted = NodeFromJS(_returned)
	_result = _converted
	return
}

func (_this *Node) AppendChild(node *Node) (_result *Node) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := node.jsValue
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("appendChild", _args[0:_end]...)
	var (
		_converted *Node // javascript: Node _what_return_name
	)
	_converted = NodeFromJS(_returned)
	_result = _converted
	return
}

func (_this *Node) ReplaceChild(node *Node, child *Node) (_result *Node) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := node.jsValue
	_args[0] = _p0
	_end++
	_p1 := child.jsValue
	_args[1] = _p1
	_end++
	_returned := _this.jsValue.Call("replaceChild", _args[0:_end]...)
	var (
		_converted *Node // javascript: Node _what_return_name
	)
	_converted = NodeFromJS(_returned)
	_result = _converted
	return
}

func (_this *Node) RemoveChild(child *Node) (_result *Node) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := child.jsValue
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("removeChild", _args[0:_end]...)
	var (
		_converted *Node // javascript: Node _what_return_name
	)
	_converted = NodeFromJS(_returned)
	_result = _converted
	return
}

/****************************************************************************
* NodeFilter
*****************************************************************************/

const (
	FILTER_ACCEPT               int  = 1
	FILTER_REJECT               int  = 2
	FILTER_SKIP                 int  = 3
	SHOW_ALL                    uint = 0xFFFFFFFF
	SHOW_ELEMENT                uint = 0x1
	SHOW_ATTRIBUTE              uint = 0x2
	SHOW_TEXT                   uint = 0x4
	SHOW_CDATA_SECTION          uint = 0x8
	SHOW_ENTITY_REFERENCE       uint = 0x10
	SHOW_ENTITY                 uint = 0x20
	SHOW_PROCESSING_INSTRUCTION uint = 0x40
	SHOW_COMMENT                uint = 0x80
	SHOW_DOCUMENT               uint = 0x100
	SHOW_DOCUMENT_TYPE          uint = 0x200
	SHOW_DOCUMENT_FRAGMENT      uint = 0x400
	SHOW_NOTATION               uint = 0x800
)

// NodeFilter is a callback interface.
type NodeFilter interface {
	AcceptNode(node *Node) (_result int)
}

// NodeFilterValue is javascript reference value for callback interface NodeFilter.
// This is holding the underlying javascript object.
type NodeFilterValue struct {
	// Value is the underlying javascript object or function.
	jsValue js.Value

	// Functions is the underlying function objects that is allocated for the interface callback
	Functions [1]js.Func

	// Go interface to invoke
	impl      NodeFilter
	function  func(node *Node) (_result int)
	useInvoke bool
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *NodeFilterValue) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.jsValue
}

// Release is releasing all resources that is allocated.
func (t *NodeFilterValue) Release() {
	for i := range t.Functions {
		if t.Functions[i].Type() != js.TypeUndefined {
			t.Functions[i].Release()
		}
	}
}

// NewNodeFilter is allocating a new javascript object that
// implements NodeFilter.
func NewNodeFilter(callback NodeFilter) *NodeFilterValue {
	ret := &NodeFilterValue{impl: callback}
	ret.jsValue = js.Global().Get("Object").New()
	ret.Functions[0] = ret.allocateAcceptNode()
	ret.jsValue.Set("acceptNode", ret.Functions[0])
	return ret
}

// NewNodeFilterFunc is allocating a new javascript
// function is implements
// NodeFilter interface.
func NewNodeFilterFunc(f func(node *Node) (_result int)) *NodeFilterValue {
	// single function will result in javascript function type, not an object
	ret := &NodeFilterValue{function: f}
	ret.Functions[0] = ret.allocateAcceptNode()
	ret.jsValue = ret.Functions[0].Value
	return ret
}

// NodeFilterFromJS is taking an javascript object that reference to a
// callback interface and return a corresponding interface that can be used
// to invoke on that element.
func NodeFilterFromJS(value js.Value) *NodeFilterValue {
	if value.Type() == js.TypeObject {
		return &NodeFilterValue{jsValue: value}
	}
	if value.Type() == js.TypeFunction {
		return &NodeFilterValue{jsValue: value, useInvoke: true}
	}
	panic("unsupported type")
}

func (t *NodeFilterValue) allocateAcceptNode() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *Node // javascript: Node node
		)
		_p0 = NodeFromJS(args[0])
		var _returned int
		if t.function != nil {
			_returned = t.function(_p0)
		} else {
			_returned = t.impl.AcceptNode(_p0)
		}
		_converted := _returned
		return _converted
	})
}

func (_this *NodeFilterValue) AcceptNode(node *Node) (_result int) {
	if _this.function != nil {
		return _this.function(node)
	}
	if _this.impl != nil {
		return _this.impl.AcceptNode(node)
	}
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := node.jsValue
	_args[0] = _p0
	_end++
	var _returned js.Value
	if _this.useInvoke {

		// invoke a javascript function
		_returned = _this.jsValue.Invoke(_args[0:_end]...)
	} else {
		_returned = _this.jsValue.Call("acceptNode", _args[0:_end]...)
	}
	var (
		_converted int // javascript: unsigned short _what_return_name
	)
	_converted = (_returned).Int()
	_result = _converted
	return
}

/****************************************************************************
* NodeIterator
*****************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/NodeIterator
type NodeIterator struct {
	jsValue js.Value
}

// NodeIteratorFromJS is casting a js.Value into NodeIterator.
func NodeIteratorFromJS(value js.Value) *NodeIterator {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &NodeIterator{}
	ret.jsValue = value
	return ret
}

// Root returning attribute 'root' with
// type Node (idl: Node).
func (_this *NodeIterator) Root() *Node {
	var ret *Node
	value := _this.jsValue.Get("root")
	ret = NodeFromJS(value)
	return ret
}

// ReferenceNode returning attribute 'referenceNode' with
// type Node (idl: Node).
func (_this *NodeIterator) ReferenceNode() *Node {
	var ret *Node
	value := _this.jsValue.Get("referenceNode")
	ret = NodeFromJS(value)
	return ret
}

// PointerBeforeReferenceNode returning attribute 'pointerBeforeReferenceNode' with
// type bool (idl: boolean).
func (_this *NodeIterator) PointerBeforeReferenceNode() bool {
	var ret bool
	value := _this.jsValue.Get("pointerBeforeReferenceNode")
	ret = (value).Bool()
	return ret
}

// WhatToShow returning attribute 'whatToShow' with
// type uint (idl: unsigned long).
func (_this *NodeIterator) WhatToShow() uint {
	var ret uint
	value := _this.jsValue.Get("whatToShow")
	ret = (uint)((value).Int())
	return ret
}

// Filter returning attribute 'filter' with
// type NodeFilter (idl: NodeFilter).
func (_this *NodeIterator) Filter() NodeFilter {
	var ret NodeFilter
	value := _this.jsValue.Get("filter")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = NodeFilterFromJS(value)
	}
	return ret
}

func (_this *NodeIterator) NextNode() (_result *Node) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.jsValue.Call("nextNode", _args[0:_end]...)
	var (
		_converted *Node // javascript: Node _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = NodeFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *NodeIterator) PreviousNode() (_result *Node) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.jsValue.Call("previousNode", _args[0:_end]...)
	var (
		_converted *Node // javascript: Node _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = NodeFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *NodeIterator) Detach() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.jsValue.Call("detach", _args[0:_end]...)
	return
}

/****************************************************************************
* NodeListForEach
*****************************************************************************/

type NodeListForEachFunc func(currentValue *Node, currentIndex int, listObj *NodeList)

// NodeListForEach is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type NodeListForEach js.Func

func NodeListForEachToJS(callback NodeListForEachFunc) *NodeListForEach {
	if callback == nil {
		return nil
	}
	ret := NodeListForEach(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *Node     // javascript: Node currentValue
			_p1 int       // javascript: long currentIndex
			_p2 *NodeList // javascript: NodeList listObj
		)
		_p0 = NodeFromJS(args[0])
		_p1 = (args[1]).Int()
		_p2 = NodeListFromJS(args[2])
		callback(_p0, _p1, _p2)

		// returning no return value
		return nil
	}))
	return &ret
}

func NodeListForEachFromJS(_value js.Value) NodeListForEachFunc {
	return func(currentValue *Node, currentIndex int, listObj *NodeList) {
		var (
			_args [3]interface{}
			_end  int
		)
		_p0 := currentValue.jsValue
		_args[0] = _p0
		_end++
		_p1 := currentIndex
		_args[1] = _p1
		_end++
		_p2 := listObj.jsValue
		_args[2] = _p2
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

/****************************************************************************
* NodeList
*****************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/NodeList
type NodeList struct {
	// Value holds a reference to a javascript value
	jsValue js.Value
}

// NodeListFromJS is casting a js.Value into NodeList.
func NodeListFromJS(value js.Value) *NodeList {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &NodeList{}
	ret.jsValue = value
	return ret
}

// Length returning attribute 'length' with
// type uint (idl: unsigned long).
func (_this *NodeList) Length() uint {
	var ret uint
	value := _this.jsValue.Get("length")
	ret = (uint)((value).Int())
	return ret
}

func (_this *NodeList) Index(index uint) (_result *Node) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("item", _args[0:_end]...)
	var (
		_converted *Node // javascript: Node _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = NodeFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *NodeList) Item(index uint) (_result *Node) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("item", _args[0:_end]...)
	var (
		_converted *Node // javascript: Node _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = NodeFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *NodeList) Entries() (_result *NodeListEntryIterator) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.jsValue.Call("entries", _args[0:_end]...)
	var (
		_converted *NodeListEntryIterator // javascript: NodeListEntryIterator _what_return_name
	)
	_converted = NodeListEntryIteratorFromJS(_returned)
	_result = _converted
	return
}

func (_this *NodeList) ForEach(callback *NodeListForEach, optionalThisForCallbackArgument interface{}) {
	var (
		_args [2]interface{}
		_end  int
	)

	var __callback0 js.Value
	if callback != nil {
		__callback0 = (*callback).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	if optionalThisForCallbackArgument != nil {
		_p1 := optionalThisForCallbackArgument
		_args[1] = _p1
		_end++
	}
	_this.jsValue.Call("forEach", _args[0:_end]...)
	return
}

func (_this *NodeList) Keys() (_result *NodeListKeyIterator) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.jsValue.Call("keys", _args[0:_end]...)
	var (
		_converted *NodeListKeyIterator // javascript: NodeListKeyIterator _what_return_name
	)
	_converted = NodeListKeyIteratorFromJS(_returned)
	_result = _converted
	return
}

func (_this *NodeList) Values() (_result *NodeListValueIterator) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.jsValue.Call("values", _args[0:_end]...)
	var (
		_converted *NodeListValueIterator // javascript: NodeListValueIterator _what_return_name
	)
	_converted = NodeListValueIteratorFromJS(_returned)
	_result = _converted
	return
}

/****************************************************************************
* NodeListEntryIteratorValue
*****************************************************************************/

// dictionary: NodeListEntryIteratorValue
type NodeListEntryIteratorValue struct {
	Value []js.Value
	Done  bool
}

// NodeListEntryIteratorValueFromJS is allocating a new
// NodeListEntryIteratorValue object and copy all values in the value javascript object.
func NodeListEntryIteratorValueFromJS(value js.Value) *NodeListEntryIteratorValue {
	var out NodeListEntryIteratorValue
	var (
		value0 []js.Value // javascript: sequence<any> {value Value value}
		value1 bool       // javascript: boolean {done Done done}
	)
	__length0 := value.Get("value").Length()
	__array0 := make([]js.Value, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 js.Value
		__seq_in0 := value.Get("value").Index(__idx0)
		__seq_out0 = __seq_in0
		__array0[__idx0] = __seq_out0
	}
	value0 = __array0
	out.Value = value0
	value1 = (value.Get("done")).Bool()
	out.Done = value1
	return &out
}

/****************************************************************************
* NodeListKeyIteratorValue
*****************************************************************************/

type NodeListKeyIteratorValue struct {
	Value uint
	Done  bool
}

// NodeListKeyIteratorValueFromJS is allocating a new
// NodeListKeyIteratorValue object and copy all values in the value javascript object.
func NodeListKeyIteratorValueFromJS(value js.Value) *NodeListKeyIteratorValue {
	var out NodeListKeyIteratorValue
	var (
		value0 uint // javascript: unsigned long {value Value value}
		value1 bool // javascript: boolean {done Done done}
	)
	value0 = (uint)((value.Get("value")).Int())
	out.Value = value0
	value1 = (value.Get("done")).Bool()
	out.Done = value1
	return &out
}

/****************************************************************************
* NodeListValueIteratorValue
*****************************************************************************/

type NodeListValueIteratorValue struct {
	Value *Node
	Done  bool
}

// NodeListValueIteratorValueFromJS is allocating a new
// NodeListValueIteratorValue object and copy all values in the value javascript object.
func NodeListValueIteratorValueFromJS(value js.Value) *NodeListValueIteratorValue {
	var out NodeListValueIteratorValue
	var (
		value0 *Node // javascript: Node {value Value value}
		value1 bool  // javascript: boolean {done Done done}
	)
	value0 = NodeFromJS(value.Get("value"))
	out.Value = value0
	value1 = (value.Get("done")).Bool()
	out.Done = value1
	return &out
}

/****************************************************************************
* NodeListEntryIterator
*****************************************************************************/

type NodeListEntryIterator struct {
	jsValue js.Value
}

// NodeListEntryIteratorFromJS is casting a js.Value into NodeListEntryIterator.
func NodeListEntryIteratorFromJS(value js.Value) *NodeListEntryIterator {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &NodeListEntryIterator{}
	ret.jsValue = value
	return ret
}

func (_this *NodeListEntryIterator) Next() (_result *NodeListEntryIteratorValue) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.jsValue.Call("next", _args[0:_end]...)
	var (
		_converted *NodeListEntryIteratorValue // javascript: NodeListEntryIteratorValue _what_return_name
	)
	_converted = NodeListEntryIteratorValueFromJS(_returned)
	_result = _converted
	return
}

/****************************************************************************
* NodeListKeyIterator
*****************************************************************************/

type NodeListKeyIterator struct {
	jsValue js.Value
}

// NodeListKeyIteratorFromJS is casting a js.Value into NodeListKeyIterator.
func NodeListKeyIteratorFromJS(value js.Value) *NodeListKeyIterator {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &NodeListKeyIterator{}
	ret.jsValue = value
	return ret
}

func (_this *NodeListKeyIterator) Next() (_result *NodeListKeyIteratorValue) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.jsValue.Call("next", _args[0:_end]...)
	var (
		_converted *NodeListKeyIteratorValue // javascript: NodeListKeyIteratorValue _what_return_name
	)
	_converted = NodeListKeyIteratorValueFromJS(_returned)
	_result = _converted
	return
}

/****************************************************************************
* NodeListValueIterator
*****************************************************************************/

type NodeListValueIterator struct {
	jsValue js.Value
}

// NodeListValueIteratorFromJS is casting a js.Value into NodeListValueIterator.
func NodeListValueIteratorFromJS(value js.Value) *NodeListValueIterator {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &NodeListValueIterator{}
	ret.jsValue = value
	return ret
}

func (_this *NodeListValueIterator) Next() (_result *NodeListValueIteratorValue) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.jsValue.Call("next", _args[0:_end]...)
	var (
		_converted *NodeListValueIteratorValue // javascript: NodeListValueIteratorValue _what_return_name
	)
	_converted = NodeListValueIteratorValueFromJS(_returned)
	_result = _converted
	return
}
