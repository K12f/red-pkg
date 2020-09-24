package rekpkg

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
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

func (i imageR) Scan(col ColorR, position uint) (Result, error) {
	var result Result
	var err error
	var im = i.im
	width := im.Bounds().Max.X
	height := im.Bounds().Max.Y
	//2.扫描屏幕到下一步
	widthMid := int(math.Ceil(float64(width / 2)))
	heightMid := int(math.Ceil(float64(height / 2)))
	heightStart := height
	heightEnd := 0
	widthStart := 0
	widthEnd := width

	switch position {
	//适配 开红包的位置
	case 1:
		fixValue := int(math.Ceil(float64(width / 10)))

		heightStart -= fixValue
		heightEnd += fixValue

		widthStart = widthMid - fixValue
		widthEnd = widthMid + fixValue
	case 2:
		//适配点击红包的位置
		fixValue := int(math.Ceil(float64(width / 10)))
		heightStart = int(math.Ceil(float64(height*3/4))) - fixValue
		heightEnd = heightMid + fixValue

		widthStart = widthMid - fixValue
		widthEnd = widthMid + fixValue
	}

	black := color.NRGBA{0, 0, 0, 255}
	red := color.NRGBA{255, 0, 0, 255}

	path := fmt.Sprintf("./images/screen%s.png", "tmep")
	des, _ := os.Create(path)
	defer des.Close()
	newIm := image.NewRGBA(im.Bounds())
	draw.Draw(newIm, im.Bounds(), im, newIm.Bounds().Min, draw.Src)

	for w := widthStart; w < widthEnd; w++ {
		for h := heightStart; h > heightEnd; h-- {
			pointColor := im.At(w, h)

			r := pointColor.(color.NRGBA).R
			g := pointColor.(color.NRGBA).G
			b := pointColor.(color.NRGBA).B

			//debug(im, black, 3, w, h)

			if r >= uint8(col.R-20) && r <= uint8(col.R) &&
				g >= uint8(col.G-20) && g <= uint8(col.G+20) &&
				b >= uint8(col.B-20) && b <= uint8(col.B+20) {
				newIm.Set(w, h, red)
				//debug(im, black, 2, w, h)

				err = png.Encode(des, newIm)
				if err != nil {
					log.Fatal(err)
				}

				return Result{w, h}, err
			} else {
				newIm.Set(w, h, black)
			}
		}
	}
	err = png.Encode(des, newIm)
	if err != nil {
		log.Fatal(err)
	}
	err = errors.New("未发现相似的rgb")
	return result, err
}

func debug(im image.Image, color color.NRGBA, name int, width, height int) {
	path := fmt.Sprintf("./images/screen%d.png", name)
	des, _ := os.Create(path)
	//_, err = io.Copy(des, file)

	defer des.Close()
	newIm := image.NewRGBA(im.Bounds())
	draw.Draw(newIm, im.Bounds(), im, newIm.Bounds().Min, draw.Src)
	fmt.Println(width, height)

	//newIm.Set(width, height, color)

	for i := 0; i < 100; i++ {
		newIm.Set(width+i, height, color)
	}
	_ = png.Encode(des, newIm)
}
