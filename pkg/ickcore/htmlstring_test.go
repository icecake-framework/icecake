package ickcore

import (
	"bytes"
	"fmt"
	"io"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type Unmanaged struct{}

type sniph1 struct {
	BareSnippet
	IsOk bool
	Text string
	HTML HTMLString
	I    int
	F    float64
	D    time.Duration
	U    *url.URL
	Unmanaged
}

func (s *sniph1) BuildTag() Tag {
	s.Tag().SetTagName("div")
	s.Tag().AddClassIf(s.IsOk, "ok")
	return *s.Tag()
}

func (s *sniph1) RenderContent(out io.Writer) error {
	RenderStringIf(s.Text != "", out, s.Text)
	RenderChild(out, s, &s.HTML)
	RenderStringIf(s.I != 0, out, fmt.Sprintf("%v", s.I))
	RenderStringIf(s.F != 0, out, fmt.Sprintf("%v", s.F))
	RenderStringIf(s.D != 0, out, fmt.Sprintf("%v", s.D+(time.Hour*1)))
	RenderStringIf(s.U != nil, out, fmt.Sprintf("<a href='%v'></a>", s.U))
	return nil
}

func TestRenderHTML_1(t *testing.T) {

	ResetRegistry()
	out := new(bytes.Buffer)

	out.Reset()
	AddRegistryEntry("ick-testsnippet0", &BareSnippet{})
	err := renderHTML(out, nil, *ToHTML("<ick-testsnippet0/>"))
	require.NoError(t, err)
	require.Equal(t, ``, out.String())

	out.Reset()
	AddRegistryEntry("ick-tstsniph1", &sniph1{})
	err = renderHTML(out, nil, *ToHTML("<ick-tstsniph1/>"))
	require.NoError(t, err)
	require.Equal(t, `<div name="sniph1"></div>`, out.String())

	out.Reset()
	err = renderHTML(out, nil, *ToHTML("<ick-tstsniph1 test/>"))
	require.NoError(t, err)
	require.Equal(t, `<div name="sniph1" test></div>`, out.String())

	out.Reset()
	err = renderHTML(out, nil, *ToHTML("<ick-tstsniph1 Unmanaged='test'/>"))
	require.NoError(t, err)
	require.Equal(t, `<!--ick-tstsniph1: "Unmanaged" attribute: unmanaged type ickcore.Unmanaged-->`, out.String())

	out.Reset()
	err = renderHTML(out, nil, *ToHTML("<ick-tstsniph1 IsOk=true/>"))
	require.NoError(t, err)
	require.Equal(t, `<div name="sniph1" class="ok"></div>`, out.String())

	out.Reset()
	err = renderHTML(out, nil, *ToHTML("<ick-tstsniph1 Text=success/>"))
	require.NoError(t, err)
	require.Equal(t, `<div name="sniph1">success</div>`, out.String())

	out.Reset()
	err = renderHTML(out, nil, *ToHTML(`<ick-tstsniph1 HTML="<strong>STRONG</strong>"/>`))
	require.NoError(t, err)
	require.Equal(t, `<div name="sniph1"><strong>STRONG</strong></div>`, out.String())

	out.Reset()
	err = renderHTML(out, nil, *ToHTML(`<ick-tstsniph1 I=777/>`))
	require.NoError(t, err)
	require.Equal(t, `<div name="sniph1">777</div>`, out.String())

	out.Reset()
	err = renderHTML(out, nil, *ToHTML(`<ick-tstsniph1 F=777.777/>`))
	require.NoError(t, err)
	require.Equal(t, `<div name="sniph1">777.777</div>`, out.String())

	out.Reset()
	err = renderHTML(out, nil, *ToHTML(`<ick-tstsniph1 D=5h30m40s/>`))
	require.NoError(t, err)
	require.Equal(t, `<div name="sniph1">6h30m40s</div>`, out.String())

	out.Reset()
	err = renderHTML(out, nil, *ToHTML(`<ick-tstsniph1 U="/icecake.dev"/>`))
	require.NoError(t, err)
	require.Equal(t, `<div name="sniph1"><a href='/icecake.dev'></a></div>`, out.String())
}

type sniph2 struct {
	BareSnippet
	Test int
}

func (s *sniph2) BuildTag() Tag {
	switch s.Test {
	case 1:
		s.Tag().SetTagName("span")
	case 2:
		s.Tag().SetTagName("div")
		s.Tag().ParseAttributes(`class="ts2a ts2b" tabindex=2 style="display=test;" a2`)
	}
	return *s.Tag()
}

func TestRenderHTML_2(t *testing.T) {
	tstset := []struct {
		name string
		in   string
		want string
		err  bool
	}{
		{name: "none",
			in:   ``,
			want: ``,
			err:  false},
		{name: "text",
			in:   `Hello`,
			want: `Hello`,
			err:  false},
		{name: "notrim",
			in:   `  Hell o  `,
			want: `  Hell o  `,
			err:  false},
		{name: "html",
			in:   `<strong>Hello</strong>`,
			want: `<strong>Hello</strong>`,
			err:  false},

		{name: "icktag1",
			in:   `<ick-tstsniph2 Test=1/>`,
			want: `<span name="sniph2"></span>`,
			err:  false},
		{name: "icktag2",
			in:   `<ick-tstsniph2 Test=1 />`,
			want: `<span name="sniph2"></span>`,
			err:  false},
		{name: "icktag3",
			in:   `<ick-tstsniph2 Test=1 >`,
			want: ``,
			err:  true},
		{name: "icktag4",
			in:   `<ick-/>`,
			want: ``,
			err:  true},
		{name: "attrib1",
			in:   `<ick-tstsniph2 Test=1 a/>`,
			want: `<span name="sniph2" a></span>`,
			err:  false},
		{name: "attrib2",
			in:   `<ick-tstsniph2 Test=1 a />`,
			want: `<span name="sniph2" a></span>`,
			err:  false},
		{name: "attrib3",
			in:   `<ick-tstsniph2 Test=1  c a1  b />`,
			want: `<span name="sniph2" a1 b c></span>`,
			err:  false},
		{name: "attrib4",
			in:   `<ick-tstsniph2 Test=1 ; />`,
			want: ``,
			err:  true},
		{name: "attrib5",
			in:   `<ick-tstsniph2 Test=1 a😀b />`,
			want: `<span name="sniph2" a😀b></span>`,
			err:  false},
		{name: "attrib6",
			in:   `<ick-tstsniph2 Test=1 1a />`,
			want: ``,
			err:  true},
		{name: "value without attribute 1",
			in:   `<ick-tstsniph2=/>`,
			want: ``,
			err:  true},
		{name: "value without attribute 2",
			in:   `<ick-tstsniph2 Test=1 = />`,
			want: ``,
			err:  true},
		{name: "missing value 1",
			in:   `<ick-tstsniph2 Test=1 a=/>`,
			want: ``,
			err:  true},
		{name: "missing value 2",
			in:   `<ick-tstsniph2 Test=1 a= />`,
			want: ``,
			err:  true},
		{name: "value 1",
			in:   `<ick-tstsniph2 Test=1 a=1/>`,
			want: `<span name="sniph2" a=1></span>`,
			err:  false},
		{name: "value 2",
			in:   `<ick-tstsniph2 Test=1 abc=1 />`,
			want: `<span name="sniph2" abc=1></span>`,
			err:  false},
		{name: "quoted value 1",
			in:   `<ick-tstsniph2 Test=1 a='x'/>`,
			want: `<span name="sniph2" a="x"></span>`,
			err:  false},
		{name: "quoted value 2",
			in:   `<ick-tstsniph2 Test=1 a= 'x' />`,
			want: `<span name="sniph2" a="x"></span>`,
			err:  false},
		{name: "quoted value 3",
			in:   `<ick-tstsniph2 Test=1 a='x' />`,
			want: `<span name="sniph2" a="x"></span>`,
			err:  false},

		{name: "value with spaces",
			in:   `<ick-tstsniph2 Test=1 a=' x '/>`,
			want: `<span name="sniph2" a=" x "></span>`,
			err:  false},
		{name: "double quote value",
			in:   `<ick-tstsniph2 Test=1 a="y"/>`,
			want: `<span name="sniph2" a="y"></span>`,
			err:  false},
		{name: "mixed quotes value",
			in:   `<ick-tstsniph2 Test=1 a="y'z;"/>`,
			want: `<span name="sniph2" a="y'z;"></span>`,
			err:  false},
		{name: "string with quote value",
			in:   `<ick-tstsniph2 Test=1 a=y'z/>`,
			want: `<span name="sniph2" a="y'z"></span>`,
			err:  false},
		{name: "html value",
			in:   `<ick-tstsniph2 Test=1 a="<ok></>"/>`,
			want: `<span name="sniph2" a="<ok></>"></span>`,
			err:  false},

		{name: "text + embedding + text",
			in:   ` Hello <ick-tstsniph2 Test=1/> folks <ick-tstsniph2 Test=1/> ! `,
			want: ` Hello <span name="sniph2"></span> folks <span name="sniph2"></span> ! `,
			err:  false},

		{name: "simple embedding with attributes",
			in:   `<ick-tstsniph2 Test=1 a b=1 c="x"/>`,
			want: `<span name="sniph2" a b=1 c="x"></span>`,
			err:  false},
		{name: "multi embedding with attributes",
			in:   `<ick-tstsniph2 Test=1 a b=1 c="x"/><ick-tstsniph2 Test=1 c="y" d/>`,
			want: `<span name="sniph2" a b=1 c="x"></span><span name="sniph2" c="y" d></span>`,
			err:  false},
		{name: "empty attributes",
			in:   `<ick-tstsniph2 Test=1 empty1='' empty2="" bool/>`,
			want: `<span name="sniph2" bool empty1 empty2></span>`,
			err:  false},
		{name: "setup attributes",
			in:   `<ick-tstsniph2 Test=2 class='text' d='test'/>`,
			want: `<div name="sniph2" class="ts2a ts2b" a2 d="test" style="display=test;" tabindex=2></div>`,
			err:  false},
	}

	// restet the component registrey for tests
	ResetRegistry()
	AddRegistryEntry("ick-tstsniph2", &sniph2{})

	output := new(bytes.Buffer)
	for i, tst := range tstset {
		output.Reset()
		err := renderHTML(output, nil, *ToHTML(tst.in))
		if tst.err {
			if err == nil {
				fmt.Printf("T%v) %s\n", i, tst.name)
				fmt.Println("   no error but error expected!")
				t.FailNow()
			}
		} else {
			if err == nil {
				if tst.want != output.String() {
					fmt.Printf("T%v) %s\n", i, tst.name)
					fmt.Println("   want:", tst.want)
					fmt.Println("   out :", output.String())
					t.FailNow()
				}
			} else {
				fmt.Printf("T%v) %s\n", i, tst.name)
				fmt.Println("   unexpected error:", err.Error())
				t.FailNow()
			}
		}
	}
}

type sniph3 struct {
	BareSnippet
	Text string
	Test int
}

func (s *sniph3) BuildTag() Tag {
	switch s.Test {
	case 0:
		s.Tag().SetTagName("span")
	case 1:
		return EmptyTag()
	case 2:
		t := NewTag("div")
		t.ParseAttributes(`class="ts2a ts2b" tabindex=2 style="display=test;" a2`)
		return *t
	}
	return *s.Tag()
}

func (s *sniph3) RenderContent(out io.Writer) (errx error) {
	switch s.Test {
	case 1:
		_, errx = RenderStringIf(s.Text != "", out, s.Text)
	}
	return errx
}

func TestRenderHTML_3(t *testing.T) {

	tstset := []struct {
		name string
		in   string
		want string
		err  bool
	}{
		{name: "empty span",
			in:   `Hello <ick-tstsniph3 Test=0/>`,
			want: `Hello <span name="sniph3"></span>`,
			err:  false},
		{name: "simple body, no tagnam",
			in:   `Hello <ick-tstsniph3 Test=1 Text="Bob"/>`,
			want: `Hello Bob`,
			err:  false},
		{name: "attributes with html syntax",
			in:   `Hello <ick-tstsniph3 Test=1 Text="<br/>"/>`,
			want: `Hello <br/>`,
			err:  false},
		{name: "two components",
			in:   `Hello <ick-tstsniph3 Test=1 Text="Bob"/>, Hello <ick-tstsniph3 Test=1 Text="Alice"/>`,
			want: `Hello Bob, Hello Alice`,
			err:  false},

		{name: "setup attributes",
			in:   `Hello <ick-tstsniph3 Test=0 style='color=red;' tabindex=1 a=0 class='text'/>`,
			want: `Hello <span name="sniph3" class="text" a=0 style="color=red;" tabindex=1></span>`,
			err:  false},
		{name: "overloading class", // not possible with testsnippet2, forced by the TagBuilder
			in:   `<ick-tstsniph3 Test=2 class='text'/>`,
			want: `<div name="sniph3" class="ts2a ts2b" a2 style="display=test;" tabindex=2></div>`,
			err:  false},
		{name: "overloading tabindex", // not possible with testsnippet2, forced by the TagBuilder
			in:   `<ick-tstsniph3 Test=2 tabindex=1/>`,
			want: `<div name="sniph3" class="ts2a ts2b" a2 style="display=test;" tabindex=2></div>`,
			err:  false},
		{name: "overloading style", // not possible with testsnippet2, forced by the TagBuilder
			in:   `<ick-tstsniph3 Test=2 style='color=red;'/>`,
			want: `<div name="sniph3" class="ts2a ts2b" a2 style="display=test;" tabindex=2></div>`,
			err:  false},
	}

	// restet the component registrey for tests
	ResetRegistry()
	AddRegistryEntry("ick-tstsniph3", &sniph3{})

	// running tests
	out := new(bytes.Buffer)
	for i, tst := range tstset {
		out.Reset()
		err := renderHTML(out, nil, *ToHTML(tst.in))
		if tst.err {
			if err == nil {
				fmt.Printf("T%v) %s\n", i, tst.name)
				fmt.Println("   no error but error expected!")
				t.FailNow()
			}
		} else {
			if err == nil {
				if tst.want != out.String() {
					fmt.Printf("T%v) %s\n", i, tst.name)
					fmt.Println("   want:", tst.want)
					fmt.Println("   out :", out.String())
					t.FailNow()
				}
			} else {
				fmt.Printf("T%v) %s\n", i, tst.name)
				fmt.Println("   unexpected error:", err.Error())
				t.FailNow()
			}
		}
	}
}

type sniph4 struct {
	BareSnippet
}

func (s *sniph4) RenderContent(out io.Writer) (errx error) {
	return RenderChild(out, s, ToHTML("<ick-tstrecursive/>"))
}

func TestRenderHTML_4(t *testing.T) {

	// restet the component registrey for tests
	ResetRegistry()
	AddRegistryEntry("ick-tstrecursive", &sniph4{})

	out := new(bytes.Buffer)
	err := renderHTML(out, nil, *ToHTML("<ick-tstrecursive/>"))
	require.Error(t, err)
}
