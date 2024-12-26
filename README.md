# circlebuffer

[![Go Reference](https://pkg.go.dev/badge/cattlecloud.net/go/circlebuffer.svg)](https://pkg.go.dev/cattlecloud.net/go/circlebuffer)
[![License](https://img.shields.io/github/license/cattlecloud/circlebuffer?color=7C00D8&style=flat-square&label=License)](https://github.com/cattlecloud/circlebuffer/blob/main/LICENSE)
[![Build](https://img.shields.io/github/actions/workflow/status/cattlecloud/circlebuffer/ci.yaml?style=flat-square&color=0FAA07&label=Tests)](https://github.com/cattlecloud/circlebuffer/actions/workflows/ci.yaml)

`circlebuffer` provides a modern [circular buffer](https://en.wikipedia.org/wiki/Circular_buffer) Go library with support for generics.

### Getting Started

The `circlebuffer` package can be added to a project with `go-get`.

```shell
go get cattlecloud.net/go/circlebuffer@latest
```

```go
import "cattlecloud.net/go/circlebuffer"
```

### Examples

##### Inserting elements

```go
buf := circlebuffer.New[string](1024)
buf.Insert("alice")
buf.Insert("bob")
```

##### Iterating elements

```go
for item := range buf.All() {
  // ...
}
```

### License

The `cattlecloud.net/go/circlebuffer` module is open source under the [BSD](LICENSE) license.
