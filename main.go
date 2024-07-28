package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
)

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	img, _, err := image.Decode(f)
	return img, err
}

func main() {
	fmt.Println("Hello, World!")
	s, _ := os.Getwd()
	img, err := getImageFromFilePath(s + "\\tree.png")
	fmt.Println("Hello, World!")
	fmt.Println(err)
	fmt.Println(img)
}
