package html

import (
	"bytes"
	"testing"
)

func TestDefaultPage(t *testing.T) {

	dft := NewHtml5Page("en")
	out := new(bytes.Buffer)
	dft.Generate(out)

	if out.String() != `<!doctype html><html lang="en"><head><META charset="utf-8"></head><body></body></html>` {
		t.Fail()
	}

}
