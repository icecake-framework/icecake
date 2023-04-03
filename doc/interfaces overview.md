# Interfaces Overview


## HTMLComposer

```go
type HTMLComposer interface {
	Id() string
	Template(_data *DataState) SnippetTemplate
	SetAttribute(key string, value String, overwrite bool)
	Attributes() String
	Embed(id string, cmp any)
	Embedded() map[string]any
}
```

## JSValueProvider

```go
type JSValueProvider interface {
	Value() JSValue
}
```

## JSValueWrapper

```go
type JSValueWrapper interface {
	Wrap(JSValueProvider)
}
```

## UIComposer

```go
type UIComposer interface {
	html.HTMLComposer
	js.JSValueWrapper
	Mount()
	AddListeners()
	RemoveListeners()
	UnMount()
}
```