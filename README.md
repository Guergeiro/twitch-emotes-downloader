# twitch-emotes-downloader

Downloads twitch emotes in bulk from https://www.twitchmetrics.net/

## Table of Contents

- [Quick Start](#quick-start)
- [Using Source](#using-source)
- [Creator](#creator)
- [License](#license)

## Quick Start

1. Head into the releases page.
2. Download the corresponding binary to your OS.
3. Run it.

```console
# To see how to use and extra flags
$ twe-dl --help

# To run with default values
$ twe-dl

# To run with one or more urls
$ twe-dl https://www.twitchmetrics.net/emotes \
    https://www.twitchmetrics.net/c/26261471-asmongold/emotes \
    https://www.twitchmetrics.net/c/149747285-twitchpresents/emotes
```

## Using Source

twitch-emotes-downloader is an open source project, distributed under a
[GPLv3 license](./LICENSE). This document explains how to check out the sources,
build them on your own machine, and run them.

This project uses [Go](https://go.dev) as a programming language, so you need to
install it first.

1. Clone the repository.
2. Install its dependencies.
3. Execute `go run main.go`

## Creator

[Breno Salles](https://brenosalles.com)

## License

By contributing your code, you agree to license your contribution under the
[GPLv3](./LICENSE).
