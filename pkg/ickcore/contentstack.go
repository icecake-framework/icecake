package ickcore

import (
	"io"
	"reflect"

	"github.com/huandu/go-clone"
)

// ContentStack is a stack of ContentComposer than can easily be embedded into any custom snippet.
// Call Push to feed the stack and call RenderStack into the RenderContent fonction of the custom snippet.
type ContentStack struct {
	Stack []ContentComposer
}

func (c ContentStack) Clone() ContentStack {
	var n ContentStack
	if len(c.Stack) > 0 {
		copy := clone.Clone(c.Stack)
		n.Stack = copy.([]ContentComposer)
	}
	return n
}

// Clear clears the rendering stack
func (c *ContentStack) ClearContent() {
	if c != nil {
		c.Stack = c.Stack[:0]
	}
}

// NeedRendering returns true is the content stack is not nil and it contains at least on item
func (c *ContentStack) NeedRendering() bool {
	return c != nil && c.Stack != nil && len(c.Stack) > 0
}

// Push adds one or many composers to the rendering stack.
// Returns the snippet to allow chaining calls.
//
// Warning: Struct embedding ICKSnippet should be car of Push returns an ICKSnippet and not the parent stuct type.
func (c *ContentStack) Push(content ...ContentComposer) {
	if c.Stack == nil {
		c.Stack = make([]ContentComposer, 0)
	}
	if len(content) > 0 {
		for _, cmp := range content {
			if cmp != nil && reflect.TypeOf(cmp).Kind() == reflect.Ptr && !reflect.ValueOf(cmp).IsNil() {
				c.Stack = append(c.Stack, cmp)
			}
		}
	}
}

// RenderStack writes the HTML string corresponding to the content of the HTML element.
// The default implementation for an HTMLSnippet snippet is to render all the internal stack of composers inside an enclosed HTML tag.
func (c *ContentStack) RenderStack(out io.Writer, parent RMetaProvider) (err error) {
	if c.Stack != nil && len(c.Stack) > 0 {
		for _, child := range c.Stack {
			err := RenderChild(out, parent, child)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
