# aide-go
A go (lang) helper lib by [Hotel Urbano](http://www.hotelurbano.com) 

![gophers image](https://blog.pyyoshi.com/content/images/2016/09/gopherswrench.jpg)

## Setup

```bash
make setup
```

## Using 

### To use this lib, in your app root path, get this package using:

```bash
go get github.com/hotelurbano/aide-go/lib
```

or

```bash
go get gopkg.in/hotelurbano/aide-go.v1
```

### In your go file, import this using:

```go
import "github.com/hotelurbano/aide-go/lib"
```

or

```go
import "gopkg.in/hotelurbano/aide-go.v1"
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
