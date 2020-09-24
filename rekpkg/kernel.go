package rekpkg

import (
	"fmt"
	"log"
	"time"
)

type Kernel struct {
}

func NewKernel() Kernel {
	return Kernel{}
}

func (k Kernel) StartUp() {
	k.handle()
}

func (k Kernel) handle() {
	//1.读取配置文件
	fmt.Println("正在读取配置...")
	path := "./config.json"
	c := NewConfig()

	config, err := c.Get(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("正在拉取截图...")

	//2. 检测并拉去手机截图
	name := "screen.png"
	target := "./images/"
	adb := NewAdb()
	//
	fmt.Println("红包...")
	err = adb.Run(name, target, config.Red, 1)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("开...")
	nameOpen := "screen_open.png"
	err = adb.Run(nameOpen, target, config.Open, 2)
	if err != nil {
		log.Fatal(err)
	}
}
