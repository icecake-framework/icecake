package wick

import (
	"fmt"

	"github.com/sunraylab/icecake/pkg/console"
)

const (
	UNDEFINED_NODE string = "<undefined node>"
)

/****************************************************************************
* enums NodeType, NodePosition
*****************************************************************************/

type NODE_TYPE int

const (
	NT_ALL                    NODE_TYPE = 0xFFFF
	NT_UNDEF                  NODE_TYPE = 0x0000
	NT_ELEMENT                NODE_TYPE = 0x0001 // An Element node like <p> or <div>. aka ELEMNT_NODE.
	NT_ATTRIBUTE              NODE_TYPE = 0x0002 // An Attribute of an Element. aka ATTRIBUTE_NODE.
	NT_TEXT                   NODE_TYPE = 0x0003 // The actual Text inside an Element or Attr. aka TEXT_NODE.
	NT_CDATA_SECTION          NODE_TYPE = 0x0004 // A CDATASection, such as <!CDATA[[ … ]]>. aka CDATA_SECTION_NODE.
	NT_PROCESSING_INSTRUCTION NODE_TYPE = 0x0007 // A ProcessingInstruction of an XML document, such as <?xml-stylesheet … ?>.
	NT_COMMENT                NODE_TYPE = 0x0008 // A Comment node, such as <!-- … -->. aka COMMENT_NODE.
	NT_DOCUMENT               NODE_TYPE = 0x0009 // A Document node. aka DOCUMENT_NODE.
	NT_DOCUMENT_TYPE          NODE_TYPE = 0x000A // A DocumentType node, such as <!DOCTYPE html>. aka DOCUMENT_TYPE_NODE.
	NT_DOCUMENT_FRAGMENT      NODE_TYPE = 0x000B // A DocumentFragment node. aka DOCUMENT_FRAGMENT_NODE.

	NT_ENTITY_REFERENCE NODE_TYPE = 0x0005 // Deprecated
	NT_ENTITY           NODE_TYPE = 0x0006 // Deprecated
	NT_NOTATION         NODE_TYPE = 0x000C // Deprecated
)

func (nt NODE_TYPE) String() string {
	switch nt {
	case NT_ELEMENT:
		return "Element"
	case NT_ATTRIBUTE:
		return "Attribute"
	case NT_TEXT:
		return "Text"
	case NT_CDATA_SECTION:
		return "Data Section"
	case NT_PROCESSING_INSTRUCTION:
		return "Processing Instruction"
	case NT_COMMENT:
		return "Comment"
	case NT_DOCUMENT:
		return "Document"
	case NT_DOCUMENT_TYPE:
		return "Document Type"
	case NT_DOCUMENT_FRAGMENT:
		return "Document Fragment"
	}
	return fmt.Sprintf("unmanaged node type: %d", nt)
}

/****************************************************************************
* enums: NodeFilter
*****************************************************************************/

// An integer value representing otherNode's position relative to node as a bitmask.
type NODE_POSITION int

const (
	NODEPOS_UNDEF                   NODE_POSITION = 0x00
	NODEPOS_DISCONNECTED            NODE_POSITION = 0x01
	NODEPOS_PRECEDING               NODE_POSITION = 0x02
	NODEPOS_FOLLOWING               NODE_POSITION = 0x04
	NODEPOS_CONTAINS                NODE_POSITION = 0x08
	NODEPOS_CONTAINED_BY            NODE_POSITION = 0x10
	NODEPOS_IMPLEMENTATION_SPECIFIC NODE_POSITION = 0x20
)

/****************************************************************************
* Node
*****************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/Node
type Node struct {
	EventTarget
}

// CastNode is casting a js.Value into Node.
func CastNode(_jsvp JSValueProvider) *Node {
	if _jsvp.Value().Type() != TYPE_OBJECT {
		console.Errorf("casting Node failed\n")
		return new(Node)
	}
	cast := new(Node)
	cast.jsvalue = _jsvp.Value().jsvalue
	return cast
}

func CastNodes(_jsvp JSValueProvider) []*Node {
	nodes := make([]*Node, 0)
	if _jsvp.Value().Type() != TYPE_OBJECT {
		console.Errorf("casting Nodes failed\n")
		return nodes
	}
	len := _jsvp.Value().GetInt("length")
	for i := 0; i < len; i++ {
		_returned := _jsvp.Value().Call("item", uint(i))
		node := CastNode(_returned)
		nodes = append(nodes, node)
	}
	return nodes
}

// IsDefined returns true if the Element is not nil AND it's type is not TypeNull and not TypeUndefined
func (_node *Node) IsDefined() bool {
	if _node == nil {
		return false
	}
	return _node.JSValue.IsDefined()
}

// IsSameNode tests whether two nodes are the same (in other words, whether they reference the same object).
// https://developer.mozilla.org/en-US/docs/Web/API/Node/isSameNode
func (_node *Node) IsSameNode(_otherNode *Node) bool {
	return _node.Call("isSameNode", _otherNode.jsvalue).Truthy()
}

/****************************************************************************
* Node's method and properties
*****************************************************************************/

// NodeType It distinguishes different kind of nodes from each other, such as elements, text and comments.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeType
func (_node *Node) NodeType() NODE_TYPE {
	if !_node.IsDefined() {
		return NT_UNDEF
	}
	nt := _node.GetInt("nodeType")
	return NODE_TYPE(nt)
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
// NodeName returns the name of the current node as a string.
func (_node *Node) NodeName() string {
	if !_node.IsDefined() {
		return UNDEFINED_NODE
	}
	return _node.GetString("nodeName")
}

// BaseURI returns the absolute base URL of the document containing the node.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/baseURI
func (_node *Node) BaseURI() string {
	if !_node.IsDefined() {
		return UNDEFINED_NODE
	}
	return _node.GetString("baseURI")
}

// IsInDOM returns a boolean indicating whether the node is within a Document object.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/isConnected
// func (_node Node) IsConnected() bool {
func (_node Node) IsInDOM() bool {
	return _node.GetBool("isConnected")
}

// GetRootNode returns the context object's root, which optionally includes the shadow root if it is available.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/getRootNode
func (_node Node) RootNode() *Node {
	return CastNode(_node.Call("getRootNode"))
}

// ParentNode returns the parent of the specified node in the DOM tree.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/parentNode
func (_node Node) ParentNode() *Node {
	return CastNode(_node.Get("parentNode"))
}

// ParentElement returns the DOM node's parent Element, or null if the node either has no parent, or its parent isn't a DOM Element.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/parentElement
func (_node Node) ParentElement() *Element {
	return CastElement(_node.Get("parentElement"))
}

// HasChildNodes returns a boolean value indicating whether the given Node has child nodes or not.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/hasChildNodes
func (_node *Node) HasChildren() bool {
	return _node.Call("hasChildNodes").Truthy()
}

// ChildNodes returns a ~live~ static NodeList of child nodes of the given element where the first child node is assigned index 0.
// Child nodes include elements, text and comments.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/childNodes
func (_node *Node) Children() []*Node {
	return CastNodes(_node.Get("childNodes"))
}

// FilteredChildren make a slice of nodes, scaning recursively children of every existing nodes.
// Only nodes matching _filter AND the optional _match function are included.
func (_root *Node) FilteredChildren(_filter NODE_TYPE, _match func(*Node) bool) []*Node {
	if !_root.IsDefined() || !_root.HasChildren() {
		return make([]*Node, 0)
	}
	return _root.filteredChildren(_filter, 999, _match)
}

func (_root *Node) filteredChildren(_filter NODE_TYPE, _deepmax int, _match func(*Node) bool) []*Node {
	nodes := make([]*Node, 0)

	for _, scan := range _root.Children() {
		// check filtered node type
		fn := _filter == NT_ALL || scan.NodeType() == _filter

		// apply the filter to children if not too deep and the type node is selected
		if fn && scan.HasChildren() {
			if _deepmax > 0 {
				sub := scan.filteredChildren(_filter, _deepmax-1, _match)
				nodes = append(nodes, sub...)
			} else {
				console.Warnf("FilteredChildren reached max level")
			}
		}

		// apply matching function
		if fn && _match != nil {
			fn = fn && _match(scan)
		}

		if fn {
			nodes = append(nodes, scan)
		}
	}
	return nodes
}

// FirstChild returns the node's first child in the tree, or null if the node has no children.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/firstChild
func (_node *Node) ChildFirst() *Node {
	return CastNode(_node.Get("firstChild"))
}

// LastChild returns the last child of the node.
// If its parent is an element, then the child is generally an element node, a text node, or a comment node.
// It returns null if there are no child nodes.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/lastChild
func (_node *Node) ChildLast() *Node {
	return CastNode(_node.Get("lastChild"))
}

// PreviousSibling  returns the node immediately preceding the specified one in its parent's childNodes list, or null if the specified node is the first in that list.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/previousSibling
func (_node *Node) SiblingPrevious() *Node {
	return CastNode(_node.Get("previousSibling"))
}

// NextSibling returns the node immediately following the specified one in their parent's childNodes, or returns null if the specified node is the last child in the parent element.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/nextSibling
func (_node *Node) SiblingNext() *Node {
	return CastNode(_node.Get("nextSibling"))
}

// NodeValue is a string containing the value of the current node, if any.
//
// For the document itself, nodeValue returns null.
// For text, comment, and CDATA nodes, nodeValue returns the content of the node.
// For attribute nodes, the value of the attribute is returned.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeValue
func (_node *Node) NodeValue() string {
	if !_node.IsDefined() {
		return UNDEFINED_NODE
	}
	return _node.GetString("nodeValue")
}

// NodeValue is a string containing the value of the current node, if any.
//
// For the document itself, nodeValue returns null.
// For text, comment, and CDATA nodes, nodeValue returns the content of the node.
// For attribute nodes, the value of the attribute is returned.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeValue
func (_node *Node) SetNodeValue(value string) *Node {
	if value != "" {
		_node.Set("nodeValue", value)
	} else {
		_node.Set("nodeValue", nil)
	}
	return _node
}

// TextContent represents the text content of the node and its descendants.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent
func (_node *Node) TextContent() string {
	if !_node.IsDefined() {
		return UNDEFINED_NODE
	}
	return _node.GetString("textContent")
}

// TextContent represents the text content of the node and its descendants.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent
func (_node *Node) SetTextContent(value string) *Node {
	if value != "" {
		_node.Set("textContent", value)
	} else {
		_node.Set("textContent", nil)
	}
	return _node
}

// CompareDocumentPosition reports the position of its argument node relative to the node on which it is called.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/compareDocumentPosition
func (_onenode *Node) ComparePosition(_other *Node) NODE_POSITION {
	if !_onenode.IsDefined() {
		return NODEPOS_UNDEF
	}
	_returned := _onenode.Call("compareDocumentPosition", _other.jsvalue)
	return NODE_POSITION(_returned.Int())
}

// InsertBefore inserts a newnode before a refnode. if refnode is nil, then newNode is inserted at the end of node's child nodes.
//
// If the given node already exists in the document, insertBefore() moves it from its current position to the new position.
// (That is, it will automatically be removed from its existing parent before appending it to the specified new parent.)
//
// Returns the _parent to enable chaining calls.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/insertBefore
func (_parentnode *Node) InsertBefore(newnode *Node, refnode *Node) *Node {
	_parentnode.Call("insertBefore", newnode.jsvalue, refnode.jsvalue)
	return _parentnode
}

// AppenChild adds a node to the end of the list of children of a specified parent node.
// If the given child is a reference to an existing node in the document, appendChild() moves it from its current position to the new position.
//
// Returns the _parent to enable chaining calls.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/appendChild
func (_parentnode *Node) AppendChild(newnode *Node) *Node {
	_parentnode.Call("appendChild", newnode.jsvalue)
	return _parentnode
}

// ReplaceChild replaces a child node within the given (parent) node.
//
// Returns the _parent to enable chaining calls.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/replaceChild
func (_parentnode *Node) ReplaceChild(_newchild *Node, _oldchild *Node) *Node {
	_parentnode.Call("replaceChild", _newchild.jsvalue, _oldchild.jsvalue)
	return _parentnode
}

// RemoveChild removes a child node from the DOM
//
// Returns the _parent to enable chaining calls.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/removeChild
func (_parentnode *Node) RemoveChild(_child *Node) *Node {
	_parentnode.Call("removeChild", _child.jsvalue)
	return _parentnode
}
