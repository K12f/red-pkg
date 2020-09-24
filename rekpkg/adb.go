package rekpkg

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
)

type adb struct {
}

func NewAdb() *adb {
	return &adb{}
}

func (a adb) command(arg ...string) error {
	var err error
	if len(arg) < 3 {
		return errors.New("参数有误")
	}
	fmt.Println(arg)
	if len(arg) == 4 {
		err = exec.Command("adb", arg[0], arg[1], arg[2], arg[3]).Run()
	} else if len(arg) == 3 {
		err = exec.Command("adb", arg[0], arg[1], arg[2]).Run()
	} else if len(arg) == 5 {
		err = exec.Command("adb", arg[0], arg[1], arg[2], arg[3], arg[4]).Run()
	} else if len(arg) == 7 {
		err = exec.Command("adb", arg[0], arg[1], arg[2], arg[3], arg[4], arg[5], arg[6]).Run()
	} else {
		err = errors.New("不支持的参数数量")
	}
	if err != nil {
		err = errors.New("截图失败，请检查开发者选项中的 USB 调试安全设置是否打开" + err.Error())
		return err
	}
	return err
}

func (a adb) Run(name, target string, colorR ColorR, position uint) error {
	var err error
	err = a.Pull(name, target)
	if err != nil {
		return err

	}

	fmt.Println("正在读取分析截图...")
	//3.读取分析截图
	imageK := NewimageR()
	filename := fmt.Sprintf("%s%s", target, name)
	err = imageK.ReadPNG(filename)
	if err != nil {
		return err
	}
	fmt.Println("开始扫描图片...")

	//4.拿到句柄 开始扫描图片
	redPositionResult, err := imageK.Scan(colorR, position)
	if err != nil {
		return err
	}

	fmt.Println("开始模拟点击...", redPositionResult)

	//5.点击
	err = a.Touch(redPositionResult)
	if err != nil {
		return err
	}
	return err
}

// 拉去截屏的图片到 target 目录
func (a adb) Pull(name, target string) error {
	var err error
	imageName := fmt.Sprintf("/sdcard/%s", name)
	fmt.Println("screencap")
	err = a.command("shell", "screencap", "-p", imageName)
	if err != nil {
		return err
	}
	fmt.Println("pull")
	err = a.command("pull", "-p", imageName, target)
	if err != nil {
		return err
	}
	fmt.Println("rm")
	err = a.command("shell", "rm", imageName)
	if err != nil {
		return err
	}
	return err
}

func (a adb) Touch(result Result) error {
	var err error
	touchX, touchY := strconv.Itoa(result.x), strconv.Itoa(result.y)
	err = a.command("shell", "input", "tap", touchX, touchY)
	if err != nil {
		err = errors.New("模拟触摸失败，请检查开发者选项中的 USB 调试安全设置是否打开")
	}
	return err
}

func (a adb) Swipe() error {
	var err error
	err = a.command("shell", "input", "swipe", "800", "1000", "200", "300")
	if err != nil {
		err = errors.New("模拟触摸失败，请检查开发者选项中的 USB 调试安全设置是否打开")
	}
	return err
}

func (a adb) Click() error {
	var err error
	err = a.command("shell", "input", "tap", "200", "200")
	if err != nil {
		err = errors.New("模拟触摸失败，请检查开发者选项中的 USB 调试安全设置是否打开")
	}
	return err
}
