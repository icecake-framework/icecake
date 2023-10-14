## Icecake Interfaces

**ickcore.RMetaProvider**

	. RMeta() *RMetaData

**ickcore.Composer**
	- RMetaProvider

	. NeedRendering() bool

**ickcore.TagBuilder**
	- Composer

	. SetAttribute(name string, value string)
	. BuildTag() Tag

**ickcore.ContentComposer**
	- Composer

	. RenderContent(out io.Writer) error

**ickcore.ElementComposer**
	- TagBuilder
	- ContentComposer

**dom.UIComposer**
    - ElementComposer

	. Wrap(js.JSValueProvider)
	. AddListeners()
	. RemoveListeners()

