package main

import (
  "fmt"
  "math"

  "image/color"
  "image"
  _ "image/png"
  _ "image/jpeg"
)

const ASCII = " .:-=+*#%@"

func ToGrayscale(img image.Image) {
  bounds := img.Bounds()
  width, height := bounds.Max.X, bounds.Max.Y

  for y := 0; y < height; y += 2 {
    for x := 0; x < width; x += 1 {
      gray := colorToGrayscale(img.At(x, y))
      normalized := float32(gray.Y) / 255.0
      normalized *= float32(len(ASCII) - 1.0)
      index := int(normalized)
      fmt.Print(ASCII[index:index + 1])
    }
    fmt.Println()
  }
}

func colorToGrayscale(pixel color.Color) color.Gray {
  r, g, b, _ := pixel.RGBA()
  grayscale := grayValue(r, 0.2125) + grayValue(g, 0.7154) + grayValue(b, 0.0721)
  return color.Gray{uint8(uint16(math.Pow(grayscale, 1.0 / 2.2) + 0.5) >> 8)}
}

func grayValue(color uint32, multiplier float64) float64 {
  return multiplier * math.Pow(float64(color), 2.2)
}
