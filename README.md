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
	// 数字转中文
	fmt.Print(zh.Uint64(1234))
	// Output: 一千二百三十四
}
```

```go
func main() {
	// 中文转数字
	var num zh.Uint64
	_, err := fmt.Sscan("三百八十六万七千三百五十一", &num)
	if err != nil {
		// ....
	}

	fmt.Println(uint64(num))
	// Output: 3867351
}
```