package main

import (
  "os"

  "github.com/nfnt/resize"
)

func main() {
  img, err := OpenImageFile(os.Args[1])
  HandleError(err)
  resizedImg := resize.Resize(0, 100, img, resize.NearestNeighbor)
  ToGrayscale(resizedImg)
}
