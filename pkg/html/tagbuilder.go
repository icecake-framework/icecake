package html

type TagBuilder interface {
	// Tag returns a valid reference to a tag.
	Tag() *Tag

	// BuildTag builds the tag used to render the html element.
	// This tag builder can update the given tag or overwrite its properties.
	// The composer rendering processes call BuildTag once.
	// If the implementer builds an empty tag, only the body will be rendered.
	BuildTag(tag *Tag)
}
