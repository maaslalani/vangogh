package text

import (
  "fmt"
  "math"
  "os"

  "image"
  "image/color"
  _ "image/jpeg"
  _ "image/png"

  "github.com/nfnt/resize"
)

const (
  maxHeight      = 75
  maxWidth       = 0
  resizeStrategy = resize.Lanczos2
  asciiMap       = " .,:~=+?I7$ZO8DNM"
  modifier       = 2.2
)

func Command(filepath string) {
  file, err := os.Open(filepath)
  defer file.Close()

  if err != nil {
    panic("Could not open file")
  }

  img, _, err := image.Decode(file)

  if err != nil {
    panic("Unable to decode file")
  }

  img = resize.Resize(maxWidth, maxHeight, img, resizeStrategy)

  bounds := img.Bounds()
  width, height := bounds.Max.X, bounds.Max.Y

  for y := 0; y < height; y += 2 {
    for x := 0; x < width; x++ {
      pixel := img.At(x, y)
      red, green, blue, _ := pixel.RGBA()

      grayRed := 0.2125 * math.Pow(float64(red), modifier)
      grayBlue := 0.7154 * math.Pow(float64(blue), modifier)
      grayGreen := 0.0721 * math.Pow(float64(green), modifier)

      grayPixel := math.Pow(grayRed+grayBlue+grayGreen, 1.0/modifier)
      grayColor := color.Gray{uint8(uint16(grayPixel+0.5) >> 8)}

      normalized := float32(grayColor.Y) / 255.0
      normalized *= float32(len(asciiMap) - 1.0)

      index := int(normalized)
      fmt.Print(asciiMap[index : index+1])
    }
    fmt.Println()
  }
}
