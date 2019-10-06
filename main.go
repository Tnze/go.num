package main

import (
	"fmt"

	"github.com/Tnze/go.num/v2/zh"
)

func main() {
	var num zh.Uint64
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(uint64(num))
}
