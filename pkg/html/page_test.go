package html

import (
	"bytes"
	"testing"
)

func TestDefaultPage(t *testing.T) {

	dft := NewPage("en")
	out := new(bytes.Buffer)
	dft.Render(out)

	if out.String() != `<!doctype html><html lang="en"><head><META charset="utf-8"></head><body></body></html>` {
		t.Fail()
	}

}
