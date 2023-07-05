package html

type TagBuilder interface {

	// Meta returns a reference to render meta data
	RMetaProvider

	// BuildTag builds the tag used to render the html element.
	// The composer rendering processes call BuildTag once.
	// If the implementer builds an empty tag, only the body will be rendered.
	//
	// The returned tag can be built from scratch or on top of an embedded tag in the snippet.
	BuildTag() Tag

	// SetAttribute creates a tag attribute and set its value.
	// If the attribute already exists then it is updated.
	//
	// Attribute's name is case sensitive
	SetAttribute(name string, value string)
}
