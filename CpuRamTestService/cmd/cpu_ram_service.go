package main

import (
	"fmt"
	"time"
)

// hello world BountyCloud 测试服务
// @Author: hyc
// @Description:
// @Date: 2021/12/26 10:07
func main() {


	// CPU 打满
	go func() {
		for {
			a := 1
			b := 2
			c := a + b
			fmt.Printf("%+v\n",c)
		}
	}()

	// RAM 打满
	go func() {
		a := make([]byte,0)
		//for i := 0; i < 1024 * 1024 * 1024; i++ {
		//	a = append(a, 1)
		//}
		for i := 0; i < 1024; i++ {
			a = append(a, 1)
		}
		fmt.Printf("%+v\n","end")
		time.Sleep(time.Second * 1000)
	}()

	time.Sleep(time.Second * 1000)
}

