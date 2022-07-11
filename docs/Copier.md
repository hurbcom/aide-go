# Copier
## A library for creating deep copies of structs during looping.

This library employs the same notation for tags as [jinzshu/copier](https://github.com/jinzhu/copier). However, this 
solution should be more efficient because it analyzes the destination struct at initialization time, which is executed 
at first struct copy. If no tags was provided, the library will use the field name to perform the copy.

## Usage

```go
package main

import (
	"fmt"

	aidego "github.com/hurbcom/aide-go/v4"
)

type Source struct {
	IntValue      int
	StringValue   string
	FloatValue    float64
	BoolValue     bool
	IntPointer    *int
	StringPointer *string
	FloatPointer  *float64
	BoolPointer   *bool
}

type Destination struct {
	A bool     `copier:"BoolValue"`
	B *float64 `copier:"FloatPointer"`
	C string   `copier:"StringValue"`
	D *string  `copier:"StringPointer"`
	E *int     `copier:"IntPointer"`
	F float64  `copier:"FloatValue"`
	G int      `copier:"IntValue"`
	H *bool    `copier:"BoolPointer"`
}

func main() {
	
	copier, err := aidego.NewCopier()
	if err != nil {
		panic(err)
	}

	data := GetSourceVector()
	for _, source := range data {

		var destination Destination
        if err = copier.Copy(&source, &destination); err != nil {
			fmt.Println(err)
			continue
		}
		
		ProccessDestination(destination)
    }
	
}

```

# TODOs

- Implement a benchmark comparing with jinzshu/copier.
- During copying, there are more opportunities to use cached data.
