package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestImage_countZeros(t *testing.T) {
	img := Image{pxs: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2}, height: 2, width: 3}

	if img.numLayers() != 2 {
		t.Errorf("Num layers: %v", img.numLayers())
		t.Fail()
	}

	countZeros := img.countNum(0)

	if countZeros[1] != 1 {
		t.Errorf("Num zeros: %v", countZeros)
		t.Fail()
	}

}

func TestRenderImage(t *testing.T) {

	img := Image{height: 2, width: 2, pxs: digitize("0222112222120000")}

	rendered := img.Render()

	assert.Equal(t, rendered.pxs, digitize("0110"))

}
