package main

import (
	"flag"
	"red-package/rekpkg"
)

var flagV int

func init() {
	flag.IntVar(&flagV, "plat", 2, "输入你想抢红包的平台: 1.微信,2:飞书")
}

func main() {
	// 把用户传递的命令行参数解析为对应变量的值
	flag.Parse()
	k := rekpkg.NewKernel()
	k.StartUp(flagV)
}
