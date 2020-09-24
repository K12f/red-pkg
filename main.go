package main

import (
	"bufio"
	"fmt"
	"os"
	"red-package/rekpkg"
)

var flagV int

//func init() {
//	flag.IntVar(&flagV, "plat", 1, "输入你想抢红包的平台: 1.微信,2:飞书")
//}

func main() {
	// 把用户传递的命令行参数解析为对应变量的值
	fmt.Println("输入你想抢红包的平台: 1.微信,2:飞书 ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	k := rekpkg.NewKernel()
	k.StartUp(input.Text())
}
