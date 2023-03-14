package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	wick "github.com/sunraylab/icecake/pkg/wicecake"
)

func TestNode(t *testing.T) {

	t.Run("IsInDOM", func(t *testing.T) {

		assert.False(t, new(wick.Element).IsInDOM()) // console warning --> unable to get "isConnected": undefined js value

		div := wick.App.CreateElement("DIV").SetId("tstisindom")
		assert.False(t, div.IsInDOM())

		wick.App.ChildById("test-container").AppendChild(&div.Node)
		assert.True(t, div.IsInDOM())
	})

	t.Run("Children", func(t *testing.T) {

		assert.False(t, new(wick.Node).HasChildren())

		dive := wick.App.ChildById("tstisindom")
		assert.False(t, dive.HasChildren())
		assert.True(t, wick.App.ChildById("test-container").HasChildren())

		// has := _node.Call("hasChildNodes")
	})

}
