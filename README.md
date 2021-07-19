## wordpress-hash-go

WordPress hashes implemented with Golang

## Usage

### Start use it

Download and install

```bash
$ go get github.com/thundernet8/wordpress-hash-go
```

Import it in your code

```go
import (
    wphash "github.com/thundernet8/wordpress-hash-go"
)
```

### Canonical example:

```go
package main

import (
	"fmt"

	wphash "github.com/thundernet8/wordpress-hash-go"
)

func main() {
	password := "123456"
	// hash password
	hash := wphash.HashPassword(password)
	fmt.Printf("Password <%s> hash result is <%s>", password, hash)

	// verify password and hash
	match := wphash.CheckWordPressPasswordHash(password, "$P$BmIaPlVaAl6kEsffVZGdASCVH.i1cZ0")
	fmt.Printf("Check password <%s> with hash <%s> result is %t", match)
}
```