# Steadybit Extension Scaffold

This repository contains a scaffold with a sample implementation of a [Steadybit extension](https://docs.steadybit.com/integrate-with-steadybit/extensions). You may find this repository helpful…

 - [When you want to understand what Steadybit extensions are](#understanding-the-extension-mechanism).
 - [When you want to build a Steadybit extension](#for-extension-authors)

Please follow one of the links above to move to the appropriate documentation sections.

## Understanding the Extension Mechanism

One of the best ways to understand the extension mechanism is to run an extension and experiment with its APIs. We have prepared Gitpod and GitHub codespaces setups to make this as easy as possible for you.

When you click one of these buttons, you will be directed to an online editor with a locally running extension, and the file `README.http` will open. This file contains documentation and HTTP calls you can execute to learn about extensions and this specific sample implementation.

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](http://gitpod.io/#https://github.com/steadybit/extension-scaffold/blob/main/README.http)


[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://github.com/codespaces/new?hide_repo_select=true&ref=main&repo=595972094)


## For Extension Authors

TODO

# Using this

 - template repository
 - download as zip file
 - codespaces
 - ejection procedure
 - add DOCKER_USERNAME and DOCKER_PASSWORD to secrets
 - add PERSONAL_ACCESS_TOKEN_USED_BY_CLA_FROM_ANSGAR as secret if you want to use the CLA support
 - Docker image published where/how/how to make the image publicly accessible
 - Helm chart
   - remember to bumb the version
     > Releasing charts...
     Error: error creating GitHub release steadybit-extension-scaffold-1.0.0: POST https://api.github.com/repos/steadybit/extension-scaffold/releases: 422 Validation Failed [{Resource:Release Field:tag_name Code:already_exists Message:}]
 - when are image tags set

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
