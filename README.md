# Icecake Go Wasm Framework

Icecake is a framework designed to build Web Assembly SPA 100% in GO.

Icecake is an experimental project aiming to implement web assembly technologies in go :
- web assembly [see wasm doc](https://developer.mozilla.org/fr/docs/WebAssembly)
- fullstack in go (no JS) [see this post to use wasm in GO](https://tutorialedge.net/golang/writing-frontend-web-framework-webassembly-go/)

## Why

Ealy 2023 webassembly generated code is smoosly becoming the new standard for running frontend web applications. But although [webassembly source code for front can be easily written in C++, C# or Rust](https://www.webassemblyman.com/webassembly_front_end_web_development.html), there's still few solutions to write it in GO.

Existing solutions are too restrictive or too exploratory, furthermore nothing exists when it comes to build a nice app with components as we can do with Angular, React or Vue, and mixing one of these framework to generate webassembly code does not make sens.

So the idea here with icecake is to provide a native GUI framework to develop nice and modern SPA very quickly in GO.

Sources of inspiration: 
- https://github.com/maxence-charriere/go-app
- https://github.com/gowebapi/webapi

## Backlog

### Front side

- [ ] pages and routes

### Back side

- [ ] "hello world" wasm served by a SPA server, with dev environment setup.

## Tech

- Go 1.20 and it's wasm compiler
- based on the ``syscall/js`` package
- CSS responsive framework, without any JS code: [Bulma](https://bulma.io/)

## Project layout

The Directory structure follows [go ecosystem recommendation](https://github.com/golang-standards/project-layout).

```bash
icecake
│   ├── readme.md
│   └── .gitignore
│
├── configs                         # congigurations files, loaded at server startup
│   ├── dev.env                 
│   └── prod.env          
│
├── build                           # build scripts
│   └── Taskfile.yaml               # building task configuration, ic. autobuild the front
│
├── doc                             # overall doc
│   └── [*.*]                       
│
├── cmd
│   └── icecake                     # the icecake CLI command required to run the SPA server
│       └── icecake.go          
│
├── pkg
│   ├── spaserver                   # SPA server 
│   │   ├── webserver.go                   
│   │   └── middleware.go                   
│   ├── ick                         # icecake package with framework primitives, ic WebAPI embedded 
│   │   └── [*.go]                   
│   ├── spasdk                      # SDK for any SPA client willing to call SPA APIs
│   │   └── [*.go]                   
│   ├── dom                         # 
│   │   └── [*.go]                   
│
├── web                             # source codes and assets required by the front, even server side described in the go file
│   ├── static
│   │   ├── wasm_exec.js            # this file is mandatory and is provided by the go compiler
│   │   ├── wasm_spa.js             # WebAssembly initialization for this app
│   │   ├── index.html              # SPA icecake index file
│   │   ├── icecake.css             # CSS file for icecake
│   │   └── [*.*]                   # any other static assets like img
│   └── wasm
│       └── main.go                 # the front app entry point, uses components
│
├── website                         # the self sufficient dir to serve the app in production, built with prod tasks (see Taskfile.yaml)
│   ├── *.*
│
├── tmp                             # built with dev tasks (see Taskfile.yaml)
│   ├── *.*

```

### About Web Assembly with go

Some documentation available here https://tinygo.org/docs/guides/webassembly/ and here https://github.com/golang/go/wiki/WebAssembly

Go provides a specific js file called `wasm_exec.js` that need to be served by your webpapp. This file mustbe part of the static assets to be served by the server. To get the latest version you can _extract_ it from you go installation: `cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./web/static`

## Development

We use ``task`` as task runner. See [Taskfile Installation](https://taskfile.dev/installation/) doc to install it.

In development mode run the `dev_front` task from the root path with the `--watch flag`:

```bash
$ task -t ./build/Taskfile.yaml dev_front --watch
```

This task:
1. moves any changed files in ``./web/static/`` to ``./tmp/website/``
1. builds/rebuilds any frontend components and the .wasm file
1. builds/rebuilds the ``./tmp/website/spa.wasm`` file according to changes in the ``web/wasm/main.go``

Start the server either in debug mode with the `F5` in vscode, or by running the `dev_back` task:

```bash
$ task -t ./build/Taskfile.yaml dev_back
```

### Editor Configuration

If you are using Visual Studio Code, you can use workspace settings to configure the environment variables for the go tools.

Your settings.json file should look something like this:

```json
{
    "go.toolsEnvVars": { "GOARCH": "wasm", "GOOS": "js" }
}
```

## Testing

Useful read about [go data race detector](https://go.dev/doc/articles/race_detector#How_To_Use)

To be able to test wasm code on the browser, you need to install [wasmbrowsertest](https://github.com/agnivade/wasmbrowsertest):

```bash
$ go install github.com/agnivade/wasmbrowsertest@latest
$ mv $(go env GOBIN)/wasmbrowsertest $(go env GOBIN)/go_js_wasm_exec
```

and if you're working on WSL you need to customize and create a command to enable lanching chrome from the command line:

[see build/google-chrome/google-chrome.go](./build/google-chrome/google-chrome.go)

```bash
cd build
go build -o ~/google-chrome ./google-chrome/google-chrome.go
```

Run the `unit_test` task to run both testing pkg and wasm:

```bash
$ task -t ./build/Taskfile.yaml unit_test
```

## Licence

[LICENCE](LICENCE)
 