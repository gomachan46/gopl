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

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // パレットの最初の色
	blackIndex = 1 // パレットの次の色
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
  // 図形の設定値
	const (
		cycles  = 5 // 発振器xが完了する周回回数
		res     = 0.001 // 回転の分解能
		size    = 100 // 画像キャンパスは[-size..+size]の範囲
		nframes = 64 // アニメーションフレーム数
		delay   = 8 // 10ms単位でのフレーム間遅延
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ { // 各フレームごとの処理
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
	gif.EncodeAll(out, &anim) // 注意: エンコードエラーを無視
}

// コンポジットリテラル: []color.Color{...}とかgif.GIF{...}みたいなやつ。要素の値の並びからGoのコンポジット型のインスタンスを生成するための簡潔な表記
