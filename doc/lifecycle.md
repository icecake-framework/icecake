# icecake component lifecyle

## WIP25

RenderComponent `C`

1. create the HTML component
    1. lookup for registered component properties: ickname, typ, css
    2. get a fresh component unique id
    3. extract `C.Container()` properties: tagname, classes, attributes
    4. **create the HTML element**
    5. set the container classes and attributes
    6. wrap the HTML element to the component
    7. set the id
    8. > report class & attributes
    9. create the style element into document's head
2. name the component
3. unfold `C.Body()` template
    1. parse the template
    2. Execute the template with the _data
    3. scan the generated html looking for an embedded and registered`ick-*` component
    4. for every component found
        - instantiate the component object
        - **create the HTML component**
        - extract any embedded attributes
        - force setting component ID
        - for every embedded attribute
            if the attribute's name match with a _data object property name
                - set attribute's value to the _data object property
            otherwise
                - set the attribute to the new component
        - recursively unfold embedded component 
        - fill unfolding string with the outerhtml of the component
4. set InnerHTML with the unfolded body html
5. insert the HTML component into the DOM
6. add event listeners for every unfolded component
7. add event listeners for `C`
8. show every unfolded component
9. show `C`

---

## WIP26

RenderComponent `C`

A. compose HTML
B. mount DOM ELEMENT

A. compose `C` HTML

1. check the recursive depth
2. Get a unique id
3. open the HTML Element
    1. extract `C.Container()` properties: tagname, classes, attributes, styles
    2. add container's properties to existing component's properties
    3. write to the buffer
3. unfold the HTML body 
    1. parse the template
    2. Execute the template with the _data
    3. scan the generated html looking for an embedded and registered`ick-*` component
    4. for every component found
        - instantiate the `C.SUB` component object
        - extract any embedded attributes
        - force setting component ID
        - for every embedded attribute
            if the attribute's name match with a _data object property name
                - set attribute's value to the _data object property
            otherwise
                - set the attribute to the new component
        - recursively compose `C.SUB` HTML
4. close the HTML Element
5. return the component HTML string

B. Mount `C` DOM.Element

1. create the HTML Element
2. set inner `C` HTML
3. create the global component Style Element
4. set `C` style
5. add `C` listeners

