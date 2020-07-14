package main

import (
	"bytes"
	"image"
	"reflect"
	"sync"
	"syscall/js"
	"unsafe"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/imgio"
)

type Ctx struct {
	SetFileArrCb    js.Value
	SetImageToHueCb js.Value
}

func setFile(ctx *Ctx, fileJsArr js.Value, length int) {
	bs := make([]byte, length)
	ptr := (*reflect.SliceHeader)(unsafe.Pointer(&bs)).Data
	ctx.SetFileArrCb.Invoke(fileJsArr, ptr)

	img, _, _ := image.Decode(bytes.NewReader(bs))
	buf := &bytes.Buffer{}
	imgio.JPEGEncoder(93)(buf, adjust.Hue(img, -150))

	bs = buf.Bytes()
	ptr = (*reflect.SliceHeader)(unsafe.Pointer(&bs)).Data
	ctx.SetImageToHueCb.Invoke(ptr, len(bs))
}

func main() {
	jsGlobal := js.Global()
	ctx := &Ctx{
		SetFileArrCb:    jsGlobal.Get("setFileArrCb"),
		SetImageToHueCb: jsGlobal.Get("setImageToHueCb"),
	}

	goFuncs := jsGlobal.Get("goFuncs")
	goFuncs.Set("setFile", js.NewCallback(func(args []js.Value) {
		setFile(ctx, args[0], args[1].Int())
	}))

	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}