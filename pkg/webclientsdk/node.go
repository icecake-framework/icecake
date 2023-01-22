package browser

import (
	"syscall/js"
)

/****************************************************************************
* enums NodeType, NodePosition
*****************************************************************************/

type NT_NodeType int

const (
	NT_ELEMENT                NT_NodeType = 1  // An Element node like <p> or <div>. aka ELEMNT_NODE.
	NT_ATTRIBUTE              NT_NodeType = 2  // An Attribute of an Element. aka ATTRIBUTE_NODE.
	NT_TEXT                   NT_NodeType = 3  // The actual Text inside an Element or Attr. aka TEXT_NODE.
	NT_CDATA_SECTION          NT_NodeType = 4  // A CDATASection, such as <!CDATA[[ … ]]>. aka CDATA_SECTION_NODE.
	NT_PROCESSING_INSTRUCTION NT_NodeType = 7  // A ProcessingInstruction of an XML document, such as <?xml-stylesheet … ?>.
	NT_COMMENT                NT_NodeType = 8  // A Comment node, such as <!-- … -->. aka COMMENT_NODE.
	NT_DOCUMENT               NT_NodeType = 9  // A Document node. aka DOCUMENT_NODE.
	NT_DOCUMENT_TYPE          NT_NodeType = 10 // A DocumentType node, such as <!DOCTYPE html>. aka DOCUMENT_TYPE_NODE.
	NT_DOCUMENT_FRAGMENT      NT_NodeType = 11 // A DocumentFragment node. aka DOCUMENT_FRAGMENT_NODE.

	ENTITY_REFERENCE_NODE NT_NodeType = 5  // Deprecated
	ENTITY_NODE           NT_NodeType = 6  // Deprecated
	NOTATION_NODE         NT_NodeType = 12 // Deprecated
)

type NodePosition int

const (
	DOCUMENT_POSITION_DISCONNECTED            NodePosition = 0x01
	DOCUMENT_POSITION_PRECEDING               NodePosition = 0x02
	DOCUMENT_POSITION_FOLLOWING               NodePosition = 0x04
	DOCUMENT_POSITION_CONTAINS                NodePosition = 0x08
	DOCUMENT_POSITION_CONTAINED_BY            NodePosition = 0x10
	DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC NodePosition = 0x20
)

/****************************************************************************
* Node
*****************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/Node
type Node struct {
	EventTarget
}

// MakeNodeFromJS is casting a js.Value into Node.
func MakeNodeFromJS(value js.Value) *Node {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Node{}
	ret.jsValue = value
	return ret
}

// NodeType It distinguishes different kind of nodes from each other, such as elements, text and comments.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeType
func (_this *Node) NodeType() NT_NodeType {
	value := _this.jsValue.Get("nodeType")
	return NT_NodeType((value).Int())
}

// Values for the different types of nodes are:
//   - Attr: the qualified name of the attribute.
//   - CDATASection: "#cdata-section".
//   - Comment: "#comment".
//   - Document: "#document".
//   - DocumentFragment: "#document-fragment".
//   - DocumentType: the value of DocumentType.name
//   - Element: the uppercase name of the element tag if an HTML element, or the lowercase element tag if an XML element (like a SVG or MATHML element).
//   - ProcessingInstruction: The value of ProcessingInstruction.target
//   - Text: "#text".
//
// NodeName returning attribute 'nodeName' with
func (_this *Node) NodeName() string {
	value := _this.jsValue.Get("nodeName")
	return (value).String()
}

// BaseURI returns the absolute base URL of the document containing the node.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/baseURI
func (_this *Node) BaseURI() string {
	value := _this.jsValue.Get("baseURI")
	return (value).String()
}

// IsDocConnected returns a boolean indicating whether the node is connected (directly or indirectly) to a Document object.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/isConnected
func (_this *Node) IsDocConnected() bool {
	value := _this.jsValue.Get("isConnected")
	return (value).Bool()
}

// Doc returns the top-level document object of the node, the top-level object in which all the child nodes are created.
// on a node that is itself a document, the value is null.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/ownerDocument
func (_this *Node) Doc() *Document {
	value := _this.jsValue.Get("ownerDocument")
	return MakeDocumentFromJS(value)
}

// ParentNode returns the parent of the specified node in the DOM tree.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/parentNode
func (_this *Node) ParentNode() *Node {
	var ret *Node
	value := _this.jsValue.Get("parentNode")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = MakeNodeFromJS(value)
	}
	return ret
}

// ParentElement returns the DOM node's parent Element, or null if the node either has no parent, or its parent isn't a DOM Element.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/parentElement
func (_this *Node) ParentElement() *Element {
	value := _this.jsValue.Get("parentElement")
	return MakeElementFromJS(value)
}

// ChildNodes returns a live NodeList of child nodes of the given element where the first child node is assigned index 0. Child nodes include elements, text and comments.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/childNodes
func (_this *Node) ChildNodes() *NodeList {
	value := _this.jsValue.Get("childNodes")
	return MakeNodeListFromJS(value)
}

// FirstChild returns the node's first child in the tree, or null if the node has no children.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/firstChild
func (_this *Node) FirstChild() *Node {
	value := _this.jsValue.Get("firstChild")
	return MakeNodeFromJS(value)
}

// LastChild returns the last child of the node. If its parent is an element, then the child is generally an element node, a text node, or a comment node. It returns null if there are no child nodes.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/lastChild
func (_this *Node) LastChild() *Node {
	value := _this.jsValue.Get("lastChild")
	return MakeNodeFromJS(value)
}

// PreviousSibling  returns the node immediately preceding the specified one in its parent's childNodes list, or null if the specified node is the first in that list.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/previousSibling
func (_this *Node) PreviousSibling() *Node {
	var ret *Node
	value := _this.jsValue.Get("previousSibling")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = MakeNodeFromJS(value)
	}
	return ret
}

// NextSibling returns the node immediately following the specified one in their parent's childNodes, or returns null if the specified node is the last child in the parent element.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/nextSibling
func (_this *Node) NextSibling() *Node {
	var ret *Node
	value := _this.jsValue.Get("nextSibling")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = MakeNodeFromJS(value)
	}
	return ret
}

// NodeValue is a string containing the value of the current node, if any.
// For the document itself, nodeValue returns null.
// For text, comment, and CDATA nodes, nodeValue returns the content of the node.
// For attribute nodes, the value of the attribute is returned.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeValue
func (_this *Node) NodeValue() string {
	value := _this.jsValue.Get("nodeValue")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		return ""
	}
	return (value).String()
}

// NodeValue is a string containing the value of the current node, if any.
// For the document itself, nodeValue returns null.
// For text, comment, and CDATA nodes, nodeValue returns the content of the node.
// For attribute nodes, the value of the attribute is returned.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeValue
func (_this *Node) SetNodeValue(value string) {
	var input interface{}
	if value != "" {
		input = value
	} else {
		input = nil
	}
	_this.jsValue.Set("nodeValue", input)
}

// TextContent represents the text content of the node and its descendants.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent
func (_this *Node) TextContent() string {
	value := _this.jsValue.Get("textContent")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		return ""
	}
	return (value).String()
}

// TextContent represents the text content of the node and its descendants.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent
func (_this *Node) SetTextContent(value string) {
	var input interface{}
	if value != "" {
		input = value
	} else {
		input = nil
	}
	_this.jsValue.Set("textContent", input)
}

// GetRootNode returns the context object's root, which optionally includes the shadow root if it is available.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/getRootNode
func (_this *Node) GetRootNode() (_result *Node) {
	var _args [1]interface{}
	_returned := _this.jsValue.Call("getRootNode", _args[0:0]...)
	return MakeNodeFromJS(_returned)
}

// HasChildNodes returns a boolean value indicating whether the given Node has child nodes or not.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/hasChildNodes
func (_this *Node) HasChildNodes() bool {
	var _args [0]interface{}
	_returned := _this.jsValue.Call("hasChildNodes", _args[0:0]...)
	return (_returned).Bool()
}

// IsSameNode tests whether two nodes are the same (in other words, whether they reference the same object).
// https://developer.mozilla.org/en-US/docs/Web/API/Node/isSameNode
func (_this *Node) IsSameNode(otherNode *Node) bool {
	var _args [1]interface{}
	_args[0] = otherNode.jsValue
	_returned := _this.jsValue.Call("isSameNode", _args[0:1]...)
	return (_returned).Bool()
}

// CompareDocumentPosition reports the position of its argument node relative to the node on which it is called.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/compareDocumentPosition
func (_this *Node) CompareDocumentPosition(other *Node) NodePosition {
	var _args [1]interface{}
	_args[0] = other.jsValue
	_returned := _this.jsValue.Call("compareDocumentPosition", _args[0:1]...)
	return NodePosition((_returned).Int())
}

// InsertBefore inserts a node before a reference node as a child of a specified parent node.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/insertBefore
func (_this *Node) InsertBefore(node *Node, child *Node) (_result *Node) {
	var _args [2]interface{}
	_args[0] = node.jsValue
	_args[1] = child.jsValue
	_returned := _this.jsValue.Call("insertBefore", _args[0:2]...)
	return MakeNodeFromJS(_returned)
}

// AppenChild adds a node to the end of the list of children of a specified parent node.
// If the given child is a reference to an existing node in the document, appendChild() moves it from its current position to the new position.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/appendChild
func (_this *Node) AppendChild(node *Node) (_result *Node) {
	var _args [1]interface{}
	_args[0] = node.jsValue
	_returned := _this.jsValue.Call("appendChild", _args[0:1]...)
	return MakeNodeFromJS(_returned)
}

// ReplaceChild replaces a child node within the given (parent) node.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/replaceChild
func (_this *Node) ReplaceChild(node *Node, child *Node) (_result *Node) {
	var _args [2]interface{}
	_args[0] = node.jsValue
	_args[1] = child.jsValue
	_returned := _this.jsValue.Call("replaceChild", _args[0:2]...)
	return MakeNodeFromJS(_returned)
}

// RemoveChild removes a child node from the DOM and returns the removed node.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/removeChild
func (_this *Node) RemoveChild(child *Node) (_result *Node) {
	var _args [1]interface{}
	_args[0] = child.jsValue
	_returned := _this.jsValue.Call("removeChild", _args[0:1]...)
	return MakeNodeFromJS(_returned)
}

/****************************************************************************
* NodeList
*****************************************************************************/

// NodeList is a *live* list of notdes, returned by properties such as Node.childNodes.
//
// It can be converted into a static list calling MakeNodes(_this.Item(0))
//
// https://developer.mozilla.org/en-US/docs/Web/API/NodeList
type NodeList struct {
	jsValue js.Value
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *NodeList) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.jsValue
}

// MakeNodeListFromJS is casting a js.Value into NodeList.
func MakeNodeListFromJS(value js.Value) *NodeList {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &NodeList{}
	ret.jsValue = value
	return ret
}

// Length returns the number of items in a NodeList.
//
// https://developer.mozilla.org/en-US/docs/Web/API/NodeList/length
func (_this *NodeList) Length() uint {
	var ret uint
	value := _this.jsValue.Get("length")
	ret = (uint)((value).Int())
	return ret
}

// Item Returns a node from a NodeList by index.
//
// This method doesn't throw exceptions as long as you provide arguments.
// A value of null is returned if the index is out of range, and a TypeError is thrown if no argument is provided.
//
// https://developer.mozilla.org/en-US/docs/Web/API/NodeList/item
func (_this *NodeList) Item(index uint) (_result *Node) {
	var _args [1]interface{}
	_args[0] = index
	_returned := _this.jsValue.Call("item", _args[0:1]...)
	return MakeNodeFromJS(_returned)
}
