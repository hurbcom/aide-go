# aide-go
A go (lang) helper lib by [Hotel Urbano](http://www.hotelurbano.com) 

![gophers image](https://blog.pyyoshi.com/content/images/2016/09/gopherswrench.jpg)

## Setup

```
make setup
```

## Using 

### To use this lib, in your app root path, get this package using:

```
go get github.com/hotelurbano/aide-go/lib
```

### In your go file, import this using:

```go
import "github.com/hotelurbano/aide-go/lib"
```

or

```go
import (
    "github.com/hotelurbano/aide-go/lib"
)
```

### And use like this:

```go
fmt.Printf("Testing build at ", lib.BeginningOfToday())
```

## Test

```
make test
```

## Coverage

```
make coverage
```
