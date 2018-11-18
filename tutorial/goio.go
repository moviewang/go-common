package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	reader2 "golang.org/x/tour/reader"
	"image"
	"image/color"
	"io"
	"os"
	"strings"
)

type MyReader struct {
}

func (r MyReader) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = 'A'
	}
	return len(p), err
}

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	len, err := r.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return len, err
}

type Image struct {
	Height, Width int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Height, i.Width)
}

func (i Image) At(x, y int) color.Color {
	u := uint8(x ^ y)
	return color.RGBA{u,u, 255, 255}
}

func main() {
	s := "Hello, Reader!"
	fmt.Println("s len", len(s))
	reader := strings.NewReader(s)
	bytes := make([]byte, 8)
	for {
		n, err := reader.Read(bytes)
		fmt.Printf("n=%v, err = %v, bytes = %v\n", n, err, bytes)
		fmt.Printf("b[:n]= %q\n", bytes[:n])
		if io.EOF == err {
			break
		}
	}
	reader2.Validate(MyReader{})

	ss := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{ss}
	io.Copy(os.Stdout, &r)
	fmt.Println()
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0,0).RGBA())

	img := Image{256, 256}
	pic.ShowImage(img)
}
