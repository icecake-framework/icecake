
icecake can be used to build Single-Page Application (SPA) in Go in different ways:

1. Building and Enhancing a static HTML website
1. Building and Embedding Web Components on a stateless website
1. Building a full SPA with state and API calls

Example3 illustrates how to use icecake for Building and Enhancing a static HTML website:


1. demonstrate simple UI components
1. demonstrate simple embedded components


## Build

Example3 requires a simple build of the wasm code where the source is located in ./wasm/example3.go.

To build it run the build_ex3 task:

```go
# from the icecake root directory:
task -t ./build/Taskfile.yaml build_ex3
```

The build will be located in the ./website ditectory.

## Run

Because this is a static webpage, you can serve the ./website directory with any webserver. 

We've setup `liveserver` to serve the exemple directory. see `.vscode/settings.json`.

There's no index.html file so open `localhost:5510/example3/website/example3.html` URL.

