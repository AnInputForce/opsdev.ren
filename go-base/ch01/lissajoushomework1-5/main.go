// lissajous 创造一个lissajous的git动画
// 练习1.5 ：修改前面的lissajous程序里的调色板，由黑色改为绿色。我们可以用
// color.RGBA{0xRR, 0xGG, 0xBB, 0xff}来得到 #RRGGBB这个色值，三个十六进制的字符串分别代表
// 红、绿、蓝像素
// 注意：此处给出的是伪代码。实际需要的绿色，十六进制表示为{0x00, 0xFF, 0x00, 0xff}

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.RGBA{0x00, 0xFF, 0x00, 0xff}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in pallette

)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

/* output:
ChinaDreams:lissajous kangcunhua$ go build
ChinaDreams:lissajous kangcunhua$ ./lissajoushomework1.5 > lissajous.gif
*/
