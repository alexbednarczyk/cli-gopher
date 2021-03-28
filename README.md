# cli-gopher

**[Table of Contents]**

- [Synopsis](#synopsis)
- [Motivation](#motivation)
- [Getting Started](#getting-started)
  - [Help And Available Commands](#help-and-available-commands)
- [Build Commands](#build-commands)
- [Built With](#built-with)
- [License](#license)

This is an example or starter providing guidance in creating CLIs using Golang.

## Synopsis

This repo provides a central location of guidance for best practices when creating CLIs using Golang

## Motivation

The goal of this repo is to create a location for folks to get information they need to create a CLI using Golang.

In the end this repo will provide guidance for the following:

- Basic and Advanced CLI examples
- Binary releases

## Getting Started

What things you need to install the software and how to install them

- [golang](https://golang.org/doc/install)

If you want to download and run, the [latest pre-built binaries available](https://github.com/alexbednarczyk/cli-gopher/releases/latest)

For Linux and MacOS with [Gatekeeper enabled](https://support.apple.com/en-us/HT202491)
```
chmod +x <binary file name>

./<binary file name>
```

### Help And Available Commands

A list of available commands

```
go run main.go
```

```
Example cli tool based on golang!

Usage:
  cli-gopher [command]

Available Commands:
  config      Simple config command
  count       A simple count command
  date        A simple date command
  env         A simple environvment command
  help        Help about any command
  password    A simple password command
  version     The version number of cli-gopher

Flags:
  -h, --help   help for cli-gopher

Use "cli-gopher [command] --help" for more information about a command.
```

## Build Commands

Build binary locally on OSX
`CGO_ENABLED=0 go build -ldflags="-X 'github.com/alexbednarczyk/cli-gopher/cmd.version=$(date +%Y.%U.%j)' -w -s" -v -o cli-gopher .`

## Built With

- [Golang](https://golang.org/) - Go, the open source programming language
- [Cobra](https://github.com/spf13/cobra) - A Commander for modern Go CLI interactions
- [go.1password.io/spg](https://pkg.go.dev/go.1password.io/spg) - A go package for strong password generation.
- [GitHub Actions](https://github.com/features/actions) - Automate your workflow from idea to produciton

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
