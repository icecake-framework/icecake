# Icecake Framework Documentation

The Icecake Framework documentation is composed of a website available at https://icecake.dev and a full go package documentation available at https://pkg.go.dev/github.com/icecake-framework/icecake. This go package documentation is automatically generated when the package is published. 

The published website is generated within the `/docs` directory and served by github pages. Sources of this website are in `/website/docs` directory. 

[Go doc comments](https://go.dev/doc/comment) are written directly at package level within the `/pkg` and `/internal` directories. 

## Rebuild the doc website

Run the command `task -t ./build/Taskfile.yaml makedocwebsite` to rebuild all the wasm files and the full `/website/docs`.
