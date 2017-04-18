# Please building a non-trivial Golang project


This repo is an PoC project structure with [Please.build](http://please.build) for Golang apps.

### Outline:

  - `third_party/go` - vendor packages, will be available to our application code during builds / tests.
  - `tests/` - folder for integration tests
  - `proto/` - protobuff IDL, we restrict code generation just to Golang

### Goals:

  - support fast incremental builds
  - support IDE / Editor autocomplete features


### Approaches / Workarounds

Currently there is no way to make editors support the dependency declaration from plz-BUILD files. By limiting our supported languages just to Golang and respecting the expected package folder structure dictated by Go, we could get very fast incremental build speeds for tests / packages generation AND keep the convenience of good editor support.


### Requirements:
  - working installation of `plz` (http://please.build)
  - properly setup `direnv` (https://direnv.net/)
  - Visual Studio Code with Golang support (https://code.visualstudio.com/) + (https://github.com/Microsoft/vscode-go)
  - direnv provides the right ENV variables in shell, editors are started from via terminal integration, eg: `code` for Visual Studio Code.

-----------

#### Problem:

support autocomplete for third_party/go packages

Solution:

use [direnv](https://direnv.net/) to augment GOPATH, like:

    $ echo .envrc
    > export GOPATH=$PWD/plz-out/gen/third_party/go

-----------

#### Problem:

support autocomplete for our application packages

Solution:

This would require our packages to be in Go workspace, like $GOPATH/src. We could nest our golang code in a top-level `src` folder and tell `direnv` to set GOPATH to the root our our repo. `Plz` would then control `src/` folder and things should kinda work.

    $ echo .envrc
    > export GOPATH=$PWD:$PWD/src/plz-out/gen/third_party/go:$GOPATH


-----------

#### Problem:

now autocompletion on generated artefacts is not working, e.g.: proto. The actual `kitten.pb.go` file is in `plz-out/gen/proto/kitten.pb.go`. It would be great, if we could trick the autocomplete daemon for Golang to somehow use (some?) of those folders...

Solution:

???