# aide-go
GO (lang) helpers lib

![gophers image](https://blog.pyyoshi.com/content/images/2016/09/gopherswrench.jpg)

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

## Setup

```
make setup
```

## Test

```
make test
```

## Coverage

```
make coverage
```
