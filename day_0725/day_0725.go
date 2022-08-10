package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

//绘制正弦函数
func drawSinImg() {
	//图片分辨率
	const size = 300
	//创建灰度图
	img := image.NewGray(image.Rect(0, 0, size, size))
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			img.SetGray(i, j, color.Gray{Y: 255})
		}
	}
	//绘制曲线
	for x := 0; x < size; x++ {
		s := float64(x) / size * 2 * math.Pi
		y := size/2 - math.Sin(s)*size/2
		img.SetGray(x, int(y), color.Gray{})
	}
	//保存至图片文件
	file, err := os.Create("../images/sin.png")
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

var name = flag.String("name", "", "名称")
var age = flag.Int("age", 0, "年龄")

func main() {
	flag.Parse()
	testParseArgs()
}

//测试命令行参数的读取
func testParseArgs() {
	fmt.Printf("姓名%s,年龄%d", *name, *age)
	println()
}
