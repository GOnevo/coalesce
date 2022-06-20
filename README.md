# Introduction

Coalesce is a simple function that returns the first non-nil value.

## Installation

```shell
go get github.com/gonevo/coalesce
```

## Example

```go
package main

import (
	"github.com/gonevo/coalesce"
	"log"
)

func main() {
	var err error
	log.Println(coalesce.Coalesce(nil, err, true))
}
```

## License

The coalesce library is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).
