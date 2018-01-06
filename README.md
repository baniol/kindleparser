# Go package for parsing Kindle My Clippings file

## Example with json output

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/baniol/kindleparser"
)

func main() {

	records := kindleparser.ParseClippngs("My Clippings.txt")

	res, err := json.Marshal(records)

	if err != nil {
		panic("Parsing error")
	}

	fmt.Println(string(res))
}
```
