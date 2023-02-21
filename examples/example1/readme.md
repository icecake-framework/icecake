# icecake example1

icecake can be used to build Single-Page Application (SPA) in Go in different ways:
1. Building and Enhancing a static HTML website
1. Building and Embedding Web Components on a stateless website
1. Building a full SPA with state and API calls

Example1 illustrate how to use icecake for Building and Enhancing a static HTML website: 

1. demonstrate the use of the go HTML templating package to build page content directly on the front-end. 
1. demonstrate how to generate HTML content from a markdown source, directly on the front-side.

## Build

Example1 requires a simple build of the wasm code where the source is located in ./wasm/example1.go.

To build it run the build_ex1 task:

```
# from the icecake root directory:
task -t ./build/Taskfile.yaml build_ex1
```

The build will be located in the ./website ditectory.

## Run

Because this is a static webpage, you can serve the ./website directory with any webserver. 

We've setup `liveserver` to serve the exemple directory. see `.vscode/settings.json`.

There's no index.html file so open `localhost:5510/example1/website/example1.html` URL.

