package main

import (
	"testing"

	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T) {

	t.Run("IsInDOM", func(t *testing.T) {

		// console warning --> unable to get "isConnected": undefined js value
		assert.False(t, new(dom.Element).IsInDOM())

		div := dom.CreateElement("DIV").SetId("tstisindom")
		assert.False(t, div.IsInDOM())

		dom.Doc().ChildById("test-container").AppendChild(&div.Node)
		assert.True(t, div.IsInDOM())
	})

	t.Run("Children", func(t *testing.T) {

		// console warning --> unable to call "hasChildNodes": undefined js value
		assert.False(t, new(dom.Node).HasChildren())

		dive := dom.Doc().ChildById("tstisindom")
		assert.False(t, dive.HasChildren())
		assert.True(t, dom.Doc().ChildById("test-container").HasChildren())
	})

}
