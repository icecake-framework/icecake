## Icecake Interfaces

**ickcore.RMetaProvider**

	. RMeta() *RMetaData

**html.TagBuilder**
	- RMetaProvider

	. BuildTag() Tag
	. SetAttribute(name string, value string)

**html.ContentComposer**
	- RMetaProvider

	. RenderContent(out io.Writer) error

**html.ElementComposer**
	- TagBuilder
	- ContentComposer

**dom.UIComposer**
    - ElementComposer

	. Wrap(js.JSValueProvider)
	. AddListeners()
	. RemoveListeners()

