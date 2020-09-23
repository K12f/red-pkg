package rekpkg

import (
	"errors"
	"fmt"
	"os/exec"
)

type adb struct {
}

func NewAdb() adb {
	return adb{}
}

func (a adb) command(arg ...string) (err error) {

	if len(arg) < 4 {
		return errors.New("参数有误")
	}
	err = exec.Command("adb", arg[0], arg[1], arg[2], arg[3]).Run()
	if err != nil {
		err = errors.New("截图失败，请检查开发者选项中的 USB 调试安全设置是否打" + err.Error())
		return err
	}
	return err
}

// 拉去截屏的图片到 target 目录
func (a adb) Pull(name, target string) error {
	var err error
	imageName := fmt.Sprintf("/sdcard/%s", name)
	err = a.command("shell", "screencap", "-p", imageName)
	if err != nil {
		return err
	}
	err = a.command("shell", "pull", "-p", imageName, target)
	if err != nil {
		return err
	}
	err = a.command("shell", "pull", "rm", imageName, target)
	if err != nil {
		return err
	}
	return err
}
