# Tent

Tent is a tiny webserver intended for serving files and directories from
a single path on the filesystem. Kinda like Python's http.server but with better
defaults.

No non-standard dependencies. Essentially an http.ServeFile() wrapper.

## Installation

Tent can be compiled with `make` or `go build`, and installed system-wide by
running `make install` with root-level permissions.

## Usage

```
Usage of ./tent:
  -host string
        IP address to listen on (default "127.0.0.1")
  -port uint
        Port to listen on (default random)
  -path string
        Absolute or relative path to serve (default "./")
```
