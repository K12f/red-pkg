package rekpkg

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type imageR struct {
	im image.Image
}

type Result struct {
	x, y int
}

func NewimageR() *imageR {
	return &imageR{}
}

func (i *imageR) ReadPNG(filename string) error {
	var err error

	file, err := os.Open(filename)
	if err != nil {
		err = errors.New("读取图片失败")
		return err
	}
	defer file.Close()
	i.im, err = png.Decode(file)
	if err != nil {
		err = errors.New("PNG 截图解码失败" + err.Error())
		return err
	}
	return err
}

func (i imageR) Scan(col ColorR) (Result, error) {
	var result Result
	var err error
	var im = i.im
	width := im.Bounds().Max.X
	height := im.Bounds().Max.Y
	//2.扫描屏幕到下一步
	//widthMid := int(math.Ceil(float64(width / 2)))
	//heightMid := int(math.Ceil(float64(height / 2)))
	for h := height; h > 0; h-- {
		for w := 0; w < width; w++ {
			pointColor := im.At(w, h)

			r := pointColor.(color.NRGBA).R
			g := pointColor.(color.NRGBA).G
			b := pointColor.(color.NRGBA).B

			if r >= uint8(col.R-20) && r <= uint8(col.R) &&
				g >= uint8(col.G-20) && g <= uint8(col.G+20) &&
				b >= uint8(col.B-20) && b <= uint8(col.B+20) {

				fmt.Println(r, g, b)
				pointW := w
				pointH := h

				debug(im, w, pointH)

				return Result{pointW, pointH}, err
			}
		}
	}
	return result, errors.New("未发现相似的rgb")
}

func debug(im image.Image, width, height int) {
	des, _ := os.Create("./images/screen1.png")
	//_, err = io.Copy(des, file)

	defer des.Close()
	newIm := image.NewRGBA(im.Bounds())
	draw.Draw(newIm, im.Bounds(), im, newIm.Bounds().Min, draw.Src)
	red := color.NRGBA{0, 0, 0, 255}
	fmt.Println(width, height)
	for i := 0; i < 100; i++ {
		newIm.Set(width+i, height, red)
	}
	_ = png.Encode(des, newIm)
}
