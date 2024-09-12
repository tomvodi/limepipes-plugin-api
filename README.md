# limepipes-plugin-api

LimePipes plugin API is a repository that contains the [gRPC](https://grpc.io/) API definition for LimePipes plugins. 
These plugins are used to parse and/or export different file formats (Bagpipe Music Writer, MusicXML, ...).
The underlying communication is done via [Hashicorp's Go Plugin System](https://github.com/hashicorp/go-plugin).

## Project Structure

This repository contains the API definition for LimePipes plugins and a music model that represents the parsed tunes 
from the supported file formats. The returned data structure from file parsing is called the music model which is 
a common representation of the parsed tunes. The music model structs contain struct tags to be exported to `.yaml` files.

### Directory Structure

`proto`

Here are the protobuf files that define the gRPC API for LimePipes plugins and the music model.

`musicmodel`

The output directory for the generated code from the protobuf files for the music model. All files that do not end 
with `.pb.go` store utility methods for objects from the model. Many structs have a corresponding `.yaml.go` file
which contains the `MarshalYAML` and `UnmarshalYAML` methods to serialize and deserialize them to and from YAML.

`cmd`

This directory contains some test and example code to implement a LimePipes plugin and the LimePipes plugin host.

`plugin`

The generated code from the protobuf files for the LimePipes plugin API. This code is used to implement the LimePipes plugins.
It also contains the boilerplate code for [go-plugin](https://github.com/hashicorp/go-plugin) to create and use a plugin.

## Build

`buf generate` builds all the generated code from the protobuf files.


## Run

The limepipes application needs a configuration file called `limepipes.env` beside the executable or the environment
variables from this file directly set to the environment of the application.

## Develop

### Prerequisites

In order to build the protobuf implementations, you need to have the [buf](https://buf.build/) tool installed 
and in your PATH.

Mocks are generated with `[vektra/mockery](https://vektra.github.io/mockery/latest/installation)` this also has to
be in your PATH.

### Build

The `Makefile` contains targets for many build and test tasks. The most important ones are:

- `make build` cleans all generated files and builds the protobuf code
- `make format` formats all protobuf files
- `make test` runs all tests on go files
- `make check-coverage` checks the test coverage of the go files

### Quality assurance

The [buf](https://buf.build/) tool is used to lint and format the protobuf files. It is mandatory to have formatted and 
linted protobuf files in order to push changes to the repository. There is a GitHub action that checks the linting and 
formatting of the protobuf files. It also checks for breaking changes in the protobuf files so it should not be possible
to introduce breaking changes without noticing.

### Conventions

The API must utilize the return errors according to the [gRPC status codes](https://github.com/grpc/grpc/blob/master/src/proto/grpc/status/status.proto)
and the error codes defined in the [grpc status codes page](https://github.com/grpc/grpc/blob/master/doc/statuscodes.md) used by Google.

### Used libraries and tools

- [buf](https://buf.build/) for protobuf handling
- [Go Plugin System](https://github.com/hashicorp/go-plugin) for plugin handling
