### Readme

<ick-icecake-brand/> can be used to build Single-Page Application (SPA) in Go in different ways:

1. Building and enhancing a static HTML website
1. Building and embedding Web Components on a stateless website
1. Building a full SPA with state and API calls

This example illustrates how to use <ick-icecake-brand/> for building and enhancing a static HTML website: 

1. uses the go HTML templating package to build page content directly on the front-end. 
1. generates HTML content from an embedded markdown source, directly on the front-side.

The markdown source includes two components wich are rendered with the icecake markdown extension:

`<ick-icecake-brand/>`: <ick-icecake-brand/> 

`<ick-button/>`: <ick-button Title="See Other Examples" HRef="/" Class="is-link is-light is-small is-outlined"/> <ick-button Title="See Source Code" HRef="https://github.com/icecake-framework/icecake/blob/main/examples/example1/wasm/example1.go" Class="is-link is-light is-small is-outlined"/>

### Build

This example requires a simple build of the wasm code where the source is located in `./examples/example1/wasm/example1.go`

To build it run the `build_eample` task as follow:

```
# from the icecake root directory:
EXAMPLE=example1 task -t ./build/Taskfile.yaml build_example
```

The build will be located in the `./examples/website` ditectory.

### Run

Because this is a static webpage, you can serve the `./examples/website` directory with any webserver. 

We've setup `liveserver` to serve it, see `.vscode/settings.json`.

Open `localhost:5510/example1.html` URL.
