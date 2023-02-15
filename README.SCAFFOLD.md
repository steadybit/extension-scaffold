# README

# Using this

 - template repository
 - download as zip file
 - clear changelog/contributing/README/update LICENSE/remove CLA step? (TODO via makefile?)
 - add DOCKER_USERNAME and DOCKER_PASSWORD to secrets
 - add PERSONAL_ACCESS_TOKEN_USED_BY_CLA_FROM_ANSGAR as secret if you want to use the CLA support
 - codespaces
 - Docker image published where/how/how to make the image publicly accessible
 - Helm chart

## Getting started

Make sure that you're in the root of the project directory, fetch the dependencies with `go mod tidy`, then run the application using `go run ./cmd/web`:

```
$ go mod tidy
$ go run ./cmd/web
```

Then visit [http://localhost:4444](http://localhost:4444) in your browser.

## Configuration settings



## Admin tasks

The `Makefile` in the project root contains commands to easily run common admin tasks:

|     |     |
| --- | --- |
| `$ make tidy` | Format all code using `go fmt` and tidy the `go.mod` file. |
| `$ make audit` | Run `go vet`, `staticheck`, execute all tests and verify required modules. |
| `$ make build` | Build a binary for the `cmd/web` application and store it in the `bin` folder. |
| `$ make run` | Build and then run a binary for the `cmd/web` application. |

## Changing the module path

The module path is currently set to `foobar`. If you want to change this please find and replace all instances of `foobar` in the codebase with your own module path.
