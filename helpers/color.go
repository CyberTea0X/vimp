package helpers

import (
	"fmt"
	"image/color"
)

func RGBFromInt(rgb int) color.RGBA {
	r := uint8((rgb >> 16) & 0xFF)
	g := uint8((rgb >> 8) & 0xFF)
	b := uint8(rgb & 0xFF)
	return color.RGBA{R: r, G: g, B: b, A: 255}
}

func RGBAFromString(s string) color.RGBA {
	var r, g, b, a uint8
	if len(s) < 8 {
		return color.RGBA{}
	}
	r, g, b = RGBFromString(s[0:8])
	if len(s) == 10 {
		// 0xAABBEEFF
		// 0123456789
		fmt.Sscan("0x"+s[8:10], &a)
	} else {
		a = 255
	}
	return color.RGBA{R: r, G: g, B: b, A: a}
}

func RGBFromString(s string) (uint8, uint8, uint8) {
	if len(s) != 8 {
		return 0, 0, 0
	}
	// 0xAABBEE
	// 01234567
	var r, g, b uint8
	fmt.Sscan("0x"+s[2:4], &r)
	fmt.Sscan("0x"+s[4:6], &g)
	fmt.Sscan("0x"+s[6:8], &b)
	return r, g, b
}
