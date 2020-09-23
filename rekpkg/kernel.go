package rekpkg

import (
	"fmt"
	"log"
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
	path := "../config.json"
	config := NewConfig()

	config.Parse(path)
	//2. 检测并拉去手机截图
	name := "screen.png"
	target := "./images/"
	adb := NewAdb()

	err := adb.Pull(name, target)
	if err != nil {
		log.Fatal(err)
	}

	//3.读取分析截图
	imageK := NewimageR()
	filename := fmt.Sprintf("%s%s", target, name)
	err = imageK.ReadPNG(filename)
	if err != nil {
		log.Fatal(err)
	}
	//4.拿到句柄 开始扫描图片
}
