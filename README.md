# Icecake Go Wasm Framework

:warning: This code is an Alpha version with work in progress.

Icecake is a framework designed to build Web Assembly SPA 100% in GO.

Icecake is an experimental project aiming to implement web assembly technologies in go :
- web assembly [see wasm doc](https://developer.mozilla.org/fr/docs/WebAssembly)
- fullstack in go (no JS) [see this post to use wasm in GO](https://tutorialedge.net/golang/writing-frontend-web-framework-webassembly-go/)

[![Go Reference](https://pkg.go.dev/badge/github.com/icecake-framework/icecake.svg)](https://pkg.go.dev/github.com/icecake-framework/icecake)

## Why

Ealy 2023 webassembly generated code is smoosly becoming the new standard for running frontend web applications. But although [webassembly source code for front can be easily written in C++, C# or Rust](https://www.webassemblyman.com/webassembly_front_end_web_development.html), there's still few solutions to write it in GO.

Existing solutions are too restrictive or too exploratory, furthermore nothing exists when it comes to build a nice app with components as we can do with Angular, React or Vue, and mixing one of these framework to generate webassembly code does not make sens.

So the idea here with icecake is to provide a native GUI framework to develop nice and modern SPA very quickly in GO.

Sources of inspiration: 
- https://github.com/maxence-charriere/go-app
- https://github.com/gowebapi/webapi

## Backlog

This repo is an Alpha version with many work in progress.

Current works on Alpha4 focus on:
- unfolding html containing ick-components tags, including custom attributes
- instantiating ick-components, embedded or not  
- rendering any html containing ick-component on the front side, with event handling
- provide examples to use UI Snippets 
- provide some core UI Snippets such as: 
    - button
    - card
    - image
    - input
    - message
    - notify

In parallel we're doing many refactoring to make code more coder friendly.

## Tech

- Go 1.21 and it's wasm compiler
- based on the ``syscall/js`` package
- Snippets use [Bulma](https://bulma.io/) CSS framework. Bulma is a pure CSS framework without any JS code.

## Examples

Open `examples/website/index.html` in any browser to run examples.

You can use the [vscode liverserver extension](https://marketplace.visualstudio.com/items?itemName%3Dritwickdey.LiveServer) to do so.


## Project directory layout

The Directory structure follows [go ecosystem recommendation](https://github.com/golang-standards/project-layout).

```bash
icecake
│   ├── readme.md
│   ├── go.mod
│   ├── go.sum
│   ├── LICENCE
│   └── .gitignore
│
├── build                           # build scripts
│   └── Taskfile.yaml               # building task configuration, ic. autobuild the front
│
├── cmd
│   └── icecake                     # the icecake CLI command required to run the SPA server
│       └── icecake.go          
│
├── configs                         # congigurations files, loaded at server startup
│   ├── dev.env                 
│   └── prod.env          
│
├── doc                             # overall icecake doc
│   ├── cheatsheet.ods              # overview of icecake structs and methods
│   └── [*.*]                       
│
├── internal
│   ├── helper
│   └── testwasm                    # wasm test environment
│
├── pkg
│   ├── *                           # see description of packages here under
│
├── web                             # source codes and assets required by the front
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

## Packages

| pkg           | description |
| -             | - |
| `browser`     | provides primitives to interact directly with the browser, its navigation hystory and the local or the session storages.
| `clock`       | provides a timer and a ticker with possibility to add callback functions at every tic and at the end of the timer.
| `console`     | provides helpers to raise enhanced messages in the browser console.
| `dom`         | provides primitives to interact with the DOM of a webpage. Traditional node, element, and document's methods can be call in go here. An UISnippet struct and an UIComposer Interface are provided to allow rendering of HTMLSnippet and to handle event listening.
| `event`       | defines all types of the dom event handlers with their methods
| `html`        | provides an HTMLSnippet struct and an Interfaces to allow html rendering of ick-name tags in different ways.
| `ick`         | core Snippets with html rendering
|  └── `ickui`  | UI of core Snippets with event handler. Compiles with the wasm compiler
| `ickcore`     | provides the render metadata provider and the global registry of Composers, this needs to be fully reworks.
| `icksdk`      | usefull functions to call API on the spaserver.
| `ickserver`   | provides a configurable webserver dedicated to serve an spa with wasm code.
| `js`          | ``syscall/js`` extended with ``console`` 
| `stringpattern` | provides functions to check validity of an HTML name such a a token name or the name of an attribute.



## About Web Assembly with go

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

### Testing wasm code

Testing front with wasm and DOM components can not be executed with the standard `go test` command. To test wasm code into a browser well we need to run a dedicated html page which will execute the wasm test code.

Wasm test code and environement is located into `internal/testwasm`.

Run the task `task -t ./build/Taskfile.yaml test_wasm` to build all `test*.go` files. `tests.go` is the wasm entry point that will be executed when you load the dedicated HTML page `internal/testswasm/index.html`

### Testing the back 

```bash
$ task -t ./build/Taskfile.yaml unit_test
```

_Useful read about [go data race detector](https://go.dev/doc/articles/race_detector#How_To_Use)_

## Licence

[LICENCE](LICENCE)
 