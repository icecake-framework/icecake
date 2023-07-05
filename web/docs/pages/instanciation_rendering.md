## Instanciation and Rendering a snippet

According to your goal there's multiple way to instantiate and to render a snipper: 

- you need a simple html template and generate an html string interpretable by browsers. This is usefull to format a content according to some properties,
- you need to unfold the same content many times and avoid to duplicate it,
- you need to render the template content of an SPA or blog pages on the server side and to be able to listen events on your snippet,
- or your need a standalone component...

And in every case your snippet can be used by other snippet, but with or without a dependance between each other. We call it a parent/child relationship. So every snippet can be instantiated and rendered as an orphan or as a child of another composer.

TODO: doc: examples of parent/child relationship

In both cases a snippet can be instantiated with the snippet factory or by [unfolding an `ick-tag`], in this former case your code can not access the snippet instance directly like for a snippet _hidden_ within another one.

We'r considering the following types of cases:

1. the snippet is fully instantiated on the front or only to render the static page, or on both.
1. the snippet is instantiated at the top level, usually within the &lt;body&gt; element of a page, or hidden.
1. the snippet needs to listen events or not.


### Two faces of a snippet

Every snippet is represented by an html string. This string can be rendered ont the server side or direclty on the front. 

Then a snippet can interact with the UI by listening events or by updating its html representation. In this case the wasm code running on the front needs to access the instance of the snippet. You may have a reference to it or not, usually when the snippet has already been rendered on loaded page or when it is hidden. So it's common to get its corresponding DOM element either by its ID attribute, either by its name attribute or its class attribute. 

The snippet ID can be assigned to the snippet by the process instantiating it and must ensure its uniqueness. This process could be:
- the top level rendering of the static page builder. It's up to the dev to choose the Id.
- the rendering of another snippet (within the static page builder, or dynamically on the front). The ID is a sub-ID assigned by the parent.
- the unfolding process of an `ick-tag`. In this case the ID property of the `ick-tag` should be build dynamically to avoid duplicate.

Then Icecake provides a generic function to wrap a DOM element by its ID to any snippet. This function is `WrapById`. 

The wasm composer must implement the `dom.UIComposer` interface with a `Wrap`, a `AddListeners` and a `RemoveListeners` methods.

The wrapped snippet is unaware of the properties and unable to get it. Guessing some properties from rendered attributes could be done in some case if it was intented. So, properties of a wrapped snippet are initialized with their zero go value and may not reflect the render. Snippet render `IsWrapped` metadata is turned-on and may be used to know that these propeties may not reflect the render.

### Summary

| Snippet Instanciation                    | use cases  | Id    | HTML rendering | Event Listening   | 
| -                                        | -          | -     | -                 | -                     |
| **within the static page builder:**
|   - no need to listen events             | - format a content<br>- duplicate same content | useless | - RenderChild | N/A
|   - need to listen events                | - page template | ID assigned by the dev | - RenderChild | `WrapById()`, useless properties. |
| **within another snippet:** 
|   - no need to listen events             | - format a content<br>- duplicate same content | useless | - RenderChild<br>- Unfolding&nbsp;`ick-tag` | N/A
|   - need to listen events                | - standalone snippet | automatically assigned by the ick-apps  | - RenderChild<br>- Unfolding&nbsp;`ick-tag` | listeners automatically added by the ick-app loader
|   - need to listen events                | - sub-snippet | sub-ID assigned by the parent | - RenderChild<br>- Unfolding&nbsp;`ick-tag` | `WrapById()` call within the `WrapById()` of the parent.

_NOTA: Unfolding `ick-tag` is also the way to render a snippet within a markdown content._

[unfolding an `ick-tag`]: https://icecake-framework.html/unfolding-ick-tag
