# go.num

Mainly implemented int->Chinese
主要实现了整数转中文算法

```go
package main

import (
	"fmt"
	"github.com/Tnze/go.num/v2/zh"
)

func main() {
	fmt.Print(zh.Uint64(1234).String())
	// Output: 一千二百三十四
}
```
