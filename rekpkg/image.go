package rekpkg

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
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

			if r >= uint8(col.r-10) && r <= uint8(col.r) &&
				g >= uint8(col.g-20) && g <= uint8(col.g+20) &&
				b >= uint8(col.b-20) && b <= uint8(col.b+20) {

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

					return Result{pointW, pointH}, err
				}
			}
		}
	}
	return result, errors.New("未发现相似的rgb")
}
