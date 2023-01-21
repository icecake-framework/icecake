package webclientsdk

import "syscall/js"

/****************************************************************************
* enums: NodeFilter
*****************************************************************************/

type NodeTypeFilter uint

const (
	SHOW_ALL                    NodeTypeFilter = 0xFFFFFFFF // Shows all nodes.
	SHOW_ELEMENT                NodeTypeFilter = 0x001
	SHOW_TEXT                   NodeTypeFilter = 0x004
	SHOW_PROCESSING_INSTRUCTION NodeTypeFilter = 0x040
	SHOW_COMMENT                NodeTypeFilter = 0x080
	SHOW_DOCUMENT               NodeTypeFilter = 0x100
	SHOW_DOCUMENT_TYPE          NodeTypeFilter = 0x200
	SHOW_DOCUMENT_FRAGMENT      NodeTypeFilter = 0x400

	SHOW_ATTRIBUTE        NodeTypeFilter = 0x002 // Deprectated
	SHOW_CDATA_SECTION    NodeTypeFilter = 0x008 // Deprectated
	SHOW_ENTITY_REFERENCE NodeTypeFilter = 0x010 // Deprectated
	SHOW_ENTITY           NodeTypeFilter = 0x020 // Deprectated
	SHOW_NOTATION         NodeTypeFilter = 0x800 // Deprectated
)

/****************************************************************************
* Nodes
*****************************************************************************/

// Nodes is a slice of node in the Document.
// Its built as a snaphot and is isefull to iterare thru.
type Nodes []*Node

// MakeNodes make a slice of nodes, scaning existing nodes from root to the last sibling node.
// Only nodes matching filter AND the optional match function are included.
func MakeNodes(root *Node, filter NodeTypeFilter, match func(*Node) bool) Nodes {
	nodes := make(Nodes, 0)
	if root == nil || root.JSValue().Type() == js.TypeNull || root.JSValue().Type() == js.TypeUndefined {
		return nodes
	}

	// deprecated filters
	if filter&(SHOW_ENTITY|SHOW_ENTITY_REFERENCE|SHOW_NOTATION|SHOW_ATTRIBUTE|SHOW_CDATA_SECTION) > 0 {
		return nodes
	}

	for scan := root; scan != nil; scan = root.NextSibling() {
		nt := scan.NodeType()
		fin := (filter & SHOW_ALL) > 0
		fin = fin || (filter&SHOW_ELEMENT) > 0 && nt == NT_ELEMENT
		fin = fin || (filter&SHOW_COMMENT) > 0 && nt == NT_COMMENT
		fin = fin || (filter&SHOW_TEXT) > 0 && nt == NT_TEXT
		fin = fin || (filter&SHOW_DOCUMENT) > 0 && nt == NT_DOCUMENT
		fin = fin || (filter&SHOW_DOCUMENT_FRAGMENT) > 0 && nt == NT_DOCUMENT_FRAGMENT
		fin = fin || (filter&SHOW_DOCUMENT_TYPE) > 0 && nt == NT_DOCUMENT_TYPE
		fin = fin || (filter&SHOW_PROCESSING_INSTRUCTION) > 0 && nt == NT_PROCESSING_INSTRUCTION
		if fin {
			if match != nil {
				fin = fin || match(scan)
			}
			nodes = append(nodes, scan)
		}
	}

	return nodes
}
