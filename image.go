package main

import (
  "os"
  "image"
  _ "image/png"
  _ "image/jpeg"
)

func OpenImageFile(filename string) (image.Image, error) {
  file, err := os.Open(filename)
  HandleError(err)

  defer file.Close()

  img, _, err := image.Decode(file)
  HandleError(err)

  return img, nil
}
