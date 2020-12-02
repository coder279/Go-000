package main

import (
	"fmt"
	"time"
)

func main() {
	Go(func() {
		panic("异常错误")
	})
	time.Sleep( 5 * time.Second)
}
//捕获野生goroutine
func Go(x func()){
	go func() {
		defer func() {
			if err := recover();err != nil {
				fmt.Println(err)
			}
		}()
		x()
	}()
}