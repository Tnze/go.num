package main

import (
	"fmt"

	"github.com/Tnze/go.num/v2/zh"
)

func main() {
	var num int64
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Print(err)
		return
	}
	if num >= 0 {
		fmt.Print(zh.String(uint64(num)))
	} else {
		fmt.Print("负" + zh.String(uint64(-num)))
	}
}
