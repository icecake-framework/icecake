package ick0

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	c1 := new(HtmlComponent)
	err := TheCmpReg.RegisterComponent(c1)
	assert.ErrorContains(t, err, "not a pointer")

	err = TheCmpReg.RegisterComponent(*c1)
	assert.ErrorContains(t, err, "must be an HtmlComposer")

	err = TheCmpReg.RegisterComponent(Text{})
	assert.NoError(t, err) // only log "already registered"

	e := TheCmpReg.LookupComponent("ick-text")
	assert.NotNil(t, e)

	ct := reflect.TypeOf(Text{})
	e = TheCmpReg.LookupComponentType(ct)
	assert.NotNil(t, e)

}
