package html

import (
	"io"
	"reflect"

	"github.com/icecake-framework/icecake/pkg/ickcore"
	"github.com/lolorenzo777/verbose"
)

// maxDEEP is the maximum HTML string unfolding levels
const maxDEEP int = 25

type ContentComposer interface {

	// Meta returns a reference to render meta data
	ickcore.RMetaProvider

	// RenderContent writes the HTML string corresponding to the content of the HTML element.
	// Return an error to stops the rendering process.
	RenderContent(out io.Writer) error
}

// ElementComposer interface
type ElementComposer interface {
	TagBuilder
	ContentComposer
}

// RenderChild renders the HTML string of the composers to out, including its tag element its properties and its content.
// Rendering the content can renders child-snippets recursively. This can be done maxDEEP times max to avoid infinite loop.
//
// If composer is a also a TagBuilder the output looks like this:
//
//	`<{tagname} id={xxx} name="{ick-tag}" [attributes]>[content]</tagname>`
//
// otherwise only the content is written.
//
// A snippet id can be setup up upfront (a) accessing any saved tag attribute within the snippet struct, or (b) within an html ick-tag attribute (for embedded snippet).
// Thes id will be lost if ther'e a parent, the snippet attributes will be overwritten with the unique id generated by the rendering process.
// Unique ids are generated by using the composer name (without "ick-" prefix) with a sequence number. sub-composer ids combine the id of the parent with the id of the sub-composer.
// if the component is not registered and so does't have a name, a global unique id is generated.
// This behaviour ensures that ids are uniques even for multiple instanciations of the same composer.
//
// snippet may have none id on request. noid snippet attribute must be set to true to render the composer without id.
// The special attribute noid can be defined within an ick-tag html or with attribute's methods.
//
// If the parent is not nil, the snippet is added to its embedded stack of sub-components.
//
// Returns rendering errors, typically with the writer, or if there's too many recursive rendering.
func RenderChild(out io.Writer, parent ickcore.RMetaProvider, child ContentComposer, siblings ...ContentComposer) error {
	err := render(out, parent, child)
	if err != nil {
		return err
	}
	for _, s := range siblings {
		err := render(out, parent, s)
		if err != nil {
			return err
		}
	}
	return nil
}

func render(out io.Writer, parent ickcore.RMetaProvider, cmp ContentComposer) error {

	// nothing to render
	if cmp == nil || reflect.TypeOf(cmp).Kind() != reflect.Ptr || reflect.ValueOf(cmp).IsNil() {
		verbose.Printf(verbose.WARNING, "Render: empty composer %s\n", reflect.TypeOf(cmp).String())
		return nil
	}

	// look for depth and ensure no infinite loop
	deep := 0
	if parent != nil {
		if deep = parent.RMeta().Deep + 1; deep > maxDEEP {
			return verbose.Error("Render", ErrTooManyRecursiveRendering)
		}
		cmp.RMeta().Deep = parent.RMeta().Deep + 1
	}
	verbose.Printf(verbose.INFO, "rendering L.%v composer %s\n", deep, reflect.TypeOf(cmp).String())

	// build the tag
	var tag Tag
	cmptag, istagger := cmp.(TagBuilder)
	if istagger && cmptag != nil {
		tag = BuildTag(cmptag)
	}

	// generate the virtual id
	cmp.RMeta().GenerateVirtualId(cmp)

	// verbose id information
	//verbose.Debug(" vid:%s --> id:%s", virtualid, cmpid)

	// render openingtag
	if cmptag != nil {
		selfclosed, err := tag.RenderOpening(out)
		if selfclosed || err != nil {
			cmp.RMeta().RError = err
			return err
		}
	}

	// Render the content
	err := cmp.RenderContent(out)
	if err != nil {
		cmp.RMeta().RError = err
		return err
	}

	// Render closingtag
	if cmptag != nil {
		err := tag.RenderClosing(out)
		if err != nil {
			cmp.RMeta().RError = err
			return err
		}
	}

	// add it to the map of embedded components
	cmp.RMeta().IsRender = true
	cmp.RMeta().Deep = 0
	if parent != nil {
		parent.RMeta().Embed(cmp)
	}

	return nil
}
