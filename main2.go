package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type ColorRGB struct {
	r, g, b int
}

func main() {
	for {
		fmt.Println("正在扫描屏幕...")
		scan()
		fmt.Println("模拟操作结束...")
		time.Sleep(1 * time.Minute)
	}
}

func autoCap() image.Image {
	err := exec.Command("adb", "shell", "screencap", "-p", "/sdcard/screen.png").Run()
	if err != nil {
		log.Fatal("截图失败，请检查开发者选项中的 USB 调试安全设置是否打", err)
	}
	err = exec.Command("adb", "pull", "/sdcard/screen.png", "./images/").Run()
	if err != nil {
		log.Fatal("截图失败，请检查开发者选项中的 USB 调试安全设置是否打开", err)
	}
	err = exec.Command("adb", "shell", "rm", "/sdcard/screen.png").Run()
	if err != nil {
		log.Fatal("截图失败，请检查开发者选项中的 USB 调试安全设置是否打开", err)
	}

	file, err := os.Open("./images/screen.png")
	if err != nil {
		log.Fatal("读取图片失败")
	}
	defer file.Close()
	im, err := png.Decode(file)

	if err != nil {
		log.Fatal("PNG 截图解码失败.", err)
	}
	return im
}

func scan() {

	colorP := ColorRGB{
		r: 255,
		g: 97,
		b: 94,
	}
	//1.截图
	im := autoCap()
	//
	des, err := os.Create("./images/screen1.png")
	//_, err = io.Copy(des, file)
	checkError(err)
	defer des.Close()

	width := im.Bounds().Max.X
	height := im.Bounds().Max.Y
	//2.扫描屏幕到下一步
	widthMid := int(math.Ceil(float64(width / 2)))
	//heightMid := int(math.Ceil(float64(height / 2)))
	newIm := image.NewRGBA(im.Bounds())
	red := color.NRGBA{0, 0, 0, 255}

	draw.Draw(newIm, im.Bounds(), im, newIm.Bounds().Min, draw.Src)

	tempW := 0
	for w := 0; w < width; w++ {
		for h := height; h > 0; h-- {
			pointColor := im.At(w, h)

			r := pointColor.(color.NRGBA).R
			g := pointColor.(color.NRGBA).G
			b := pointColor.(color.NRGBA).B

			if r >= uint8(colorP.r-10) && r <= uint8(colorP.r) &&
				g >= uint8(colorP.g-20) && g <= uint8(colorP.g+20) &&
				b >= uint8(colorP.b-20) && b <= uint8(colorP.b+20) {

				//有一次颜色一样，记录一下
				tempW = (w + widthMid)

				pointColor := im.At(tempW, h)

				r2 := pointColor.(color.NRGBA).R
				g2 := pointColor.(color.NRGBA).G
				b2 := pointColor.(color.NRGBA).B
				if r2 == r && g2 == g && b2 == b {
					fmt.Println(r, g, b)
					pointW := w + widthMid/2
					pointH := h + 10
					newIm.Set(w+widthMid/2, h+10, red)
					autoTouch(pointW, pointH)
					goto Loop
				}

			} else {
				fmt.Println("未发现..")
			}
		}
	}
Loop:
	err = png.Encode(des, newIm)
	checkError(err)
	//3.自动点击
}

func autoTouch(x, y int) {
	touchX, touchY := strconv.Itoa(x), strconv.Itoa(y)
	err := exec.Command("adb", "shell", "input", "tap", touchX, touchY).Run()
	if err != nil {
		log.Fatal("模拟触摸失败，请检查开发者选项中的 USB 调试安全设置是否打开。")
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func debug() {

}
