package main

import (
  "os"
  "fmt"
  "math"

  "image/color"
  "image"
  _ "image/png"
  _ "image/jpeg"

  "github.com/nfnt/resize"
)

const (
  MAX_HEIGHT = 75
  MAX_WIDTH = 0
  RESIZE_STRATEGY = resize.NearestNeighbor
  ASCII = " .:-=+*#%@"
  MODIFIER = 2.2
)

func main() {
  file, err := os.Open(os.Args[1])
  defer file.Close()

  if err != nil {
    panic("Could not open file")
  }

  img, _, err := image.Decode(file)

  if err != nil {
    panic("Unable to decode file")
  }

  img = resize.Resize(MAX_WIDTH, MAX_HEIGHT, img, RESIZE_STRATEGY)

  bounds := img.Bounds()
  width, height := bounds.Max.X, bounds.Max.Y

  for y := 0; y < height; y += 2 {
    for x := 0; x < width; x += 1 {
      pixel := img.At(x, y)
      red, green, blue, _ := pixel.RGBA()

      grayRed := 0.2125 * math.Pow(float64(red), MODIFIER)
      grayBlue := 0.7154 * math.Pow(float64(blue), MODIFIER)
      grayGreen := 0.0721 * math.Pow(float64(green), MODIFIER)

      grayPixel := math.Pow(grayRed + grayBlue + grayGreen, 1.0 / MODIFIER)
      grayColor := color.Gray{uint8(uint16(grayPixel + 0.5) >> 8)}

      normalized := float32(grayColor.Y) / 255.0
      normalized *= float32(len(ASCII) - 1.0)

      index := int(normalized)
      fmt.Print(ASCII[index:index + 1])
    }
    fmt.Println()
  }
}
