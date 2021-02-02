# aide-go
A go (lang) helper lib by [Hurb](https://www.hurb.com)

![gophers image](https://golang.org/doc/gopher/pencil/gopherswrench.jpg)

## Setup

```bash
make setup
```

## Using

### To use this lib, in your app root path, get this package using:

```bash
go get github.com/hurbcom/aide-go/lib
```

or

```bash
go get gopkg.in/hurbcom/aide-go.v1
```

### In your go file, import this using:

```go
import "github.com/hurbcom/aide-go/lib"
```

or

```go
import "gopkg.in/hurbcom/aide-go.v1"
```

### And use like this:

```go
fmt.Printf("Testing build at ", lib.BeginningOfToday())
```

## Test

```bash
make test
```

## Coverage

```bash
make coverage
```
