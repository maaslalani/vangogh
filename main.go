package main

import (
  "os"

  "github.com/nfnt/resize"
)

const (
  MAX_HEIGHT = 150
  MAX_WIDTH = 0
  RESIZE_STRATEGY = resize.NearestNeighbor
)

func main() {
  img, _, err := OpenImageFile(os.Args[1])
  HandleError(err)

  img = resize.Resize(MAX_WIDTH, MAX_HEIGHT, img, RESIZE_STRATEGY)

  ToGrayscale(img)
}
