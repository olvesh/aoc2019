package main

import (
	"fmt"
	"os"
	"strconv"
)

type Image struct {
	pxs    []int
	width  int
	height int
}

func (img Image) layerSize() int {
	return img.width * img.height
}

func (img Image) numLayers() int {
	return len(img.pxs) / img.layerSize()
}

func digitize(img string) []int {
	imgInt := make([]int, 0, len(img))

	for _, char := range img {
		atoi, _ := strconv.Atoi(string(char))
		imgInt = append(imgInt, atoi)
	}
	return imgInt
}

func (img Image) countNum(toCount int) []int {
	count := make([]int, img.numLayers())
	for i := 0; i < img.numLayers(); i++ {

		count[i] = img.countLayer(i, toCount)
	}
	return count

}

func (img Image) countLayer(layer int, toCount int) int {
	count := 0
	for j := 0; j < img.layerSize(); j++ {
		if img.pxs[(layer*img.layerSize())+j] == toCount {
			count++
		}
	}
	return count
}

func (img Image) Render() Image {

	//0 black
	//1 white
	//2 transp

	rendered := Image{width: img.width, height: img.height, pxs: make([]int, img.layerSize())}
	for j := 0; j < img.layerSize(); j++ {
		for layer := 0; layer < img.numLayers(); layer++ {

			if img.pxs[(layer*img.layerSize())+j] != 2 {
				rendered.pxs[j] = img.pxs[(layer*img.layerSize())+j]
				break
			}
		}
	}

	return rendered
}

func (img Image) Print(stdout *os.File) {
	for i, px := range img.pxs {
		if i%img.width == 0 {
			fmt.Println()
		}
		if px == 1 {
			fmt.Printf("%q", rune('⬜'))
		} else if px == 0 {
			fmt.Printf("%q", rune('⬛'))
		}
	}

}
