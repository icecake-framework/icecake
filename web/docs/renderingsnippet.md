# Rendering icecake snippet

An rendered icecake snippet is nothing more than an HTML string with or without some webassembly code attached to it and allowing browser event hendling.

## Focus on rendering the HTML string of the snippet

The `HTMLSnippet` struct provides several methods to sipmlify setup of Snippets and to generate the related HTML string according to multiple use cases.

A custom HTML snippet is either an instance of the `HTMLSnippet` srtuct either a custom struct embedding at least the `HTMLSnippet` and any other required **custom properties** used to set it up. This custom HTML snippet can also **redefine (overload) some interfaces** to customize the HTML rendering. Finally a custom HTML snippet can define its **own html tag** allowing easy insert into usual html syntax.

For simplest cases an HTML snippet can be rendering directly by , 

Then somewhere in your code you need to call one HTML rendering method or/and to include the custom tag of your snippet in an HTML string.

> NOTA: custom HTML tag of the icecake framework always starts by `ick-` so we use to call them the `ick-tag`. For example the core icecake button has the ick-tag `ick-button`.

### Settingup a custom HTMLSnippet



### Rendering the HTML string of the snippet

**``HTMLComposer`` interface:**

- ``WriteSnippet(writer, HTMLComposer)`` WriteSnippet writes the HTML string of the composer, its tag element and its body, to the writer. The body is unfolded to look for sub-snippet and ever sub-snippet are also written to the writer.

- ``UnfoldHTML(writer, HTMLString, DataState)`` is used to unfold an HTML string without creating an enclosing tag. 
If one property of a component is not a simple text string but if it's an HTML String and if this string can contains `ick-tag` then UnfoldHTML can be used to . 

- ``RenderSnippet(writer, HTMLComposer)`` TO COME.


**``HTMLSnippet`` methods:**

- ``HTMLSnippet.WriteChildSnippet(HTMLComposer)`` WriteChildSnippet writes the HTML string of the composer, its tag element and its body, to the writer. The body is unfolded to look for sub-snippet and ever sub-snippet are also written to the writer. If the child request an ID, WriteChildSnippet generates an ID by prefixing its parent id. In addition the child is appended into the list of sub-components.

- ``HTMLSnippet.String()``

For example 

```go
s := NewSnippet(Tag{Name: "strong"}, `class="icecake-brand"`, `Icecake`)
RegisterComposer("ick-brand", )0
```

### About Snippet ID

Every Snippet has an id unless you've explicitly specified that it must not have one.
