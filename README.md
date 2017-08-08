# gotify
a library to generate Goish variables, fields and package names from regular words

## Installation
```bash
go get -u github.com/sirkon/gotify
```
or
```bash
dep ensure -add github.com/sirkon/gotify
```

## Usage
```go
package main

import (
	"fmt"

	"github.com/sirkon/gotify"
)

func main() {
	g := gotify.New(map[string]string{
		"uid": "UID",
	})
	fmt.Println(g.Public("user_id"))
	fmt.Println(g.Private("uid_1"))
	fmt.Println(g.Private("uniqId"))
	fmt.Println(g.Package("string_sum"))
	fmt.Println(g.Goimports("string_sum"))
	fmt.Println(g.True("userID"))
	fmt.Println(g.True("1userID"))
}
```
will output
```
UserID
uid1
uniqID
stringsum
string-sum
true
false
```

