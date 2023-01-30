package dom

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

// An integer value representing otherNode's position relative to node as a bitmask.
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

// NewNodeFromJS is casting a js.Value into Node.
func NewNodeFromJS(value js.Value) *Node {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Node{}
	ret.jsValue = value
	return ret
}

// IsDefined returns true if the Element is not nil AND it's type is not TypeNull and not TypeUndefined
func (_node *Node) IsDefined() bool {
	if _node == nil || _node.jsValue.Type() == js.TypeNull || _node.jsValue.Type() == js.TypeUndefined {
		return false
	}
	return true
}

// IsSameNode tests whether two nodes are the same (in other words, whether they reference the same object).
// https://developer.mozilla.org/en-US/docs/Web/API/Node/isSameNode
func (_node *Node) IsSameNode(otherNode *Node) bool {
	var _args [1]interface{}
	_args[0] = otherNode.jsValue
	_returned := _node.jsValue.Call("isSameNode", _args[0:1]...)
	return (_returned).Bool()
}

/****************************************************************************
* Node's method and properties
*****************************************************************************/

// NodeType It distinguishes different kind of nodes from each other, such as elements, text and comments.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeType
func (_node *Node) NodeType() NT_NodeType {
	value := _node.jsValue.Get("nodeType")
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
func (_node *Node) NodeName() string {
	return _node.jsValue.Get("nodeName").String()
}

// BaseURI returns the absolute base URL of the document containing the node.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/baseURI
func (_node *Node) BaseURI() string {
	return _node.jsValue.Get("baseURI").String()
}

// IsDocConnected returns a boolean indicating whether the node is connected (directly or indirectly) to a Document object.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/isConnected
func (_node *Node) IsConnected() bool {
	value := _node.jsValue.Get("isConnected")
	return (value).Bool()
}

// Doc returns the top-level document object of the node, the top-level object in which all the child nodes are created.
// on a node that is itself a document, the value is null.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/ownerDocument
func (_node *Node) Doc() *Document {
	value := _node.jsValue.Get("ownerDocument")
	return NewDocumentFromJS(value)
}

// GetRootNode returns the context object's root, which optionally includes the shadow root if it is available.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/getRootNode
func (_node *Node) RootNode() (_result *Node) {
	var _args [1]interface{}
	_returned := _node.jsValue.Call("getRootNode", _args[0:0]...)
	return NewNodeFromJS(_returned)
}

// ParentNode returns the parent of the specified node in the DOM tree.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/parentNode
func (_node *Node) ParentNode() *Node {
	var ret *Node
	value := _node.jsValue.Get("parentNode")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = NewNodeFromJS(value)
	}
	return ret
}

// ParentElement returns the DOM node's parent Element, or null if the node either has no parent, or its parent isn't a DOM Element.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/parentElement
func (_node *Node) ParentElement() *Element {
	value := _node.jsValue.Get("parentElement")
	return NewElementFromJS(value)
}

// ChildNodes returns a ~live~ static NodeList of child nodes of the given element where the first child node is assigned index 0.
// Child nodes include elements, text and comments.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/childNodes
func (_node *Node) Children() Nodes {
	value := _node.jsValue.Get("childNodes")
	return MakeNodesFromJSNodeList(value)
}

// FirstChild returns the node's first child in the tree, or null if the node has no children.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/firstChild
func (_node *Node) ChildFirst() *Node {
	value := _node.jsValue.Get("firstChild")
	return NewNodeFromJS(value)
}

// LastChild returns the last child of the node. If its parent is an element, then the child is generally an element node, a text node, or a comment node. It returns null if there are no child nodes.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/lastChild
func (_node *Node) ChildLast() *Node {
	value := _node.jsValue.Get("lastChild")
	return NewNodeFromJS(value)
}

// PreviousSibling  returns the node immediately preceding the specified one in its parent's childNodes list, or null if the specified node is the first in that list.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/previousSibling
func (_node *Node) SiblingPrevious() *Node {
	value := _node.jsValue.Get("previousSibling")
	return NewNodeFromJS(value)
}

// NextSibling returns the node immediately following the specified one in their parent's childNodes, or returns null if the specified node is the last child in the parent element.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/nextSibling
func (_node *Node) SiblingNext() *Node {
	value := _node.jsValue.Get("nextSibling")
	return NewNodeFromJS(value)
}

// NodeValue is a string containing the value of the current node, if any.
//
// For the document itself, nodeValue returns null.
// For text, comment, and CDATA nodes, nodeValue returns the content of the node.
// For attribute nodes, the value of the attribute is returned.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeValue
func (_node *Node) NodeValue() string {
	value := _node.jsValue.Get("nodeValue")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		return ""
	}
	return value.String()
}

// NodeValue is a string containing the value of the current node, if any.
//
// For the document itself, nodeValue returns null.
// For text, comment, and CDATA nodes, nodeValue returns the content of the node.
// For attribute nodes, the value of the attribute is returned.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeValue
func (_node *Node) SetNodeValue(value string) (_ret *Node) {
	var input interface{}
	if value != "" {
		input = value
	} else {
		input = nil
	}
	_node.jsValue.Set("nodeValue", input)
	return _node
}

// TextContent represents the text content of the node and its descendants.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent
func (_node *Node) TextContent() string {
	value := _node.jsValue.Get("textContent")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		return ""
	}
	return value.String()
}

// TextContent represents the text content of the node and its descendants.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent
func (_node *Node) SetTextContent(value string) (_ret *Node) {
	var input interface{}
	if value != "" {
		input = value
	} else {
		input = nil
	}
	_node.jsValue.Set("textContent", input)
	return _node
}

// HasChildNodes returns a boolean value indicating whether the given Node has child nodes or not.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/hasChildNodes
func (_node *Node) HasChildren() bool {
	_returned := _node.jsValue.Call("hasChildNodes")
	return _returned.Bool()
}

// CompareDocumentPosition reports the position of its argument node relative to the node on which it is called.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/compareDocumentPosition
func (_onenode *Node) ComparePosition(_other *Node) NodePosition {
	_returned := _onenode.jsValue.Call("compareDocumentPosition", _other.jsValue)
	return NodePosition(_returned.Int())
}

// InsertBefore inserts a newnode before a refnode.
//
// if refnode is nil, then newNode is inserted at the end of node's child nodes.
//
// If the given node already exists in the document, insertBefore() moves it from its current position to the new position.
// (That is, it will automatically be removed from its existing parent before appending it to the specified new parent.)
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/insertBefore
func (_parentnode *Node) InsertBefore(newnode *Node, refnode *Node) (_result *Node) {
	_returned := _parentnode.jsValue.Call("insertBefore", newnode.jsValue, refnode.jsValue)
	return NewNodeFromJS(_returned)
}

// AppenChild adds a node to the end of the list of children of a specified parent node.
// If the given child is a reference to an existing node in the document, appendChild() moves it from its current position to the new position.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/appendChild
func (_parentnode *Node) AppendChild(newnode *Node) (_result *Node) {
	_returned := _parentnode.jsValue.Call("appendChild", newnode.jsValue)
	return NewNodeFromJS(_returned)
}

// ReplaceChild replaces a child node within the given (parent) node.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/replaceChild
func (_parentnode *Node) ReplaceChild(_newchild *Node, _oldchild *Node) (_result *Node) {
	_returned := _parentnode.jsValue.Call("replaceChild", _newchild.jsValue, _oldchild.jsValue)
	return NewNodeFromJS(_returned)
}

// RemoveChild removes a child node from the DOM and returns the removed node.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/removeChild
func (_parentnode *Node) RemoveChild(_newchild *Node) (_result *Node) {
	_returned := _parentnode.jsValue.Call("removeChild", _newchild.jsValue)
	return NewNodeFromJS(_returned)
}
