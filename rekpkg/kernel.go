package rekpkg

import (
	"fmt"
	"time"
)

type Kernel struct {
}

func NewKernel() Kernel {
	return Kernel{}
}

func (k Kernel) StartUp(plat string) {
	k.handle(plat)
}

func (k Kernel) handle(plat string) {
	for {

		//1.读取配置文件
		fmt.Println("正在读取配置...")
		path := "./config.json"
		c := NewConfig()

		config, err := c.Get(plat, path)
		if err != nil {
			//log.Fatal(err)
			fmt.Println(err)
			//continue
		}

		fmt.Println("正在拉取截图...")

		//2. 检测并拉去手机截图
		name := "screen.png"
		target := "./images/"
		nameOpen := "screen_open.png"

		adb := NewAdb()
		//
		fmt.Println("红包...")
		err = adb.Run(name, target, config.Red, 1)
		if err != nil {
			fmt.Println(err)
			goto swipe
		}

		//等待loading红包出来，和网速之类的有关
		time.Sleep(500 * time.Millisecond)
		fmt.Println("开...")
		err = adb.Run(nameOpen, target, config.Open, 2)
		if err != nil {
			fmt.Println(err)
			goto swipe
		}

		//点击
		err = adb.Click()
		if err != nil {
			fmt.Println(err)
		}
		//下拉
		fmt.Println("下拉...")
	swipe:
		err = adb.Swipe()
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}
