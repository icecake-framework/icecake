# Icecake Go Wasm Framework 
## Example3

<div class="subtitle">Basic use of components for building a static HTML website</div>

In this example we use the ui `Notify` component. This component uses the [BULMA Notification CSS classes](https://bulma.io/documentation/elements/notification/) and has several properties: 

``` Go
// The Message to display within the notification
Message string

// The Timeout duration after which the notification will close automatically.
// The timer starts when the notification pops up.
// Is not set, the notification will not close automatically
Timeout time.Duration
```
### Insert an UI component into the DOM, inside an element ID

Instantiate the Notify component and setup the Message property:

```Go
notif := &ui.Notify{
    Message: `This is a typical notification message <strong>including html <a href="#">link</a>.</strong> 
    Use the closing button on the right corner to remove this notification.`,
    }
```

Setup the `InitClasses` and the `InitAttributes`default properties of the component to choose for the [BULMA color](https://bulma.io/documentation/customize/variables/) and to add `aria`, `role` or other attributes.

``` Go
notif.InitClasses = ick.ParseClasses("is-warning is-light")
notif.InitAttributes, _ = ick.ParseAttributes("role='alert'")
```

Call the `RenderComponent()` function to insert it into the DOM:

```HTML
<div id="notif_container" class="block"></div>
```

```Go
webapp.ChildById("notif_container").RenderComponent(notif, nil)
```

<strong>Test it:</strong>

<button class="button is-warning" id="btn1">Insert Notity Component</button>
<button class="button is-danger" id="btn2">Auto Closing Notification</button>
<button class="button is-success is-light" id="btn3">Toast Notification</button>

<div id="notif_container" class="block"></div>

### Embed an UI component within an HTML template

Define the HTML template and Render it inside an element ID:

```Go
html := `<div class="box">
<p class="pb-2">This is an html template object embedding the &lt;ick-notify&gt; element.</p>
<div class="block">
    <ick-notify Message="This message comes from the Notify Component <strong>embedded into an html template</strong>."
    class="is-info is-light"
    role="success"/>
</div>
</box>`

webapp.ChildById("ex3_container").RenderHtml(html, nil)
```

<strong>Test it:</strong>

<button class="button is-info" id="btn4">Embed into Template</button>

<div id="ex3_container" class="block"></div>

