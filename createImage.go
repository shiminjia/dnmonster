package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)



func CreatePng(s string) {

	var g int = 10

	img := image.NewRGBA(image.Rect(0, 0, 12*g, 12*g))

	for i:=0; i< 6; i=i+1 {
		for j:=0; j<12; j=j+1{
			c := i+12*j
			fmt.Println(c)
			if string(s[c]) == "1" {
				for y:=i*g; y<(i+1)*g; y++{
					for z:=j*g; z<(j+1)*g; z++{
						img.Set(y, z, color.RGBA{255, 0, 0, 255})
					}
				}
			}
		}
	}

	for i:=0; i<6; i=i+1 {
		for j:=0; j<12; j=j+1{
			c := i+12*j
				fmt.Println(c)
			if string(s[c]) == "1" {
				for y:=i*g; y<(i+1)*g; y++{
					for z:=j*g; z<(j+1)*g; z++{
						img.Set(12*g-y-1, z, color.RGBA{255, 0, 0, 255})
					}
				}
			}
		}
	}

	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}