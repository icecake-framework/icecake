package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ick "github.com/sunraylab/icecake/pkg/icecake"
)

func TestNode(t *testing.T) {

	t.Run("IsInDOM", func(t *testing.T) {

		assert.False(t, new(ick.Element).IsInDOM()) // console warning --> unable to get "isConnected": undefined js value

		div := ick.App.CreateElement("DIV").SetId("tstisindom")
		assert.False(t, div.IsInDOM())

		ick.App.ChildById("test-container").AppendChild(&div.Node)
		assert.True(t, div.IsInDOM())
	})

	t.Run("Children", func(t *testing.T) {

		assert.False(t, new(ick.Node).HasChildren())

		dive := ick.App.ChildById("tstisindom")
		assert.False(t, dive.HasChildren())
		assert.True(t, ick.App.ChildById("test-container").HasChildren())

		// has := _node.Call("hasChildNodes")
	})

}
