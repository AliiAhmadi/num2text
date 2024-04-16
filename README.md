# num2text

converts numbers to persian text.

## Usage

### Library
```go
package main

import (
	"fmt"

	"github.com/AliiAhmadi/num2text"
)

func main() {
	number := "123"
	r, err := num2text.Convert(number)
	if err != nil {
		panic(err)
	}

	fmt.Println(r) // صد و بیست و سه
}
```

### CLI

```sh
Usage of num2text:
  -number string
        number you want to convert (default "0")
```

```sh
num2text -number=-.23
// Result: منفی بیست و سه صدم
```
